package gokalshi

type CandleSticks struct {
	AdjustedEndTs      int                  `json:"adjusted_end_ts"`
	MarketCandleSticks []MarketCandleSticks `json:"market_candlesticks"`
}

type MarketCandleSticks struct {
	EndPeriodTs  int `json:"end_period_ts"`
	OpenInterest int `json:"open_interest"`
	Price        map[string]struct {
		Close           int    `json:"close"`
		CloseDollars    string `json:"close_dollars"`
		High            int    `json:"high"`
		HighDollars     string `json:"high_dollars"`
		Low             int    `json:"low"`
		LowDollars      string `json:"low_dollars"`
		Mean            int    `json:"mean"`
		MeanDollars     string `json:"mean_dollars"`
		Open            int    `json:"open"`
		OpenDollars     string `json:"open_dollars"`
		Previous        int    `json:"previous"`
		PreviousDollars string `json:"previous_dollars"`
	} `json:"price"`
	Volume int `json:"volume"`
	YesAsk map[string]struct {
		Close        int    `json:"close"`
		CloseDollars string `json:"close_dollars"`
		High         int    `json:"high"`
		HighDollars  string `json:"high_dollars"`
		Low          int    `json:"low"`
		LowDollars   string `json:"low_dollars"`
		Mean         int    `json:"mean"`
		MeanDollars  string `json:"mean_dollars"`
		Open         int    `json:"open"`
		OpenDollars  string `json:"open_dollars"`
	} `json:"yes_ask"`
	YesBid map[string]struct {
		Close        int    `json:"close"`
		CloseDollars string `json:"close_dollars"`
		High         int    `json:"high"`
		HighDollars  string `json:"high_dollars"`
		Low          int    `json:"low"`
		LowDollars   string `json:"low_dollars"`
		Mean         int    `json:"mean"`
		MeanDollars  string `json:"mean_dollars"`
		Open         int    `json:"open"`
		OpenDollars  string `json:"open_dollars"`
	} `json:"yes_bid"`
}

func GetEventsCandleSticks(t, keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/events/"+t+"/candlesticks", "GET", keyID, keyPem, auth)
}
