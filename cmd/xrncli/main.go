package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/regen-network/regen-ledger/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"
	"os"
	"path"

	sdk "github.com/cosmos/cosmos-sdk/types"
	at "github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	crisisclient "github.com/cosmos/cosmos-sdk/x/crisis/client"
	distcmd "github.com/cosmos/cosmos-sdk/x/distribution"
	distClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrcli "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	dist "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	gv "github.com/cosmos/cosmos-sdk/x/gov"
	govClient "github.com/cosmos/cosmos-sdk/x/gov/client"
	gov "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	mintclient "github.com/cosmos/cosmos-sdk/x/mint/client"
	mintrest "github.com/cosmos/cosmos-sdk/x/mint/client/rest"
	paramcli "github.com/cosmos/cosmos-sdk/x/params/client/cli"
	paramsrest "github.com/cosmos/cosmos-sdk/x/params/client/rest"
	sl "github.com/cosmos/cosmos-sdk/x/slashing"
	slashingclient "github.com/cosmos/cosmos-sdk/x/slashing/client"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	st "github.com/cosmos/cosmos-sdk/x/staking"
	stakingclient "github.com/cosmos/cosmos-sdk/x/staking/client"
	staking "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	consortiumclient "github.com/regen-network/regen-ledger/x/consortium/client"
	dataclient "github.com/regen-network/regen-ledger/x/data/client"
	datarest "github.com/regen-network/regen-ledger/x/data/client/rest"
	espclient "github.com/regen-network/regen-ledger/x/esp/client"
	geoclient "github.com/regen-network/regen-ledger/x/geo/client"
	agentclient "github.com/regen-network/regen-ledger/x/group/client"
	claimclient "github.com/regen-network/regen-ledger/x/claim/client"
	proposalclient "github.com/regen-network/regen-ledger/x/proposal/client"
	upgraderest "github.com/regen-network/regen-ledger/x/upgrade/client/rest"
)

const (
	storeAcc      = "acc"
	storeClaim     = "claim"
	storeData     = "data"
	storeAgent    = "group"
	storeProposal = "proposal"
	storeUpgrade  = "upgrade"
)

func main() {
	// Configure cobra to sort commands
	cobra.EnableCommandSorting = false

	// Instantiate the codec for the command line application
	cdc := app.MakeCodec()

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.Seal()

	mc := []sdk.ModuleClient{
		govClient.NewModuleClient(gv.StoreKey, cdc, paramcli.GetCmdSubmitProposal(cdc), distrcli.GetCmdSubmitProposal(cdc)),
		distClient.NewModuleClient(distcmd.StoreKey, cdc),
		stakingclient.NewModuleClient(st.StoreKey, cdc),
		mintclient.NewModuleClient(mint.StoreKey, cdc),
		slashingclient.NewModuleClient(sl.StoreKey, cdc),
		crisisclient.NewModuleClient(sl.StoreKey, cdc),
		espclient.NewModuleClient(cdc),
		proposalclient.NewModuleClient(storeProposal, cdc),
		geoclient.NewModuleClient(cdc),
		dataclient.NewModuleClient(storeData, cdc),
		agentclient.NewModuleClient(storeAgent, cdc),
		consortiumclient.NewModuleClient(cdc),
		claimclient.NewModuleClient(storeClaim, cdc),
	}

	rootCmd := &cobra.Command{
		Use:   "xrncli",
		Short: "Command line interface for interacting with xrnd",
	}

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(app.DefaultCLIHome),
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.Cmd,
		client.NewCompletionCmd(rootCmd, true),
	)

	// Add flags and prefix all env exposed with GA
	executor := cli.PrepareMainCmd(rootCmd, "GA", app.DefaultCLIHome)

	err := executor.Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command: %s, exiting...\n", err)
		os.Exit(1)
	}
}

func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClient) *cobra.Command {
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
		authcmd.GetAccountCmd(at.StoreKey, cdc),
	)

	for _, m := range mc {
		mQueryCmd := m.GetQueryCmd()
		if mQueryCmd != nil {
			queryCmd.AddCommand(mQueryCmd)
		}
	}

	return queryCmd
}

func txCmd(cdc *amino.Codec, mc []sdk.ModuleClient) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
		client.LineBreak,
	)

	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}

	return txCmd
}

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	rpc.RegisterRPCRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterTxRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, at.StoreKey)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	dist.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, distcmd.StoreKey)
	staking.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashing.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	gov.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, paramsrest.ProposalRESTHandler(rs.CliCtx, rs.Cdc), dist.ProposalRESTHandler(rs.CliCtx, rs.Cdc))
	mintrest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	datarest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeData)
	upgraderest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, "upgrade-plan", storeUpgrade)
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
