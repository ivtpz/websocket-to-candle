FROM golang:latest

RUN go get github.com/gorilla/websocket
RUN go get gopkg.in/mgo.v2

# To use this as base for poloniex-socket-to-candle build:
# docker build -t poloniex-socket-to-candle/golang .