package Poloniex

type Trade struct {
	GlobalTradeId int64 `json:"globalTradeID"`
	TradeId int64 `json:"tradeID"`
	Rate string `json:"rate"`
	Amount string `json:"amount"`
	Total string `json:"total"`
}
