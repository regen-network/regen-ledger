package testsuite

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	basketclient "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/genesis"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	val     *network.Validator

	// test accounts
	addr1 sdk.AccAddress
	addr2 sdk.AccAddress

	// test values
	creditClassFee     sdk.Coins
	basketFee          sdk.Coins
	creditTypeAbbrev   string
	allowedDenoms      []string
	classID            string
	projectID          string
	projectReferenceID string
	batchDenom         string
	basketDenom        string
	sellOrderID        uint64
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	require := s.Require()

	s.T().Log("setting up integration test suite")

	// set genesis values and params
	s.setupGenesis()

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	require.NoError(err)

	_, err = s.network.WaitForHeight(1)
	require.NoError(err)

	s.val = s.network.Validators[0]

	// set test accounts
	s.setupTestAccounts()

	// create test credit class
	s.classID = s.createClass(s.val.ClientCtx, &core.MsgCreateClass{
		Admin:            s.addr1.String(),
		Issuers:          []string{s.addr1.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &s.creditClassFee[0],
	})

	// set test reference id
	s.projectReferenceID = "VCS-001"

	// create test project
	s.projectID = s.createProject(s.val.ClientCtx, &core.MsgCreateProject{
		Admin:        s.addr1.String(),
		ClassId:      s.classID,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  s.projectReferenceID,
	})

	startDate, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(err)

	endDate, err := types.ParseDate("expiration", "2021-01-01")
	require.NoError(err)

	// create test credit batch
	s.batchDenom = s.createBatch(s.val.ClientCtx, &core.MsgCreateBatch{
		Issuer:    s.addr1.String(),
		ProjectId: s.projectID,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:              s.addr1.String(),
				TradableAmount:         "10000",
				RetirementJurisdiction: "US-WA",
			},
		},
		Metadata:  "metadata",
		StartDate: &startDate,
		EndDate:   &endDate,
	})

	// create a basket and set test value
	s.basketDenom = s.createBasket(s.val.ClientCtx, &basket.MsgCreate{
		Curator:          s.addr1.String(),
		Name:             "NCT",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{s.classID},
		Fee:              s.basketFee,
	})

	// put credits in basket (for testing basket balance)
	s.putInBasket(s.val.ClientCtx, &basket.MsgPut{
		Owner:       s.addr1.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     "1000",
			},
		},
	})

	askPrice := sdk.NewInt64Coin(s.allowedDenoms[0], 10)

	// create sell orders with first test account and set test values
	sellOrderIDs := s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom:        s.batchDenom,
				Quantity:          "1000",
				AskPrice:          &askPrice,
				DisableAutoRetire: true,
			},
		},
	})

	s.sellOrderID = sellOrderIDs[0]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.Require().NoError(s.network.WaitForNextBlock())
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) setupGenesis() {
	require := s.Require()

	// set up temporary mem db
	db := dbm.NewMemDB()

	mdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(err)

	coreStore, err := api.NewStateStore(mdb)
	require.NoError(err)

	marketStore, err := marketApi.NewStateStore(mdb)
	require.NoError(err)

	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ctx := ormtable.WrapContextDefault(backend)

	// insert allowed denom
	err = marketStore.AllowedDenomTable().Insert(ctx, &marketApi.AllowedDenom{
		BankDenom:    sdk.DefaultBondDenom,
		DisplayDenom: sdk.DefaultBondDenom,
	})
	require.NoError(err)

	// set allowed denoms
	s.allowedDenoms = append(s.allowedDenoms, sdk.DefaultBondDenom)

	// set credit type abbreviation
	s.creditTypeAbbrev = "C"

	// insert credit type
	err = coreStore.CreditTypeTable().Insert(ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         "carbon",
		Unit:         "metric ton CO2 equivalent",
		Precision:    6,
	})
	require.NoError(err)

	// export genesis into target
	target := ormjson.NewRawMessageTarget()
	err = mdb.ExportJSON(ctx, target)
	require.NoError(err)

	params := core.DefaultParams()

	// set credit class and basket fees
	s.creditClassFee = params.CreditClassFee
	s.basketFee = params.BasketFee

	// merge the params into the json target
	err = genesis.MergeParamsIntoTarget(s.cfg.Codec, &params, target)
	require.NoError(err)

	// get raw json from target
	json, err := target.JSON()
	require.NoError(err)

	// set the module genesis
	s.cfg.GenesisState[ecocredit.ModuleName] = json
}

func (s *IntegrationTestSuite) setupTestAccounts() {
	// create secondary account
	info, _, err := s.val.ClientCtx.Keyring.NewMnemonic(
		"addr2",
		keyring.English,
		sdk.FullFundraiserPath,
		keyring.DefaultBIP39Passphrase,
		hd.Secp256k1,
	)
	s.Require().NoError(err)

	// set primary account
	s.addr1 = s.val.Address

	// set secondary account
	pk, err := info.GetPubKey()
	s.Require().NoError(err)
	s.addr2 = sdk.AccAddress(pk.Address())

	// fund secondary account
	s.fundAccount(s.val.ClientCtx, s.addr1, s.addr2, sdk.Coins{
		sdk.NewInt64Coin(s.cfg.BondDenom, 100000000),
	})
}

func (s *IntegrationTestSuite) commonTxFlags() []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}
}

func (s *IntegrationTestSuite) fundAccount(clientCtx client.Context, from, to sdk.AccAddress, coins sdk.Coins) {
	require := s.Require()

	out, err := banktestutil.MsgSendExec(
		clientCtx,
		from,
		to,
		coins,
		s.commonTxFlags()...,
	)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *IntegrationTestSuite) createClass(clientCtx client.Context, msg *core.MsgCreateClass) (classID string) {
	require := s.Require()

	cmd := coreclient.TxCreateClassCmd()
	args := []string{
		strings.Join(msg.Issuers, ","),
		msg.CreditTypeAbbrev,
		msg.Metadata,
		msg.Fee.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Admin),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateClass{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "class_id" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find class id in response")

	return ""
}

func (s *IntegrationTestSuite) createProject(clientCtx client.Context, msg *core.MsgCreateProject) (projectID string) {
	require := s.Require()

	cmd := coreclient.TxCreateProjectCmd()
	args := []string{
		msg.ClassId,
		msg.Jurisdiction,
		msg.Metadata,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Admin),
		fmt.Sprintf("--reference-id=%s", msg.ReferenceId),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateProject{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "project_id" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find project id in response")

	return ""
}

func (s *IntegrationTestSuite) createBatch(clientCtx client.Context, msg *core.MsgCreateBatch) (batchDenom string) {
	require := s.Require()

	bz, err := clientCtx.Codec.MarshalJSON(msg)
	require.NoError(err)

	jsonFile := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	cmd := coreclient.TxCreateBatchCmd()
	args := []string{
		jsonFile,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Issuer),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateBatch{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "batch_denom" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find batch denom in response")

	return ""
}

func (s *IntegrationTestSuite) createBasket(clientCtx client.Context, msg *basket.MsgCreate) (basketDenom string) {
	require := s.Require()

	cmd := basketclient.TxCreateBasketCmd()
	args := []string{
		msg.Name,
		fmt.Sprintf("--%s=%s", basketclient.FlagCreditTypeAbbrev, msg.CreditTypeAbbrev),
		fmt.Sprintf("--%s=%s", basketclient.FlagAllowedClasses, strings.Join(msg.AllowedClasses, ",")),
		fmt.Sprintf("--%s=%s", basketclient.FlagBasketFee, msg.Fee),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Curator),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	for _, event := range res.Logs[0].Events {
		if event.Type == proto.MessageName(&basket.EventCreate{}) {
			for _, attr := range event.Attributes {
				if attr.Key == "basket_denom" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find basket denom in response")

	return ""
}

func (s *IntegrationTestSuite) putInBasket(clientCtx client.Context, msg *basket.MsgPut) {
	require := s.Require()

	// using json because array of BasketCredit is not a proto message
	bz, err := json.Marshal(msg.Credits)
	require.NoError(err)

	jsonFile := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	cmd := basketclient.TxPutInBasketCmd()
	args := []string{
		msg.BasketDenom,
		jsonFile,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Owner),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *IntegrationTestSuite) createSellOrder(clientCtx client.Context, msg *marketplace.MsgSell) (sellOrderIDs []uint64) {
	require := s.Require()

	// using json package because array is not a proto message
	bz, err := json.Marshal(msg.Orders)
	require.NoError(err)

	jsonFile := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	cmd := marketplaceclient.TxSellCmd()
	args := []string{
		jsonFile,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Seller),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	orderIDs := make([]uint64, 0, len(msg.Orders))
	for _, event := range res.Logs[0].Events {
		if event.Type == proto.MessageName(&marketplace.EventSell{}) {
			for _, attr := range event.Attributes {
				if attr.Key == "sell_order_id" {
					orderID, err := strconv.ParseUint(strings.Trim(attr.Value, "\""), 10, 64)
					require.NoError(err)
					orderIDs = append(orderIDs, orderID)
				}
			}
		}
	}

	if len(orderIDs) == 0 {
		require.Fail("failed to find sell order id(s) in response")
	}

	return orderIDs
}
