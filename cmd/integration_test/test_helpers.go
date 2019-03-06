package clitest

import (
	"encoding/json"
	"fmt"
	"gitlab.com/regen-network/regen-ledger" // package app

	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/client/keys"
	appInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

const (
	denom        = "xrn"
	feeDenom	 = "xrn"
	fee2Denom	 = "xrn"
	fooDenom	 = "xrn"
	keyFoo       = "foo"
	keyBar       = "bar"
	keyBaz       = "baz"
	keyFooBarBaz = "foobarbaz"
	DefaultKeyPass = "12345678"
)

var (
	startCoins = sdk.Coins{
		sdk.NewCoin(denom, sdk.NewInt(1000000000)),
	}
)

//___________________________________________________________________________________
// Fixtures

// Fixtures is used to setup the testing environment
type Fixtures struct {
	ChainID  string
	RPCAddr  string
	Port     string
	XDHome   string
	XCLIHome string
	P2PAddr  string
	T        *testing.T
}

// NewFixtures creates a new instance of Fixtures with many vars set
func NewFixtures(t *testing.T) *Fixtures {
	tmpDir, err := ioutil.TempDir("", "xrn_integration_"+t.Name()+"_")
	require.NoError(t, err)
	servAddr, port, err := server.FreeTCPAddr()
	require.NoError(t, err)
	p2pAddr, _, err := server.FreeTCPAddr()
	require.NoError(t, err)
	return &Fixtures{
		T:        t,
		XDHome:   filepath.Join(tmpDir, ".xrnd"),
		XCLIHome: filepath.Join(tmpDir, ".xrncli"),
		RPCAddr:  servAddr,
		P2PAddr:  p2pAddr,
		Port:     port,
	}
}

// GenesisFile returns the path of the genesis file
func (f Fixtures) GenesisFile() string {
	return filepath.Join(f.XDHome, "config", "genesis.json")
}

// GenesisFile returns the application's genesis state
func (f Fixtures) GenesisState() app.GenesisState {
	cdc := codec.New()
	genDoc, err := appInit.LoadGenesisDoc(cdc, f.GenesisFile())
	require.NoError(f.T, err)

	var appState app.GenesisState
	require.NoError(f.T, cdc.UnmarshalJSON(genDoc.AppState, &appState))
	return appState
}

// InitFixtures is called at the beginning of a test
// and initializes a chain with 1 validator
func InitFixtures(t *testing.T) (f *Fixtures) {
	f = NewFixtures(t)

	// Reset test state
	f.UnsafeResetAll()

	// Ensure keystore has foo and bar keys
	f.KeysDelete(keyFoo)
	f.KeysDelete(keyBar)
	f.KeysDelete(keyBar)
	f.KeysDelete(keyFooBarBaz)
	f.KeysAdd(keyFoo)
	f.KeysAdd(keyBar)
	f.KeysAdd(keyBaz)
	f.KeysAdd(keyFooBarBaz, "--multisig-threshold=2", fmt.Sprintf(
		"--multisig=%s,%s,%s", keyFoo, keyBar, keyBaz))

	// Ensure that CLI output is in JSON format
	f.CLIConfig("output", "json")

	// NOTE: XDInit sets the ChainID
	f.XDInit(keyFoo)
	f.CLIConfig("chain-id", f.ChainID)

	// Start an account with tokens
	f.AddGenesisAccount(f.KeyAddress(keyFoo), startCoins)
	f.GenTx(keyFoo)
	f.CollectGenTxs()
	return
}

// Cleanup is meant to be run at the end of a test to clean up an remaining test state
func (f *Fixtures) Cleanup(dirs ...string) {
	clean := append(dirs, f.XDHome, f.XCLIHome)
	for _, d := range clean {
		err := os.RemoveAll(d)
		require.NoError(f.T, err)
	}
}

// Flags returns the flags necessary for making most CLI calls
func (f *Fixtures) Flags() string {
	return fmt.Sprintf("--home=%s --node=%s", f.XCLIHome, f.RPCAddr)
}

//___________________________________________________________________________________
// xrnd

// UnsafeResetAll is xrnd unsafe-reset-all
func (f *Fixtures) UnsafeResetAll(flags ...string) {
	cmd := fmt.Sprintf("xrnd --home=%s unsafe-reset-all", f.XDHome)
	executeWrite(f.T, addFlags(cmd, flags))
	err := os.RemoveAll(filepath.Join(f.XDHome, "config", "gentx"))
	require.NoError(f.T, err)
}

// XDInit is xrnd init
// NOTE: XDInit sets the ChainID for the Fixtures instance
func (f *Fixtures) XDInit(moniker string, flags ...string) {
	cmd := fmt.Sprintf("xrnd init -o --home=%s %s", f.XDHome, moniker)
	_, stderr := tests.ExecuteT(f.T, addFlags(cmd, flags), DefaultKeyPass)

	var chainID string
	var initRes map[string]json.RawMessage

	err := json.Unmarshal([]byte(stderr), &initRes)
	require.NoError(f.T, err)

	err = json.Unmarshal(initRes["chain_id"], &chainID)
	require.NoError(f.T, err)

	f.ChainID = chainID
}

// AddGenesisAccount is xrnd add-genesis-account
func (f *Fixtures) AddGenesisAccount(address sdk.AccAddress, coins sdk.Coins, flags ...string) {
	cmd := fmt.Sprintf("xrnd add-genesis-account %s %s --home=%s", address, coins, f.XDHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

// GenTx is xrnd gentx
func (f *Fixtures) GenTx(name string, flags ...string) {
	cmd := fmt.Sprintf("xrnd gentx --name=%s --home=%s --home-client=%s", name, f.XDHome, f.XCLIHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// CollectGenTxs is xrnd collect-gentxs
func (f *Fixtures) CollectGenTxs(flags ...string) {
	cmd := fmt.Sprintf("xrnd collect-gentxs --home=%s", f.XDHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// XDStart runs xrnd start with the appropriate flags and returns a process
func (f *Fixtures) XDStart(flags ...string) *tests.Process {
	cmd := fmt.Sprintf("xrnd start --home=%s --rpc.laddr=%v --p2p.laddr=%v", f.XDHome, f.RPCAddr, f.P2PAddr)
	proc := tests.GoExecuteTWithStdout(f.T, addFlags(cmd, flags))
	tests.WaitForTMStart(f.Port)
	tests.WaitForNextNBlocksTM(1, f.Port)
	return proc
}

// GDTendermint returns the results of xrnd tendermint [query]
func (f *Fixtures) GDTendermint(query string) string {
	cmd := fmt.Sprintf("xrnd tendermint %s --home=%s", query, f.XDHome)
	success, stdout, stderr := executeWriteRetStdStreams(f.T, cmd)
	require.Empty(f.T, stderr)
	require.True(f.T, success)
	return strings.TrimSpace(stdout)
}

// ValidateGenesis runs xrnd validate-genesis
func (f *Fixtures) ValidateGenesis() {
	cmd := fmt.Sprintf("xrnd validate-genesis --home=%s", f.XDHome)
	executeWriteCheckErr(f.T, cmd)
}

//___________________________________________________________________________________
// xrncli keys

// KeysDelete is xrncli keys delete
func (f *Fixtures) KeysDelete(name string, flags ...string) {
	cmd := fmt.Sprintf("xrncli keys delete --home=%s %s", f.XCLIHome, name)
	executeWrite(f.T, addFlags(cmd, append(append(flags, "-y"), "-f")))
}

// KeysAdd is xrncli keys add
func (f *Fixtures) KeysAdd(name string, flags ...string) {
	cmd := fmt.Sprintf("xrncli keys add --home=%s %s", f.XCLIHome, name)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// KeysAddRecover prepares xrncli keys add --recover
func (f *Fixtures) KeysAddRecover(name, mnemonic string, flags ...string) {
	cmd := fmt.Sprintf("xrncli keys add --home=%s --recover %s", f.XCLIHome, name)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass, mnemonic)
}

// KeysAddRecoverHDPath prepares xrncli keys add --recover --account --index
func (f *Fixtures) KeysAddRecoverHDPath(name, mnemonic string, account uint32, index uint32, flags ...string) {
	cmd := fmt.Sprintf("xrncli keys add --home=%s --recover %s --account %d --index %d", f.XCLIHome, name, account, index)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), DefaultKeyPass, mnemonic)
}

// KeysShow is xrncli keys show
func (f *Fixtures) KeysShow(name string, flags ...string) keys.KeyOutput {
	cmd := fmt.Sprintf("xrncli keys show --home=%s %s", f.XCLIHome, name)
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var ko keys.KeyOutput
	err := keys.UnmarshalJSON([]byte(out), &ko)
	require.NoError(f.T, err)
	return ko
}

// KeyAddress returns the SDK account address from the key
func (f *Fixtures) KeyAddress(name string) sdk.AccAddress {
	ko := f.KeysShow(name)
	accAddr, err := sdk.AccAddressFromBech32(ko.Address)
	require.NoError(f.T, err)
	return accAddr
}

//___________________________________________________________________________________
// xrncli config

// CLIConfig is xrncli config
func (f *Fixtures) CLIConfig(key, value string, flags ...string) {
	cmd := fmt.Sprintf("xrncli config --home=%s %s %s", f.XCLIHome, key, value)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

//___________________________________________________________________________________
// xrncli tx send/sign/broadcast

// TxSend is xrncli tx send
func (f *Fixtures) TxSend(from string, to sdk.AccAddress, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("xrncli tx send %s %s %v --from=%s", to, amount, f.Flags(), from)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxSign is xrncli tx sign
func (f *Fixtures) TxSign(signer, fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("xrncli tx sign %v --name=%s %v", f.Flags(), signer, fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxBroadcast is xrncli tx broadcast
func (f *Fixtures) TxBroadcast(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("xrncli tx broadcast %v %v", f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxEncode is xrncli tx encode
func (f *Fixtures) TxEncode(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("xrncli tx encode %v %v", f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

// TxMultisign is xrncli tx multisign
func (f *Fixtures) TxMultisign(fileName, name string, signaturesFiles []string,
	flags ...string) (bool, string, string) {

	cmd := fmt.Sprintf("xrncli tx multisign %v %s %s %s", f.Flags(),
		fileName, name, strings.Join(signaturesFiles, " "),
	)
	return executeWriteRetStdStreams(f.T, cmd)
}

//___________________________________________________________________________________
// xrncli query account

// QueryAccount is xrncli query account
func (f *Fixtures) QueryAccount(address sdk.AccAddress, flags ...string) auth.BaseAccount {
	cmd := fmt.Sprintf("xrncli query account %s %v", address, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var initRes map[string]json.RawMessage
	err := json.Unmarshal([]byte(out), &initRes)
	require.NoError(f.T, err, "out %v, err %v", out, err)
	value := initRes["value"]
	var acc auth.BaseAccount
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	err = cdc.UnmarshalJSON(value, &acc)
	require.NoError(f.T, err, "value %v, err %v", string(value), err)
	return acc
}

//___________________________________________________________________________________
// xrncli query txs

// QueryTxs is xrncli query txs
func (f *Fixtures) QueryTxs(page, limit int, tags ...string) []sdk.TxResponse {
	cmd := fmt.Sprintf("xrncli query txs --page=%d --limit=%d --tags='%s' %v", page, limit, queryTags(tags), f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var txs []sdk.TxResponse
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &txs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return txs
}

// QueryTxsInvalid query txs with wrong parameters and compare expected error
func (f *Fixtures) QueryTxsInvalid(expectedErr error, page, limit int, tags ...string) {
	cmd := fmt.Sprintf("xrncli query txs --page=%d --limit=%d --tags='%s' %v", page, limit, queryTags(tags), f.Flags())
	_, err := tests.ExecuteT(f.T, cmd, "")
	require.EqualError(f.T, expectedErr, err)
}

//___________________________________________________________________________________
// executors

func executeWriteCheckErr(t *testing.T, cmdStr string, writes ...string) {
	require.True(t, executeWrite(t, cmdStr, writes...))
}

func executeWrite(t *testing.T, cmdStr string, writes ...string) (exitSuccess bool) {
	exitSuccess, _, _ = executeWriteRetStdStreams(t, cmdStr, writes...)
	return
}

func executeWriteRetStdStreams(t *testing.T, cmdStr string, writes ...string) (bool, string, string) {
	proc := tests.GoExecuteT(t, cmdStr)

	// Enables use of interactive commands
	for _, write := range writes {
		_, err := proc.StdinPipe.Write([]byte(write + "\n"))
		require.NoError(t, err)
	}

	// Read both stdout and stderr from the process
	stdout, stderr, err := proc.ReadAll()
	if err != nil {
		fmt.Println("Err on proc.ReadAll()", err, cmdStr)
	}

	// Log output.
	if len(stdout) > 0 {
		t.Log("Stdout:", cmn.Green(string(stdout)))
	}
	if len(stderr) > 0 {
		t.Log("Stderr:", cmn.Red(string(stderr)))
	}

	// Wait for process to exit
	proc.Wait()

	// Return succes, stdout, stderr
	return proc.ExitState.Success(), string(stdout), string(stderr)
}

//___________________________________________________________________________________
// utils

func addFlags(cmd string, flags []string) string {
	for _, f := range flags {
		cmd += " " + f
	}
	return strings.TrimSpace(cmd)
}

func queryTags(tags []string) (out string) {
	for _, tag := range tags {
		out += tag + "&"
	}
	return strings.TrimSuffix(out, "&")
}

// Write the given string to a new temporary file
func WriteToNewTempFile(t *testing.T, s string) *os.File {
	fp, err := ioutil.TempFile(os.TempDir(), "cosmos_cli_test_")
	require.Nil(t, err)
	_, err = fp.WriteString(s)
	require.Nil(t, err)
	return fp
}

func marshalStdTx(t *testing.T, stdTx auth.StdTx) []byte {
	cdc := app.MakeCodec()
	bz, err := cdc.MarshalBinaryBare(stdTx)
	require.NoError(t, err)
	return bz
}

func unmarshalStdTx(t *testing.T, s string) (stdTx auth.StdTx) {
	cdc := app.MakeCodec()
	require.Nil(t, cdc.UnmarshalJSON([]byte(s), &stdTx))
	return
}
