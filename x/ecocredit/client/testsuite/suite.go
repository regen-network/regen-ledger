package testsuite

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/abci/types"

	sdkbase "cosmossdk.io/api/cosmos/base/v1beta1"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm/model/ormdb"
	"github.com/regen-network/regen-ledger/orm/model/ormtable"
	"github.com/regen-network/regen-ledger/orm/types/ormjson"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	baseclient "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/client"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	basketclient "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/client"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v4/genesis"
	marketclient "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/client"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
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
	creditClassFee     *sdk.Coin
	basketFee          sdk.Coins
	creditTypeAbbrev   string
	allowedDenoms      []string
	bridgeChain        string
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
	s.cfg.StakingTokens = sdkmath.NewInt(900000000)
	s.cfg.AccountTokens = sdkmath.NewInt(9000000000)
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	require.NoError(err)

	_, err = s.network.WaitForHeight(1)
	require.NoError(err)

	s.val = s.network.Validators[0]

	// set test accounts
	s.setupTestAccounts()
	require.NoError(s.network.WaitForNextBlock())
	// create test credit class
	s.classID = s.createClass(s.val.ClientCtx, &basetypes.MsgCreateClass{
		Admin:            s.addr1.String(),
		Issuers:          []string{s.addr1.String()},
		Metadata:         "metadata",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              s.creditClassFee,
	})

	// set test reference id
	s.projectReferenceID = "VCS-001"

	// create test project
	s.projectID = s.createProject(s.val.ClientCtx, &basetypes.MsgCreateProject{
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
	s.batchDenom = s.createBatch(s.val.ClientCtx, &basetypes.MsgCreateBatch{
		Issuer:    s.addr1.String(),
		ProjectId: s.projectID,
		Issuance: []*basetypes.BatchIssuance{
			{
				Recipient:              s.addr1.String(),
				TradableAmount:         "10000",
				RetirementJurisdiction: "US-WA",
			},
		},
		Metadata:  "metadata",
		StartDate: &startDate,
		EndDate:   &endDate,
		Open:      true,
	})

	// create a basket and set test value
	s.basketDenom = s.createBasket(s.val.ClientCtx, &baskettypes.MsgCreate{
		Curator:          s.addr1.String(),
		Name:             "NCT",
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{s.classID},
		Fee:              s.basketFee,
	})

	// put credits in basket (for testing basket balance)
	s.putInBasket(s.val.ClientCtx, &baskettypes.MsgPut{
		Owner:       s.addr1.String(),
		BasketDenom: s.basketDenom,
		Credits: []*baskettypes.BasketCredit{
			{
				BatchDenom: s.batchDenom,
				Amount:     "1000",
			},
		},
	})

	askPrice := sdk.NewInt64Coin(s.allowedDenoms[0], 10)

	// create sell orders with first test account and set test values
	sellOrderIDs := s.createSellOrder(s.val.ClientCtx, &markettypes.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*markettypes.MsgSell_Order{
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

	db := ormutil.NewStoreAdapter(dbm.NewMemDB())
	mdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(err)

	baseStore, err := baseapi.NewStateStore(mdb)
	require.NoError(err)

	marketStore, err := marketapi.NewStateStore(mdb)
	require.NoError(err)

	basketStore, err := basketapi.NewStateStore(mdb)
	require.NoError(err)

	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ctx := ormtable.WrapContextDefault(backend)

	// add basket fees
	err = basketStore.BasketFeeTable().Save(ctx, &basketapi.BasketFee{
		Fee: &sdkbase.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: basetypes.DefaultBasketFee.String(),
		},
	})
	require.NoError(err)

	// insert allowed denom
	err = marketStore.AllowedDenomTable().Insert(ctx, &marketapi.AllowedDenom{
		BankDenom:    sdk.DefaultBondDenom,
		DisplayDenom: sdk.DefaultBondDenom,
	})
	require.NoError(err)

	// set allowed denoms
	s.allowedDenoms = append(s.allowedDenoms, sdk.DefaultBondDenom)

	// set credit type abbreviation
	s.creditTypeAbbrev = "C"

	// insert credit type
	err = baseStore.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         "carbon",
		Unit:         "metric ton CO2 equivalent",
		Precision:    6,
	})
	require.NoError(err)

	// set credit class fees
	err = baseStore.ClassFeeTable().Save(ctx, &baseapi.ClassFee{
		Fee: &sdkbase.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: basetypes.DefaultClassFee.String(),
		},
	})
	require.NoError(err)

	// set credit class allow list
	err = baseStore.ClassCreatorAllowlistTable().Save(ctx, &baseapi.ClassCreatorAllowlist{
		Enabled: false,
	})
	require.NoError(err)

	// set allowed credit class creators
	err = baseStore.AllowedClassCreatorTable().Insert(ctx, &baseapi.AllowedClassCreator{
		Address: sdk.AccAddress("issuer1"),
	})
	require.NoError(err)
	err = baseStore.AllowedClassCreatorTable().Insert(ctx, &baseapi.AllowedClassCreator{
		Address: sdk.AccAddress("issuer2"),
	})
	require.NoError(err)

	s.bridgeChain = "polygon"

	// set allowed bridge chain
	err = baseStore.AllowedBridgeChainTable().Insert(ctx, &baseapi.AllowedBridgeChain{
		ChainName: s.bridgeChain,
	})
	require.NoError(err)

	// set batch contract for bridge testing
	err = baseStore.BatchContractTable().Insert(ctx, &baseapi.BatchContract{
		BatchKey: 1,
		ClassKey: 1,
		Contract: "0x0000000000000000000000000000000000000000",
	})
	require.NoError(err)

	// export genesis into target
	target := ormjson.NewRawMessageTarget()
	err = mdb.ExportJSON(ctx, target)
	require.NoError(err)

	// set credit class and basket fees
	s.creditClassFee = genesis.DefaultClassFee().Fee
	s.basketFee = sdk.NewCoins(*genesis.DefaultBasketFee().Fee)

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

	s.fundAccount(s.val.ClientCtx, s.addr1, s.addr2, sdk.Coins{
		sdk.NewInt64Coin(s.cfg.BondDenom, 1000),
	})
}

func (s *IntegrationTestSuite) commonTxFlags() []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdkmath.NewInt(10))).String()),
	}
}

func (s *IntegrationTestSuite) fundAccount(clientCtx client.Context, from, to sdk.AccAddress, coins sdk.Coins) {
	require := s.Require()

	out, err := cli.MsgSendExec(clientCtx, from, to, coins, addresscodec.NewBech32Codec("regen"), s.commonTxFlags()...)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
	require.NoError(s.network.WaitForNextBlock())
}

func (s *IntegrationTestSuite) createClass(clientCtx client.Context, msg *basetypes.MsgCreateClass) (classID string) {
	require := s.Require()

	// Step 1: Execute the transaction to create a class
	cmd := baseclient.TxCreateClassCmd()
	args := []string{
		strings.Join(msg.Issuers, ","),
		msg.CreditTypeAbbrev,
		msg.Metadata,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Admin),
		fmt.Sprintf("--%s=%s", baseclient.FlagClassFee, msg.Fee.String()),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync), // Use `sync` mode
	}

	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)

	// Parse transaction response
	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)

	// Step 2: Wait for the transaction to be included in a block
	require.NoError(s.network.WaitForNextBlock())

	// Step 3: Query the transaction by hash to retrieve logs
	queryRes, err := cli.GetTxResponse(s.network, clientCtx, res.TxHash)
	require.NoError(err)

	// Step 4: Parse events to extract the class ID
	classID, err = getClassIDFromEvents(queryRes.Events)
	require.NoError(err)

	return classID
}

func getClassIDFromEvents(events []abcitypes.Event) (string, error) {
	for _, ev := range events {
		if ev.Type == "regen.ecocredit.v1.EventCreateClass" {
			for _, attr := range ev.Attributes {
				if attr.Key == "class_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("class_id not found in events")
}

func (s *IntegrationTestSuite) createProject(clientCtx client.Context, msg *basetypes.MsgCreateProject) (projectID string) {
	require := s.Require()

	cmd := baseclient.TxCreateProjectCmd()
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

	require.NoError(s.network.WaitForNextBlock())

	queryRes, err := cli.GetTxResponse(s.network, clientCtx, res.TxHash)
	require.NoError(err)

	projectID, err = getProjectIDFromEvents(queryRes.Events)
	require.NoError(err)

	return projectID
}

func getProjectIDFromEvents(events []abcitypes.Event) (string, error) {
	for _, ev := range events {
		if ev.Type == "regen.ecocredit.v1.EventCreateProject" {
			for _, attr := range ev.Attributes {
				if attr.Key == "project_id" {
					return strings.Trim(attr.Value, "\""), nil
				}
			}
		}
	}
	return "", fmt.Errorf("project_id not found in events")
}

func (s *IntegrationTestSuite) createBatch(clientCtx client.Context, msg *basetypes.MsgCreateBatch) (batchDenom string) {
	require := s.Require()

	bz, err := clientCtx.Codec.MarshalJSON(msg)
	require.NoError(err)

	jsonFile := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	cmd := baseclient.TxCreateBatchCmd()
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
	require.NoError(s.network.WaitForNextBlock())

	queryRes, err := cli.GetTxResponse(s.network, clientCtx, res.TxHash)
	require.NoError(err)

	for _, ev := range queryRes.Events {
		if ev.Type == "regen.ecocredit.v1.EventCreateBatch" {
			for _, attr := range ev.Attributes {
				if attr.Key == "batch_denom" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find batch denom in response")

	return ""
}

func (s *IntegrationTestSuite) createBasket(clientCtx client.Context, msg *baskettypes.MsgCreate) (basketDenom string) {
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
	require.NoError(s.network.WaitForNextBlock())

	queryRes, err := cli.GetTxResponse(s.network, clientCtx, res.TxHash)
	require.NoError(err)

	for _, ev := range queryRes.Events {
		if ev.Type == "regen.ecocredit.basket.v1.EventCreate" {
			for _, attr := range ev.Attributes {
				if attr.Key == "basket_denom" {
					return strings.Trim(attr.Value, "\"")
				}
			}
		}
	}

	require.Fail("failed to find basket denom in response")

	return ""
}

func (s *IntegrationTestSuite) putInBasket(clientCtx client.Context, msg *baskettypes.MsgPut) {
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
	require.NoError(s.network.WaitForNextBlock())
	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *IntegrationTestSuite) createSellOrder(clientCtx client.Context, msg *markettypes.MsgSell) (sellOrderIDs []uint64) {
	require := s.Require()

	// using json package because array is not a proto message
	bz, err := json.Marshal(msg.Orders)
	require.NoError(err)

	jsonFile := testutil.WriteToNewTempFile(s.T(), string(bz)).Name()

	cmd := marketclient.TxSellCmd()
	args := []string{
		jsonFile,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, msg.Seller),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(err)
	require.NoError(s.network.WaitForNextBlock())

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
	queryRes, err := cli.GetTxResponse(s.network, clientCtx, res.TxHash)
	require.NoError(err)
	orderIDs := make([]uint64, 0, len(msg.Orders))

	for _, ev := range queryRes.Events {
		if ev.Type == "regen.ecocredit.marketplace.v1.EventSell" {
			for _, attr := range ev.Attributes {
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
