package market

import (
	"github.com/jayydoesdev/gokalshi"
)

type Events struct {
	Cursor string  `json:"cursor"`
	Events []Event `json:"events"`
}

type Event struct {
	AvailableOnBrokers   bool     `json:"available_on_brokers"`
	Category             string   `json:"category"`
	CollateralReturnType string   `json:"collateral_return_type"`
	EventTicker          string   `json:"event_ticker"`
	Markets              []Market `json:"markets"`

	MutuallyExclusive   bool        `json:"mutually_exclusive"`
	PriceLevelStructure string      `json:"price_level_structure"`
	SeriesTicker        string      `json:"series_ticker"`
	StrikeDate          interface{} `json:"strike_date"`
	StrikePeriod        string      `json:"strike_period"`
	SubTitle            string      `json:"sub_title"`
	Title               string      `json:"title"`
}

type Market struct {
	Ticker               string `json:"ticker"`
	EventTicker          string `json:"event_ticker"`
	MarketType           string `json:"market_type"`
	Title                string `json:"title"`
	Subtitle             string `json:"subtitle"`
	YesSubTitle          string `json:"yes_sub_title"`
	NoSubTitle           string `json:"no_sub_title"`
	OpenTime             string `json:"open_time"`
	CloseTime            string `json:"close_time"`
	ExpirationTime       string `json:"expiration_time"`
	LatestExpirationTime string `json:"latest_expiration_time"`

	SettlementTimerSeconds int    `json:"settlement_timer_seconds"`
	Status                 string `json:"status"`
	ResponsePriceUnits     string `json:"response_price_units"`
	NotionalValue          int    `json:"notional_value"`
	NotionalValueDollars   string `json:"notional_value_dollars"`
	TickSize               int    `json:"tick_size"`
	YesBid                 int    `json:"yes_bid"`
	YesBidDollars          string `json:"yes_bid_dollars"`
	YesAsk                 int    `json:"yes_ask"`
	YesAskDollars          string `json:"yes_ask_dollars"`
	NoBid                  int    `json:"no_bid"`
	NoBidDollars           string `json:"no_bid_dollars"`
	NoAsk                  int    `json:"no_ask"`
	NoAskDollars           string `json:"no_ask_dollars"`
	LastPrice              int    `json:"last_price"`
	LastPriceDollars       string `json:"last_price_dollars"`
	PreviousYesBid         int    `json:"previous_yes_bid"`
	PreviousYesBidDollars  string `json:"previous_yes_bid_dollars"`
	PreviousYesAsk         int    `json:"previous_yes_ask"`
	PreviousYesAskDollars  string `json:"previous_yes_ask_dollars"`
	PreviousPrice          int    `json:"previous_price"`
	PreviousPriceDollars   string `json:"previous_price_dollars"`
	Volume                 int    `json:"volume"`
	Volume24h              int    `json:"volume_24h"`
	Liquidity              int    `json:"liquidity"`
	LiquidityDollars       string `json:"liquidity_dollars"`
	OpenInterest           int    `json:"open_interest"`
	Result                 string `json:"result"`
	CanCloseEarly          bool   `json:"can_close_early"`
	ExpirationValue        string `json:"expiration_value"`
	Category               string `json:"category"`
	RiskLimitCents         int    `json:"risk_limit_cents"`
	RulesPrimary           string `json:"rules_primary"`
	RulesSecondary         string `json:"rules_secondary"`
	SettlementValue        int    `json:"settlement_value"`
	SettlementValueDollars string `json:"settlement_value_dollars"`
}

type Trades struct {
	Cursor string  `json:"cursor"`
	Trades []Trade `json:"trades"`
}

type Trade struct {
	Count           int    `json:"count"`
	CreateTime      string `json:"created_time"`
	NoPrice         int    `json:"no_price"`
	NoPriceDollars  string `json:"no_price_dollars"`
	TakerSide       string `json:"taker_side"`
	Ticker          string `json:"ticker"`
	TradeId         string `json:"trade_id"`
	YesPrice        int    `json:"yes_price"`
	YesPriceDollars string `json:"yes_price_dollars"`
}

type EventsQuery struct {
	Limit             int
	Cursor            string
	WithNestedMarkets bool
	Status            string
	SeriesTracker     string
	MinCloseTs        int
}

func (q EventsQuery) ToMap() map[string]string {
	return map[string]string{
		"limit":               gokalshi.Sprintf(q.Limit),
		"cursor":              gokalshi.Sprintf(q.Cursor),
		"with_nested_markets": gokalshi.Sprintf(q.WithNestedMarkets),
		"status":              gokalshi.Sprintf(q.Status),
		"series_tracker":      gokalshi.Sprintf(q.SeriesTracker),
		"min_close_ts":        gokalshi.Sprintf(q.MinCloseTs),
	}
}

type EventQuery struct {
	WithNestedMarkets bool
}

func (q EventQuery) ToMap() map[string]string {
	return map[string]string{
		"with_nested_markets": gokalshi.Sprintf(q.WithNestedMarkets),
	}
}

type MarketsQuery struct {
	Limit        int
	Cursor       string
	EventTicker  string
	SeriesTicker string
	MaxCloseTs   int
	MinCloseTs   int
	Status       string
	Tickers      string
}

func (q MarketsQuery) ToMap() map[string]string {
	return map[string]string{
		"limit":         gokalshi.Sprintf(q.Limit),
		"cursor":        gokalshi.Sprintf(q.Cursor),
		"event_ticker":  gokalshi.Sprintf(q.EventTicker),
		"series_ticker": gokalshi.Sprintf(q.SeriesTicker),
		"max_close_ts":  gokalshi.Sprintf(q.MaxCloseTs),
		"min_close_ts":  gokalshi.Sprintf(q.MinCloseTs),
		"status":        gokalshi.Sprintf(q.Status),
		"tickers":       gokalshi.Sprintf(q.Tickers),
	}
}

type TradesQuery struct {
	Limit  int
	Cursor string
	Ticker string
	MinTs  int
	MaxTs  int
}

func (q TradesQuery) ToMap() map[string]string {
	return map[string]string{
		"limit":  gokalshi.Sprintf(q.Limit),
		"cursor": gokalshi.Sprintf(q.Cursor),
		"ticker": gokalshi.Sprintf(q.Ticker),
		"min_ts": gokalshi.Sprintf(q.MinTs),
		"max_ts": gokalshi.Sprintf(q.MaxTs),
	}
}

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
