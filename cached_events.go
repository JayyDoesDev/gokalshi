package gokalshi

type CachedEventsForcastHistory struct {
	ForcastHistory []ForcastHistory `json:"forcast_history"`
}

type ForcastHistory struct {
	EndPeriodTs      int                `json:"end_period_ts"`
	EventTicker      string             `json:"event_ticker"`
	PercentilePoints []PercentilePoints `json:"percentile_points"`
	PeriodInterval   int                `json:"period_interval"`
}

type PercentilePoints struct {
	FormattedForcast    string `json:"formatted_forcast"`
	NumericalForcast    string `json:"numerical_forcast"`
	Percentile          int    `json:"percentile"`
	RawNumericalForcast string `json:"raw_numerical_forcast"`
}

func GetCachedEventsForcastHistory(t, keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/cached/events/"+t+"/forcast_history", "GET", keyID, keyPem, auth, map[string]string{})
}
