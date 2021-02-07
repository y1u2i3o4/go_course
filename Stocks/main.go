package main

import (
	"awesomeProject/Stocks/Binance"
	"awesomeProject/Stocks/Data"
	"awesomeProject/Stocks/Poloniex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	binance = iota
	poloniex
	sleepTime = time.Second*5
	batchSize = 1000
)



func main() {
	buffer := make(chan *Data.Trade, batchSize)
	binanceClient := Binance.GetBinanceClient()
	poloniexClient := Poloniex.GetClient()
	sqlClient := Data.GetSqlClient()

	go ReadPoloniex(buffer, poloniexClient)
	go ReadBinance(buffer, func() *Data.Trade { return fromBinanceTrade(binanceClient.ReadTrade())})

	http.HandleFunc(
		"/accounts",
		func (writer http.ResponseWriter, request *http.Request) {
			getAccounts(writer, request, sqlClient)
		})
	go http.ListenAndServe(":80", nil)
	Write(buffer, sqlClient)
}

func ReadPoloniex(ch chan *Data.Trade, client Poloniex.Client){
	for {
		trades, _ := client.GetTrades()
		if trades != nil{
			for _, trade := range trades{
				ch <- fromPoloniexTrade(&trade)
			}
		}
		time.Sleep(sleepTime)
	}
}

func getAccounts(writer http.ResponseWriter, request *http.Request, client Data.Client) {
	balances := client.GetAccountsBalance()
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(balances)
	if err != nil {
		log.Fatal("Не удалось записать ответ")
	}
}

func fromBinanceTrade(trade *Binance.Trade) *Data.Trade{
	price, _ := strconv.ParseFloat(trade.Price, 64)
	quantity, _ := strconv.ParseFloat(trade.Quantity, 64)
	return &Data.Trade{
		trade.Id,
		price,
		quantity,
		trade.BuyerId,
		trade.SellerId,
		binance,
	}
}

func fromPoloniexTrade(trade *Poloniex.Trade) *Data.Trade{
	price, _ := strconv.ParseFloat(trade.Amount, 64)
	quantity, _ := strconv.ParseFloat(trade.Rate, 64)
	return &Data.Trade{
		trade.GlobalTradeId,
		price,
		quantity,
		//Т.к. poloniex не возвращает ид клиента ид продавца не используем
		trade.TradeId,
		-1,
		poloniex,
	}
}

func ReadBinance(ch chan *Data.Trade, getter func () *Data.Trade){
	for {
		ch <- getter()
	}
}

func Write(ch chan *Data.Trade, sqlClient Data.Client){
	for {
		 trades := make([]*Data.Trade, batchSize)
		 for i := 0; i < batchSize; i++{
		 	trades[i] = <- ch
		 }
		 log.Printf("Вставка в бд %d записей", batchSize)
		 sqlClient.BulkInsert(trades)
	}
}
