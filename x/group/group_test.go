package group

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"gitlab.com/regen-network/regen-ledger/util"
	"reflect"
	"testing"
)

var cdc *codec.Codec
var ctx sdk.Context
var keeper Keeper

func setupTestInput() {
	db := dbm.NewMemDB()

	cdc = codec.New()

	groupKey := sdk.NewKVStoreKey("groupKey")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(groupKey, sdk.StoreTypeIAVL, db)
	_ = ms.LoadLatestVersion()

	keeper = NewKeeper(groupKey, cdc)
	ctx = sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
}

var privKey secp256k1.PrivKeySecp256k1

var pubKey crypto.PubKey

var group Group

var groupId sdk.AccAddress

func aPublicKeyAddress() error {
	privKey = secp256k1.GenPrivKey()
	pubKey = privKey.PubKey()
	return nil
}

func aUserCreatesAGroupWithThatAddressAndADecisionThresholdOf(t int64) error {
	mem := Member{Address: sdk.AccAddress(pubKey.Address())}
	mem.Weight.SetInt64(1)
	group = Group{
		Members: []Member{mem},
	}
	group.DecisionThreshold.SetInt64(t)
	groupId = keeper.CreateGroup(ctx, group)
	return nil
}

func theyShouldGetANewGroupAddressBack() error {
	if groupId == nil || len(groupId) <= 0 {
		return fmt.Errorf("group ID was empty")
	}
	return nil
}

func beAbleToRetrieveTheGroupDetailsWithThatAddress() error {
	groupRetrieved, err := keeper.GetGroupInfo(ctx, groupId)
	if err != nil {
		return fmt.Errorf("error retrieving group info %+v", err)
	}
	if reflect.DeepEqual(group, groupRetrieved) {
		return fmt.Errorf("retrieved group differs from committed group, expected %+v, got %+v",
			group, groupRetrieved)
	}
	return nil
}

func TestMain(m *testing.M) {
	util.GodogMain(m, "group", FeatureContext)
}

func FeatureContext(s *godog.Suite) {
	s.BeforeFeature(func(*gherkin.Feature) {
		setupTestInput()
	})
	s.Step(`^a public key address$`, aPublicKeyAddress)
	s.Step(`^they should get a new group address back$`, theyShouldGetANewGroupAddressBack)
	s.Step(`^a user creates a group with that address and a decision threshold of (\d+)$`, aUserCreatesAGroupWithThatAddressAndADecisionThresholdOf)
	s.Step(`^be able to retrieve the group details with that address$`, beAbleToRetrieveTheGroupDetailsWithThatAddress)
}
