FROM poloniex-socket-to-candle/golang:latest as build
ADD . /go/src/github.com/ivtpz/poloniex-socket-to-candle/
WORKDIR /go/src/github.com/ivtpz/poloniex-socket-to-candle

RUN cd Socket && go install
RUN go build -o main .
CMD ["/go/src/github.com/ivtpz/poloniex-socket-to-candle/main"]