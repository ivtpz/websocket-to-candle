package Socket

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/ivtpz/data-service/Mongo"
	// "time"
	// "net/http" - not sure if this is needed
)

type Connection struct {
	// Db *Mongo.DataStore
	connection *websocket.Conn
	wsuri      string
	markets    []string
	data       *Data
}

type Data struct {
	nextTime       int
	Db             *Mongo.DataStore
	currentCandles map[string]Mongo.Candle
}

func (data *Data) tick() {
	data.nextTime += 10
	// For each key in data.currentCandles
	// dataToStore := map[key]
	// market := key
	// Handle empty candles (get last candle in db)
	// go data.Db.AddCandle("poloniex_market_history", market + -10, dataToStore)
	// Reset candle with data.nextTime
}

func (conn *Connection) SetWsuri(uri string) {
	conn.wsuri = uri
}

func (conn *Connection) SetData(data *Data) {
	fmt.Println("setting data")
	fmt.Println(data)
	conn.data = data
	// Can I iterate over keys in a map?
	// data.markets := conn.markets
}

func (conn *Connection) AddMarket(market string) {
	// Do I need to handle nil values for markets?
	conn.markets = append(conn.markets, market)
}

func (conn *Connection) connect() {
	if conn.wsuri == "" {
		fmt.Println("uri not set, use SetWsuri before attempting to connect\n")
		return
	}
	fmt.Printf("attempting to connect to %s\n", conn.wsuri)
	dialer := websocket.Dialer{ /* Using defaults */ }
	wsConn, resp, err := dialer.Dial(conn.wsuri, nil)
	if err != nil {
		fmt.Println(resp.Body)
		panic(err)
	}
	conn.connection = wsConn
	fmt.Printf("Connected to %s\n", conn.wsuri)
}

type msg struct {
	Command string `json:"command"`
	Channel string `json:"channel"`
}

func (conn *Connection) listen() {
	c := conn.connection

	for _, curr := range conn.data.Db.GetCurrencies() {
		fmt.Println(curr)
		mkt := "USDT_" + curr
		err2 := c.WriteJSON(msg{"subscribe", mkt})
		if err2 != nil {
			panic(err2)
		}
	}

	for {
		messageType, p, err3 := c.ReadMessage()
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		/*
		 180: ZEC
		 121: BTC
		 122: DASH
		 173: ETC
		 123: LTC
		 126: XMR
		 149: ETH
		 191: BCH
		*/
		fmt.Printf("message type is %d\n", messageType)
		fmt.Printf("message is %s\n", p)
	}
	// Figure out how to subscribe to a market
	// Take message for market and add it to current candle for market
}

func (conn *Connection) Run() {
	// Connect to Poloniex websocket
	conn.connect()
	// Listen for trades
	conn.listen()
	// Set interval starting at time % 10 == 0
	// each interval, call conn.data.tick()
}
