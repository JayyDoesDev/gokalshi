package market

import (
	"github.com/jayydoesdev/gokalshi"
)

func GetEvents(keyID, keyPem string, q EventsQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/events", "GET", keyID, keyPem, true, q.ToMap())
}

func GetEvent(et, keyID, keyPem string, q EventQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/events/"+et, "GET", keyID, keyPem, true, q.ToMap())
}

func GetEventMeta(et, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/events/"+et+"/metadata", "GET", keyID, keyPem, true, map[string]string{})
}

func GetMarkets(keyID, keyPem string, q MarketsQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/markets", "GET", keyID, keyPem, true, q.ToMap())
}

func GetTrades(keyID, keyPem string, q TradesQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/markets/trades", "GET", keyID, keyPem, true, q.ToMap())
}

func GetMarket(t, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/markets/"+t, "GET", keyID, keyPem, true, map[string]string{})
}

func GetMarketOrderBook(t, keyID, keyPem string, q MarketOrderBookQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/markets/"+t+"/orderbook", "GET", keyID, keyPem, true, q.ToMap())
}

func GetSeriesList(keyID, keyPem string, q SeriesQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/series", "GET", keyID, keyPem, true, q.ToMap())
}

func GetSeries(st, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/series/"+st, "GET", keyID, keyPem, true, map[string]string{})
}

func GetMarketCandlesticks(st, t, keyID, keyPem string, q MarketCandlesticksQuery) ([]byte, error) {
	return gokalshi.Request[[]byte]("/series/"+st+"/markets"+t+"candlesticks", "GET", keyID, keyPem, true, q.ToMap())
}
