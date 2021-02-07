package Poloniex

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	url string = "https://poloniex.com/public?command=returnTradeHistory&currencyPair=BTC_ETH"
	responceSize = 200
)
type Client struct {
	http http.Client
}

func GetClient() Client{
	return Client{http.Client{}}
}

func (client Client) GetTrades() ([]Trade, error){

	responce, err := client.http.Get(url)
	if err != nil{
		log.Fatal("Не удалось получить данные с биржи Poloniex", err)
		return nil, err
	}
	defer responce.Body.Close()

	result := make([]Trade,0, responceSize)
	err = json.NewDecoder(responce.Body).Decode(&result)
	if err != nil{
		log.Fatal("Не удалось десереализовать данные о торгах ", err)
		return nil, err
	}
	return result, nil
}