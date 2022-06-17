package testsuite

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/gogo/protobuf/proto"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"

	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil/cli"
	"github.com/regen-network/regen-ledger/types/testutil/network"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basketclient "github.com/regen-network/regen-ledger/x/ecocredit/client/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	val     *network.Validator

	// test accounts
	addr1 sdk.AccAddress
	addr  sdk.AccAddress // TODO: addr2 (#922 / #1042)

	// test values
	creditTypeAbbrev   string
	allowedDenoms      []string
	classId            string
	projectId          string
	projectReferenceId string
	batchDenom         string
	basketDenom        string
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

	// set reference id used when creating a project
	s.projectReferenceId = "VCS-001"

	// create a class, project, and batch with first test account and set test values
	s.classId, s.projectId, s.batchDenom = s.createClassProjectBatch(s.val.ClientCtx, s.addr1.String())

	// create a basket and set test value
	s.basketDenom = s.createBasket("NCT", s.creditTypeAbbrev, s.classId, s.addr1.String())

	// credits to put in basket
	credits := basket.BasketCredit{
		BatchDenom: s.batchDenom,
		Amount:     "1",
	}

	// put credits in basket (for testing basket balance)
	s.putInBasket(s.basketDenom, credits, s.addr1.String())
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

	// merge the params into the json target
	params := core.DefaultParams()
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

	// create secondary account
	account := sdk.AccAddress(info.GetPubKey().Address())

	// fund the secondary account
	s.fundAccount(s.val.ClientCtx, s.val.Address, account, sdk.Coins{
		sdk.NewInt64Coin(s.cfg.BondDenom, 20000000000000000),
	})

	// set test accounts
	s.addr1 = s.val.Address
	s.addr = account // TODO: addr2 (#922 / #1042)
}

func (s *IntegrationTestSuite) createBasket(name, creditTypeAbbrev, classId string, curator string) (basketDenom string) {
	require := s.Require()

	cmd := basketclient.TxCreateBasket()
	args := []string{
		name,
		fmt.Sprintf("--%s=%s", basketclient.FlagCreditTypeAbbreviation, creditTypeAbbrev),
		fmt.Sprintf("--%s=%s", basketclient.FlagAllowedClasses, classId),
		fmt.Sprintf("--%s=%s", basketclient.FlagBasketFee, "20000000stake"),
		makeFlagFrom(curator),
	}
	args = append(args, s.commonTxFlags()...)
	out, err := cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))

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

func (s *IntegrationTestSuite) putInBasket(basketDenom string, credits basket.BasketCredit, owner string) {
	require := s.Require()

	// using json because array of BasketCredit is not a proto message
	bytes, err := json.Marshal([]basket.BasketCredit{credits})
	require.NoError(err)

	creditsJson := testutil.WriteToNewTempFile(s.T(), string(bytes)).Name()

	cmd := basketclient.TxPutInBasket()
	args := []string{
		basketDenom,
		creditsJson,
		makeFlagFrom(owner),
	}
	args = append(args, s.commonTxFlags()...)
	_, err = cli.ExecTestCLICmd(s.val.ClientCtx, cmd, args)
	require.NoError(err)
}
