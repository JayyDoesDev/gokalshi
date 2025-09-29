package market

type Market struct {
	AvailableOnBrokers   bool   `json:"available_on_brokers"`
	Category             string `json:"category"`
	CollateralReturnType string `json:"collateral_return_type"`
	EventTicker          string `json:"event_ticker"`
	Markets              []map[string]struct {
		Ticker      string `json:"ticker"`
		EventTicker string `json:"event_ticker"`
		MarketType  string `json:"market_type"`
		Title       string `json:"title"`
		Subtitle    string `json:"subtitle"`
	}
}
