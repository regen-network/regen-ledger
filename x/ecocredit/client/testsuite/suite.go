package testsuite

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	basketclient "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	marketplaceclient "github.com/regen-network/regen-ledger/x/ecocredit/client/marketplace"
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
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	coreclient "github.com/regen-network/regen-ledger/x/ecocredit/client"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
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
	classId            string
	projectId          string
	projectReferenceId string
	batchDenom         string
	basketDenom        string
	sellOrderIds       []uint64
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	return &IntegrationTestSuite{cfg: cfg}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	// set genesis values and params
	s.setupGenesis()

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	s.val = s.network.Validators[0]

	// set test accounts
	s.setupTestAccounts()

	// create test credit class
	s.classId, err = s.createClass(s.val.ClientCtx, &core.MsgCreateClass{
		Admin:            s.addr1.String(),
		Issuers:          []string{s.addr1.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &s.creditClassFee[0],
	})
	s.Require().NoError(err)

	// set test reference id
	s.projectReferenceId = "VCS-001"

	// create test project
	s.projectId, err = s.createProject(s.val.ClientCtx, &core.MsgCreateProject{
		Admin:        s.addr1.String(),
		ClassId:      s.classId,
		Metadata:     "metadata",
		Jurisdiction: "US-WA",
		ReferenceId:  s.projectReferenceId,
	})
	s.Require().NoError(err)

	startDate, err := types.ParseDate("start date", "2020-01-01")
	s.Require().NoError(err)

	endDate, err := types.ParseDate("expiration", "2021-01-01")
	s.Require().NoError(err)

	// create test credit batch
	s.batchDenom, err = s.createBatch(s.val.ClientCtx, &core.MsgCreateBatch{
		Issuer:    s.addr1.String(),
		ProjectId: s.projectId,
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
	s.Require().NoError(err)

	// create a basket and set test value
	s.basketDenom = s.createBasket(&basket.MsgCreate{
		Curator:          s.addr1.String(),
		Name:             "NCT",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{s.classId},
		Fee:              s.basketFee,
	})

	// put credits in basket (for testing basket balance)
	s.putInBasket(&basket.MsgPut{
		Owner:       s.addr1.String(),
		BasketDenom: s.basketDenom,
		Credits: []*basket.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     "1000",
			},
		},
	})

	askPrice := sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)

	// create test sell orders
	s.sellOrderIds, err = s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
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
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) setupGenesis() {
	// set up temporary mem db
	db := dbm.NewMemDB()
	defer db.Close()

	mdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	s.Require().NoError(err)

	coreStore, err := api.NewStateStore(mdb)
	s.Require().NoError(err)

	marketStore, err := marketApi.NewStateStore(mdb)
	s.Require().NoError(err)

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
	s.Require().NoError(err)

	// set allowed denoms
	s.allowedDenoms = append(s.allowedDenoms, sdk.DefaultBondDenom)

	s.creditTypeAbbrev = "C"

	// insert credit type
	err = coreStore.CreditTypeTable().Insert(ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         "carbon",
		Unit:         "metric ton CO2 equivalent",
		Precision:    6,
	})
	s.Require().NoError(err)

	// export genesis into target
	target := ormjson.NewRawMessageTarget()
	err = mdb.ExportJSON(ctx, target)
	s.Require().NoError(err)

	params := core.DefaultParams()
	s.creditClassFee = params.CreditClassFee
	s.basketFee = params.BasketFee

	// merge the params into the json target
	err = core.MergeParamsIntoTarget(s.cfg.Codec, &params, target)
	s.Require().NoError(err)

	// get raw json from target
	json, err := target.JSON()
	s.Require().NoError(err)

	// set the module genesis
	s.cfg.GenesisState[ecocredit.ModuleName] = json
}

func (s *IntegrationTestSuite) setupTestAccounts() {
	// create validator account
	info, _, err := s.val.ClientCtx.Keyring.NewMnemonic(
		"validator",
		keyring.English,
		sdk.FullFundraiserPath,
		keyring.DefaultBIP39Passphrase,
		hd.Secp256k1,
	)
	s.Require().NoError(err)

	// set primary account
	s.addr1 = s.val.Address

	// set secondary account
	s.addr2 = sdk.AccAddress(info.GetPubKey().Address())

	// fund the secondary account
	s.fundAccount(s.val.ClientCtx, s.val.Address, s.addr2, sdk.Coins{
		sdk.NewInt64Coin(s.cfg.BondDenom, 20000000000000000),
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
	_, err := banktestutil.MsgSendExec(
		clientCtx,
		from,
		to,
		coins,
		s.commonTxFlags()...,
	)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) createClass(clientCtx client.Context, msg *core.MsgCreateClass) (string, error) {
	var issuersStr string
	if len(msg.Issuers) == 1 {
		issuersStr = msg.Issuers[0]
	} else if len(msg.Issuers) > 1 {
		issuersStr = strings.Join(
			msg.Issuers,
			",",
		)
	}
	args := []string{
		issuersStr,
		msg.CreditTypeAbbrev,
		msg.Metadata,
		msg.Fee.String(),
	}
	flags := append(s.commonTxFlags(), fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Admin))
	args = append(args, flags...)

	cmd := coreclient.TxCreateClassCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)

	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateClass{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "class_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("class_id not found")
}

func (s *IntegrationTestSuite) createProject(clientCtx client.Context, msg *core.MsgCreateProject) (string, error) {
	cmd := coreclient.TxCreateProjectCmd()
	makeCreateProjectArgs := func(msg *core.MsgCreateProject, flags ...string) []string {
		args := []string{msg.ClassId, msg.Jurisdiction, msg.Metadata}
		return append(args, flags...)
	}

	referenceIdFlag := fmt.Sprintf("--reference-id=%s", msg.ReferenceId)
	flags := append(s.commonTxFlags(), fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Admin), referenceIdFlag)
	args := makeCreateProjectArgs(msg, flags...)

	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateProject{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "project_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("project_id not found")
}

func (s *IntegrationTestSuite) createBatch(clientCtx client.Context, msg *core.MsgCreateBatch) (string, error) {
	bytes, err := s.network.Validators[0].ClientCtx.Codec.MarshalJSON(msg)
	s.Require().NoError(err)

	batchJson := testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()

	args := []string{batchJson, fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Issuer)}
	args = append(args, s.commonTxFlags()...)
	cmd := coreclient.TxCreateBatchCmd()
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	s.Require().NoError(err)
	var res sdk.TxResponse
	s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	for _, e := range res.Logs[0].Events {
		if e.Type == proto.MessageName(&core.EventCreateBatch{}) {
			for _, attr := range e.Attributes {
				if attr.Key == "batch_denom" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("could not find batch_denom")
}

func (s *IntegrationTestSuite) createBasket(msg *basket.MsgCreate) (basketDenom string) {
	require := s.Require()

	allowedClasses := strings.Join(msg.AllowedClasses, ",")

	cmd := basketclient.TxCreateBasketCmd()
	args := []string{
		msg.Name,
		fmt.Sprintf("--%s=%s", basketclient.FlagCreditTypeAbbreviation, msg.CreditTypeAbbrev),
		fmt.Sprintf("--%s=%s", basketclient.FlagAllowedClasses, allowedClasses),
		fmt.Sprintf("--%s=%s", basketclient.FlagBasketFee, msg.Fee),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Curator),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
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

func (s *IntegrationTestSuite) putInBasket(msg *basket.MsgPut) {
	require := s.Require()

	// using json because array of BasketCredit is not a proto message
	bytes, err := json.Marshal(msg.Credits)
	require.NoError(err)

	creditsJson := testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()

	cmd := basketclient.TxPutInBasketCmd()
	args := []string{
		msg.BasketDenom,
		creditsJson,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Owner),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *IntegrationTestSuite) createSellOrder(clientCtx client.Context, msg *marketplace.MsgSell) ([]uint64, error) {
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
	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	orderIds := make([]uint64, 0, len(msg.Orders))
	for _, event := range res.Logs[0].Events {
		if event.Type == proto.MessageName(&marketplace.EventSell{}) {
			for _, attr := range event.Attributes {
				if attr.Key == "sell_order_id" {
					orderId, err := strconv.ParseUint(strings.Trim(attr.Value, "\""), 10, 64)
					require.NoError(err)
					orderIds = append(orderIds, orderId)
				}
			}
		}
	}

	if len(orderIds) == 0 {
		require.Fail("failed to find sell order ids in response")
	}

	return orderIds, nil
}
