package group_test

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/cosmos/cosmos-sdk/codec"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/errors"
// 	"github.com/cosmos/cosmos-sdk/x/auth/types"
// 	proto "github.com/gogo/protobuf/types"
// 	"github.com/regen-network/regen-ledger/orm"
// 	"github.com/regen-network/regen-ledger/x/group"
// 	"github.com/regen-network/regen-ledger/x/group/testdata"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	abci "github.com/tendermint/tendermint/abci/types"
// 	"github.com/tendermint/tendermint/crypto"
// 	"github.com/tendermint/tendermint/libs/log"
// 	dbm "github.com/tendermint/tm-db"
// )

// func createTestApp(isCheckTx bool) (*testdata.SimApp, sdk.Context) {
// 	db := dbm.NewMemDB()
// 	app := testdata.NewSimApp(log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, "", 0)
// 	genesisState := testdata.ModuleBasics.DefaultGenesis(app.AppCodec())
// 	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Initialize the chain
// 	app.InitChain(
// 		abci.RequestInitChain{
// 			Validators:    []abci.ValidatorUpdate{},
// 			AppStateBytes: stateBytes,
// 		},
// 	)
// 	app.Commit()

// 	header := abci.Header{Height: app.LastBlockHeight() + 1, Time: time.Now()}
// 	app.BeginBlock(abci.RequestBeginBlock{Header: header})

// 	ctx := app.NewContext(isCheckTx, header)
// 	return app, ctx
// }

// func TestCreateGroupScenario(t *testing.T) {
// 	app, ctx := createTestApp(false)
// 	myKey, _, myAddr := types.KeyTestPubAddr()
// 	myAccount := app.AccountKeeper.NewAccountWithAddress(ctx, myAddr)
// 	app.AccountKeeper.SetAccount(ctx, myAccount)

// 	_, _, otherAddr := types.KeyTestPubAddr()

// 	balances := sdk.NewCoins(sdk.NewInt64Coin("atom", 1000))
// 	require.NoError(t, app.BankKeeper.SetBalances(ctx, myAddr, balances))

// 	fee := types.NewTestStdFee()
// 	specs := map[string]struct {
// 		src     group.MsgCreateGroup
// 		expCode uint32
// 	}{
// 		"happy path": {
// 			src: group.MsgCreateGroup{
// 				Admin: myAddr,
// 				Members: []group.Member{{
// 					Address: myAddr,
// 					Power:   sdk.NewDec(1),
// 					Comment: "foo",
// 				}},
// 				Comment: "integration test",
// 			},
// 		},
// 		"invalid message": {
// 			src: group.MsgCreateGroup{
// 				Admin: myAddr,
// 				Members: []group.Member{{
// 					Address: myAddr,
// 					Power:   sdk.NewDec(0),
// 					Comment: "invalid power",
// 				}},
// 				Comment: "integration test",
// 			},
// 			expCode: group.ErrEmpty.ABCICode(),
// 		},
// 		"invalid signer": {
// 			src: group.MsgCreateGroup{
// 				Admin:   otherAddr,
// 				Comment: "admin and signer do not match",
// 			},
// 			expCode: errors.ErrInvalidPubKey.ABCICode(),
// 		},
// 	}
// 	var seq uint64
// 	privs, accNums := []crypto.PrivKey{myKey}, myAccount.GetAccountNumber()
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			accSeq, err := app.AccountKeeper.GetSequence(ctx, myAddr)
// 			require.NoError(t, err)
// 			tx := types.NewTestTx(ctx, []sdk.Msg{spec.src}, privs, []uint64{accNums}, []uint64{accSeq}, fee)
// 			resp := app.DeliverTx(abci.RequestDeliverTx{Tx: app.MustEncodeTx(tx)})
// 			// then
// 			require.Equal(t, spec.expCode, resp.Code, resp.Log)
// 			if spec.expCode != 0 {
// 				return
// 			}
// 			seq++
// 			assert.Equal(t, orm.EncodeSequence(seq), resp.Data)
// 			assert.True(t, app.GroupKeeper.HasGroup(ctx, resp.Data))
// 		})
// 	}
// }

// func TestCreateGroupAccountScenario(t *testing.T) {
// 	app, ctx := createTestApp(false)
// 	myKey, _, myAddr := types.KeyTestPubAddr()
// 	myAccount := app.AccountKeeper.NewAccountWithAddress(ctx, myAddr)
// 	app.AccountKeeper.SetAccount(ctx, myAccount)

// 	_, _, otherAddr := types.KeyTestPubAddr()

// 	balances := sdk.NewCoins(sdk.NewInt64Coin("atom", 10000))
// 	require.NoError(t, app.BankKeeper.SetBalances(ctx, myAddr, balances))

// 	myGroupID, err := app.GroupKeeper.CreateGroup(ctx, myAddr, nil, "integration test")
// 	require.NoError(t, err)

// 	fee := types.NewTestStdFee()
// 	specs := map[string]struct {
// 		admin     sdk.AccAddress
// 		group     group.GroupID
// 		comment   string
// 		threshold sdk.Dec
// 		timeout   proto.Duration
// 		expCode   uint32
// 	}{
// 		"happy path": {
// 			admin:     myAddr,
// 			group:     myGroupID,
// 			comment:   "integration test",
// 			threshold: sdk.OneDec(),
// 			timeout:   proto.Duration{Seconds: 1},
// 		},
// 		"second account with same group": {
// 			admin:     myAddr,
// 			group:     myGroupID,
// 			comment:   "integration test",
// 			threshold: sdk.OneDec(),
// 			timeout:   proto.Duration{Seconds: 1},
// 		},
// 		"unknown group in message": {
// 			admin:     myAddr,
// 			group:     99999,
// 			comment:   "group id does not exists",
// 			threshold: sdk.OneDec(),
// 			timeout:   proto.Duration{Seconds: 1},
// 			expCode:   orm.ErrNotFound.ABCICode(),
// 		},
// 		"invalid signer": {
// 			admin:     otherAddr,
// 			group:     myGroupID,
// 			comment:   "integration test",
// 			threshold: sdk.OneDec(),
// 			timeout:   proto.Duration{Seconds: 1},
// 			expCode:   errors.ErrInvalidPubKey.ABCICode(),
// 		},
// 	}

// 	var seq uint64
// 	privs, accNums := []crypto.PrivKey{myKey}, myAccount.GetAccountNumber()
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			m, err := group.NewMsgCreateGroupAccount(
// 				spec.admin,
// 				spec.group,
// 				spec.comment,
// 				group.NewThresholdDecisionPolicy(
// 					spec.threshold,
// 					spec.timeout,
// 				),
// 			)

// 			require.NoError(t, err)

// 			msgs := []sdk.Msg{m}
// 			accSeq, err := app.AccountKeeper.GetSequence(ctx, myAddr)
// 			require.NoError(t, err)
// 			tx := types.NewTestTx(ctx, msgs, privs, []uint64{accNums}, []uint64{accSeq}, fee)
// 			resp := app.DeliverTx(abci.RequestDeliverTx{Tx: app.MustEncodeTx(tx)})
// 			// then
// 			require.Equal(t, spec.expCode, resp.Code, resp.Log)
// 			if spec.expCode != 0 {
// 				return
// 			}
// 			seq++
// 			assert.Equal(t, group.AccountCondition(seq).Address().Bytes(), resp.Data)
// 			assert.True(t, app.GroupKeeper.HasGroupAccount(ctx, resp.Data))
// 		})
// 	}
// }

// func TestFullProposalWorkflow(t *testing.T) {
// 	app, ctx := createTestApp(false)

// 	// setup account
// 	myKey, _, myAddr := types.KeyTestPubAddr()
// 	myAccount := app.AccountKeeper.NewAccountWithAddress(ctx, myAddr)
// 	app.AccountKeeper.SetAccount(ctx, myAccount)

// 	balances := sdk.NewCoins(sdk.NewInt64Coin("atom", 100000))
// 	require.NoError(t, app.BankKeeper.SetBalances(ctx, myAddr, balances))

// 	msgFirstAccount, err := group.NewMsgCreateGroupAccount(
// 		myAddr,
// 		1,
// 		"first account",
// 		group.NewThresholdDecisionPolicy(
// 			sdk.OneDec(),
// 			*proto.DurationProto(time.Second),
// 		),
// 	)
// 	require.NoError(t, err)

// 	msgSecondAccount, err := group.NewMsgCreateGroupAccount(
// 		myAddr,
// 		1,
// 		"second account",
// 		group.NewThresholdDecisionPolicy(
// 			sdk.OneDec(),
// 			*proto.DurationProto(time.Second),
// 		),
// 	)
// 	require.NoError(t, err)

// 	// setup group
// 	msgs := []sdk.Msg{
// 		group.MsgCreateGroup{
// 			Admin: myAddr,
// 			Members: []group.Member{{
// 				Address: myAddr,
// 				Power:   sdk.OneDec(),
// 				Comment: "me",
// 			}},
// 			Comment: "integration test",
// 		},
// 		// setup group account
// 		msgFirstAccount,
// 		// and another one
// 		msgSecondAccount,
// 		// submit proposals
// 		testdata.MsgPropose{
// 			Base: group.MsgProposeBase{
// 				GroupAccount: group.AccountCondition(1).Address(), // first account
// 				Proposers:    []sdk.AccAddress{myAddr},
// 				Comment:      "ok",
// 			},
// 			Msgs: []testdata.MyAppMsg{{Sum: &testdata.MyAppMsg_A{A: &testdata.MsgAlwaysSucceed{}}}},
// 		},
// 		testdata.MsgPropose{
// 			Base: group.MsgProposeBase{
// 				GroupAccount: group.AccountCondition(2).Address(), // second account, same group
// 				Proposers:    []sdk.AccAddress{myAddr},
// 				Comment:      "other proposal",
// 			},
// 		},
// 		// vote
// 		group.MsgVote{
// 			Proposal: 1,
// 			Voters:   []sdk.AccAddress{myAddr},
// 			Choice:   group.Choice_YES,
// 			Comment:  "makes sense",
// 		},
// 		group.MsgVote{
// 			Proposal: 2,
// 			Voters:   []sdk.AccAddress{myAddr},
// 			Choice:   group.Choice_VETO,
// 			Comment:  "no way",
// 		},
// 	}

// 	fee := types.NewStdFee(200000, sdk.NewCoins(sdk.NewInt64Coin("atom", 150)))
// 	privs, accNums, seqs := []crypto.PrivKey{myKey}, myAccount.GetAccountNumber(), myAccount.GetSequence()
// 	tx := types.NewTestTx(ctx, msgs, privs, []uint64{accNums}, []uint64{seqs}, fee)

// 	resp := app.DeliverTx(abci.RequestDeliverTx{Tx: app.MustEncodeTx(tx)})
// 	require.Equal(t, uint32(0), resp.Code, resp.Log)

// 	// execute can not be in the same block so start new one
// 	app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: app.LastBlockHeight() + 1, Time: time.Now()}})

// 	// execute first proposal
// 	msgs = []sdk.Msg{
// 		group.MsgExec{
// 			Proposal: 1,
// 			Signer:   myAddr,
// 		},
// 	}
// 	myAccount = app.AccountKeeper.GetAccount(ctx, myAddr)
// 	privs, accNums, seqs = []crypto.PrivKey{myKey}, myAccount.GetAccountNumber(), myAccount.GetSequence()
// 	tx = types.NewTestTx(ctx, msgs, privs, []uint64{accNums}, []uint64{seqs}, fee)

// 	resp = app.DeliverTx(abci.RequestDeliverTx{Tx: app.MustEncodeTx(tx)})
// 	require.Equal(t, uint32(0), resp.Code, resp.Log)

// 	// then verify proposal got accepted
// 	proposal, err := app.GroupKeeper.GetProposal(ctx, 1)
// 	require.NoError(t, err)
// 	assert.Equal(t, group.ProposalResultAccepted, proposal.GetBase().Result, proposal.GetBase().Result.String())
// 	assert.Equal(t, group.ProposalStatusClosed, proposal.GetBase().Status, proposal.GetBase().Status.String())
// 	expTally := group.Tally{YesCount: sdk.OneDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.ZeroDec()}
// 	assert.Equal(t, expTally, proposal.GetBase().VoteState)

// 	// execute second proposal
// 	msgs = []sdk.Msg{
// 		group.MsgExec{
// 			Proposal: 2,
// 			Signer:   myAddr,
// 		},
// 	}
// 	myAccount = app.AccountKeeper.GetAccount(ctx, myAddr)
// 	privs, accNums, seqs = []crypto.PrivKey{myKey}, myAccount.GetAccountNumber(), myAccount.GetSequence()
// 	tx = types.NewTestTx(ctx, msgs, privs, []uint64{accNums}, []uint64{seqs}, fee)

// 	resp = app.DeliverTx(abci.RequestDeliverTx{Tx: app.MustEncodeTx(tx)})
// 	require.Equal(t, uint32(0), resp.Code, resp.Log)

// 	// verify second  proposal
// 	proposal, err = app.GroupKeeper.GetProposal(ctx, 2)
// 	require.NoError(t, err)
// 	assert.Equal(t, group.ProposalResultRejected, proposal.GetBase().Result, proposal.GetBase().Result.String())
// 	assert.Equal(t, group.ProposalStatusClosed, proposal.GetBase().Status, proposal.GetBase().Status.String())
// 	expTally = group.Tally{YesCount: sdk.ZeroDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.OneDec()}
// 	assert.Equal(t, expTally, proposal.GetBase().VoteState)
// }
