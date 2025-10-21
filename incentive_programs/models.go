package incentiveprograms

type IncentivePrograms struct {
	IncentivePrograms []IncentiveProgram `json:"incentive_programs"`
	NextCursor        string             `json:"next_cursor"`
}

type IncentiveProgram struct {
	DiscountFactorBPS int    `json:"discount_factor_bps"`
	EndDate           string `json:"end_date"`
	EventTicker       string `json:"event_ticker"`
	Id                string `json:"id"`
	IncentiveType     string `json:"incentive_type"`
	MarketTicker      string `json:"market_ticker"`
	PaidOut           bool   `json:"paid_out"`
	PeriodReward      int    `json:"period_reward"`
	SeriesTicker      string `json:"series_ticker"`
	StartDate         string `json:"start_date"`
	TargetSize        int    `json:"target_size"`
}
