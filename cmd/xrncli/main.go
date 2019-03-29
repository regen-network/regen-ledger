package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	upgradecli "github.com/cosmos/cosmos-sdk/x/upgrade/client/cli"
	upgraderest "github.com/cosmos/cosmos-sdk/x/upgrade/client/rest"
	"github.com/regen-network/regen-ledger"
	consortiumclient "github.com/regen-network/regen-ledger/x/consortium/client"
	dataclient "github.com/regen-network/regen-ledger/x/data/client"
	datarest "github.com/regen-network/regen-ledger/x/data/client/rest"
	espclient "github.com/regen-network/regen-ledger/x/esp/client"
	geoclient "github.com/regen-network/regen-ledger/x/geo/client"
	agentclient "github.com/regen-network/regen-ledger/x/group/client"
	proposalclient "github.com/regen-network/regen-ledger/x/proposal/client"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const (
	storeAcc      = "acc"
	storeData     = "data"
	storeAgent    = "group"
	storeProposal = "proposal"
	storeUpgrade  = "upgrade"
)

var defaultCLIHome = os.ExpandEnv("$HOME/.xrncli")

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.Seal()

	mc := []sdk.ModuleClients{
		espclient.NewModuleClient(cdc),
		proposalclient.NewModuleClient(storeProposal, cdc),
		geoclient.NewModuleClient(cdc),
		dataclient.NewModuleClient(storeData, cdc),
		agentclient.NewModuleClient(storeAgent, cdc),
		consortiumclient.NewModuleClient(cdc),
	}

	rootCmd := &cobra.Command{
		Use:   "xrncli",
		Short: "Regen Ledger Client",
	}

	// Construct Root Command
	rootCmd.AddCommand(
		initClientCommand(),
		rpc.StatusCommand(),
		client.ConfigCmd(defaultCLIHome),
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "XRN", defaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

const (
	// one of the following should be provided to verify the connection
	flagGenesis = "genesis"
	flagCommit  = "commit"
	flagValHash = "validator-set"
)

const defaultConfigTemplate = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The node to connect to
node = "{{ .Node }}"

# The chain ID
chain-id = "{{ .ChainId }}"
`

type cliConfig struct {
	Node    string
	ChainId string
}

// not implemented in "github.com/cosmos/cosmos-sdk/client/rpc"
// so implementing here
func initClientCommand() *cobra.Command {
	var home, node, chainId string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize light client",
		RunE: func(cmd *cobra.Command, args []string) error {
			configTemplate, err := template.New("configFileTemplate").Parse(defaultConfigTemplate)

			if err != nil {
				panic(err)
			}

			home = os.ExpandEnv(home)

			if err := cmn.EnsureDir(home, 0700); err != nil {
				cmn.PanicSanity(err.Error())
			}

			if err := cmn.EnsureDir(filepath.Join(home, "config"), 0700); err != nil {
				cmn.PanicSanity(err.Error())
			}

			configFilePath := filepath.Join(home, "config/config.toml")

			if !cmn.FileExists(configFilePath) {
				var buffer bytes.Buffer

				if err := configTemplate.Execute(&buffer, cliConfig{
					Node:    node,
					ChainId: chainId,
				}); err != nil {
					panic(err)
				}

				cmn.MustWriteFile(configFilePath, buffer.Bytes(), 0644)
			} else {
				fmt.Printf("%s already exists\n", configFilePath)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&home, cli.HomeFlag, "$HOME/.xrncli", "directory for config and data")
	cmd.Flags().StringVar(&chainId, client.FlagChainID, "", "ID of chain we connect to")
	cmd.Flags().StringVar(&node, client.FlagNode, "tcp://localhost:26657", "Node to connect to")
	//cmd.Flags().String(flagGenesis, "", "Genesis file to verify header validity")
	//cmd.Flags().String(flagCommit, "", "File with trusted and signed header")
	//cmd.Flags().String(flagValHash, "", "Hash of trusted validator set (hex-encoded)")
	//viper.BindPFlag(client.FlagChainID, cmd.Flags().Lookup(client.FlagChainID))
	//viper.BindPFlag(client.FlagNode, cmd.Flags().Lookup(client.FlagNode))

	return cmd
}

func registerRoutes(rs *lcd.RestServer) {
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeAcc)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	datarest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeData)
	upgraderest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, "upgrade-plan", storeUpgrade)
}

func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authcmd.GetAccountCmd(storeAcc, cdc),
	)

	for _, m := range mc {
		queryCmd.AddCommand(m.GetQueryCmd())
	}

	queryCmd.AddCommand(upgradecli.GetQueryCmd("upgrade-plan", storeUpgrade, cdc))

	addNodeFlags(queryCmd)

	return queryCmd
}

func txCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		client.LineBreak,
	)

	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}

	addNodeFlags(txCmd)

	return txCmd
}

func addNodeFlags(cmd *cobra.Command) {
	cmd.Flags().String(client.FlagNode, "tcp://localhost:26657", "Node to connect to")
	cmd.Flags().Bool(client.FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
	cmd.Flags().Bool(client.FlagChainID, false, "Chain ID of Tendermint node")
}
