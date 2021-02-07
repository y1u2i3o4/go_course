package Binance

type Trade struct {
	EventType string `json:"e"`
	EventTime int64 `json:"E"`
	Id int64 `json:"t"`
	Price string `json:"p"`
	Quantity string `json:"q"`
	BuyerId int64 `json:"b"`
	SellerId int64 `json:"a"`
	TradeTime int64 `json:"T"`
}