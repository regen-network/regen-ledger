package server

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

type TestSuite struct {
	suite.Suite

	ctx    sdk.Context
	goCtx  context.Context
	server Server
	ms     store.CommitMultiStore
}

func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()

	s.ms = store.NewCommitMultiStore(db)
	key := sdk.NewKVStoreKey("ecocredit")
	s.ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	s.NoError(s.ms.LoadLatestVersion())

	s.ctx = sdk.NewContext(s.ms, tmproto.Header{}, false, log.NewNopLogger())
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	s.server = NewServer(key)
}

func (s *TestSuite) TestScenario() {
	_, _, designer := testdata.KeyTestPubAddr()
	_, _, issuer1 := testdata.KeyTestPubAddr()
	_, _, issuer2 := testdata.KeyTestPubAddr()
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	_, _, addr3 := testdata.KeyTestPubAddr()

	// create class
	createClsRes, err := s.server.CreateClass(s.goCtx, &ecocredit.MsgCreateClassRequest{
		Designer: designer.String(),
		Issuers:  []string{issuer1.String(), issuer2.String()},
		Metadata: nil,
	})
	s.Require().NoError(err)
	s.Require().NotNil(createClsRes)

	clsId := createClsRes.ClassId
	s.Require().NotEmpty(clsId)

	// create batch
	t0, t1 := "10.37", "1007.3869"
	tSupply0 := "1017.7569"
	r0, r1 := "4.286", "10000.4589902"
	rSupply0 := "10004.7449902"

	createBatchRes, err := s.server.CreateBatch(s.goCtx, &ecocredit.MsgCreateBatchRequest{
		Issuer:  issuer1.String(),
		ClassId: clsId,
		Issuance: []*ecocredit.MsgCreateBatchRequest_BatchIssuance{
			{
				Recipient:      addr1.String(),
				TradeableUnits: t0,
				RetiredUnits:   r0,
			},
			{
				Recipient:      addr2.String(),
				TradeableUnits: t1,
				RetiredUnits:   r1,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(createBatchRes)

	batchDenom := createBatchRes.BatchDenom
	s.Require().NotEmpty(batchDenom)

	// query balances
	queryBalanceRes, err := s.server.Balance(s.goCtx, &ecocredit.QueryBalanceRequest{
		Account:    addr1.String(),
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t0, queryBalanceRes.TradeableUnits)
	s.Require().Equal(r0, queryBalanceRes.RetiredUnits)

	queryBalanceRes, err = s.server.Balance(s.goCtx, &ecocredit.QueryBalanceRequest{
		Account:    addr2.String(),
		BatchDenom: batchDenom,
	})
	s.Require().NoError(err)
	s.Require().NotNil(queryBalanceRes)
	s.Require().Equal(t1, queryBalanceRes.TradeableUnits)
	s.Require().Equal(r1, queryBalanceRes.RetiredUnits)

	// query supply
	querySupplyRes, err := s.server.Supply(s.goCtx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
	s.Require().NoError(err)
	s.Require().NotNil(querySupplyRes)
	s.Require().Equal(tSupply0, querySupplyRes.TradeableSupply)
	s.Require().Equal(rSupply0, querySupplyRes.RetiredSupply)

	// retire credits
	retireCases := []struct {
		name               string
		toRetire           string
		expectErr          bool
		expTradeable       string
		expRetired         string
		expTradeableSupply string
		expRetiredSupply   string
	}{
		{
			name:      "cannot retire more credits than are tradeable",
			toRetire:  "10.371",
			expectErr: true,
		},
		{
			name:      "can't use more than 7 decimal places",
			toRetire:  "10.00000001",
			expectErr: true,
		},
		{
			name:               "can retire a small amount of credits",
			toRetire:           "0.0001",
			expectErr:          false,
			expTradeable:       "10.3699",
			expRetired:         "4.2861",
			expTradeableSupply: "1017.7568",
			expRetiredSupply:   "10004.7450902",
		},
		{
			name:               "can retire more credits",
			toRetire:           "10",
			expectErr:          false,
			expTradeable:       "0.3699",
			expRetired:         "14.2861",
			expTradeableSupply: "1007.7568",
			expRetiredSupply:   "10014.7450902",
		},
		{
			name:               "can retire all credits",
			toRetire:           "0.3699",
			expectErr:          false,
			expTradeable:       "0",
			expRetired:         "14.656",
			expTradeableSupply: "1007.3869",
			expRetiredSupply:   "10015.1149902",
		},
		{
			name:      "can't retire any more credits",
			toRetire:  "1",
			expectErr: true,
		},
	}

	for _, tc := range retireCases {
		s.Run(tc.name, func() {
			ms := s.ms.CacheMultiStore()
			ctx := sdk.WrapSDKContext(sdk.NewContext(ms, tmproto.Header{}, false, log.NewNopLogger()))
			_, err := s.server.Retire(ctx, &ecocredit.MsgRetireRequest{
				Holder: addr1.String(),
				Credits: []*ecocredit.MsgRetireRequest_RetireUnits{
					{
						BatchDenom: batchDenom,
						Units:      tc.toRetire,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				ms.Write()

				// query balance
				queryBalanceRes, err = s.server.Balance(s.goCtx, &ecocredit.QueryBalanceRequest{
					Account:    addr1.String(),
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeable, queryBalanceRes.TradeableUnits)
				s.Require().Equal(tc.expRetired, queryBalanceRes.RetiredUnits)

				// query supply
				querySupplyRes, err = s.server.Supply(s.goCtx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradeableSupply, querySupplyRes.TradeableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}

	sendCases := []struct {
		name                  string
		sendTradeable         string
		sendRetired           string
		expectErr             bool
		expTradeableSender    string
		expRetiredSender      string
		expTradeableRecipient string
		expRetiredRecipient   string
		expTradeableSupply    string
		expRetiredSupply      string
	}{
		{
			name:          "can't send more tradeable than is tradeable",
			sendTradeable: "2000",
			sendRetired:   "10",
			expectErr:     true,
		},
		{
			name:          "can't send more retired than is tradeable",
			sendTradeable: "10",
			sendRetired:   "2000",
			expectErr:     true,
		},
		{
			name:                  "can send some",
			sendTradeable:         "10",
			sendRetired:           "20",
			expectErr:             false,
			expTradeableSender:    "977.3869",
			expRetiredSender:      "10000.4589902",
			expTradeableRecipient: "10",
			expRetiredRecipient:   "20",
			expTradeableSupply:    "987.3869",
			expRetiredSupply:      "10035.1149902",
		},
		{
			name:                  "can send all tradeable",
			sendTradeable:         "77.3869",
			sendRetired:           "900",
			expectErr:             false,
			expTradeableSender:    "0",
			expRetiredSender:      "10000.4589902",
			expTradeableRecipient: "87.3869",
			expRetiredRecipient:   "920",
			expTradeableSupply:    "87.3869",
			expRetiredSupply:      "10935.1149902",
		},
		{
			name:          "can't send any more",
			sendTradeable: "1",
			sendRetired:   "1",
			expectErr:     true,
		},
	}

	for _, tc := range sendCases {
		s.Run(tc.name, func() {
			ms := s.ms.CacheMultiStore()
			ctx := sdk.WrapSDKContext(sdk.NewContext(ms, tmproto.Header{}, false, log.NewNopLogger()))
			_, err := s.server.Send(ctx, &ecocredit.MsgSendRequest{
				Sender:    addr2.String(),
				Recipient: addr3.String(),
				Credits: []*ecocredit.MsgSendRequest_SendUnits{
					{
						BatchDenom:     batchDenom,
						TradeableUnits: tc.sendTradeable,
						RetiredUnits:   tc.sendRetired,
					},
				},
			})

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				ms.Write()

				// query sender balance
				queryBalanceRes, err = s.server.Balance(s.goCtx, &ecocredit.QueryBalanceRequest{
					Account:    addr2.String(),
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeableSender, queryBalanceRes.TradeableUnits)
				s.Require().Equal(tc.expRetiredSender, queryBalanceRes.RetiredUnits)

				// query recipient balance
				queryBalanceRes, err = s.server.Balance(s.goCtx, &ecocredit.QueryBalanceRequest{
					Account:    addr3.String(),
					BatchDenom: batchDenom,
				})
				s.Require().NoError(err)
				s.Require().NotNil(queryBalanceRes)
				s.Require().Equal(tc.expTradeableRecipient, queryBalanceRes.TradeableUnits)
				s.Require().Equal(tc.expRetiredRecipient, queryBalanceRes.RetiredUnits)

				// query supply
				querySupplyRes, err = s.server.Supply(s.goCtx, &ecocredit.QuerySupplyRequest{BatchDenom: batchDenom})
				s.Require().NoError(err)
				s.Require().NotNil(querySupplyRes)
				s.Require().Equal(tc.expTradeableSupply, querySupplyRes.TradeableSupply)
				s.Require().Equal(tc.expRetiredSupply, querySupplyRes.RetiredSupply)
			}
		})
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
