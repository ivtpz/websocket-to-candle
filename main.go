package main

import (
	// "github.com/ivtpz/data-service/Mongo"
	"github.com/ivtpz/data-service/Mongo"
	"github.com/ivtpz/poloniex-socket-to-candle/Socket"
	// "gopkg.in/mgo.v2"
)

func main() {
	// Connect to data tracking DB
	// session, err := mgo.Dial("mongodb://mongo-0.mongo,mongo-1.mongo:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Connected to mongo cluster")
	// ds := &Mongo.DataStore{Session: session.Copy()}
	ds := Mongo.DataStore{}
	// defer ds.Session.Close()
	// ds.Session.SetMode(mgo.Monotonic, true)
	ds.AddCurrency("ETH")
	ds.AddCurrency("BTC")
	ds.AddCurrency("BCH")
	ds.AddCurrency("ETC")
	ds.AddCurrency("LTC")
	ds.AddCurrency("ZEC")
	ds.AddCurrency("DASH")
	ds.AddCurrency("XMR")
	// ds.EnsureIndex()

	socket := Socket.Connection{}
	data := Socket.Data{Db: &ds}
	socket.SetWsuri("wss://api2.poloniex.com")
	socket.SetData(&data)
	socket.Run()

}
