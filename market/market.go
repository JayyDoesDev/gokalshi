package market

import "github.com/jayydoesdev/gokalshi"

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

func GetEvents(keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/events", "GET", keyID, keyPem, true, map[string]string{})
}
