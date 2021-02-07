package Binance

import (
	"github.com/gorilla/websocket"
	"log"
)

const tradeStream string = "wss://stream.binance.com:9443/ws/btcusdt@trade"

type Client struct {
	conn *websocket.Conn
}

func GetBinanceClient() Client {
	conn, _, err := websocket.DefaultDialer.Dial(tradeStream, nil)

	if err != nil {
		log.Fatal("Не удалось подключиться к сокету ", err)
	}
	return Client{conn: conn}
}

func (client Client) ReadTrade() *Trade {
	trade := Trade{}
	client.conn.ReadJSON(&trade)
	return  &trade
}