package orderbook

import "github.com/regen-network/regen-ledger/types/math"

type OrderBook struct {
	creditOrderBooks map[string]*CreditOrderBook
	buyOrders        []*BuyOrder
	sellOrders       []*SellOrder
}

type BuyOrder struct {
	price            math.Dec
	amount           math.Dec
	creditOrderBooks map[string]*CreditOrderBook
}

type SellOrder struct {
	price           math.Dec
	amount          math.Dec
	creditOrderBook *CreditOrderBook
}

type CreditOrderBook struct {
	batchDenom string
	buyOrders  []*BuyOrder
	sellOrders []*SellOrder
}

func (*OrderBook) OnIssueCredit() {

}

func (*OrderBook) OnTagCredit() {

}

func (*OrderBook) OnBuyOrder() {

}

func (*OrderBook) OnCancelBuyOrder() {

}

func (*OrderBook) OnCancelSellOrder() {

}

func (*OrderBook) Execute() {

}
