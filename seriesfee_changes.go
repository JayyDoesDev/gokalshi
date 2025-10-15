package gokalshi

type SeriesfeeChanges struct {
	SeriesFeeChangeArr []map[string]struct {
		FeeMultiplier int    `json:"fee_multiplier"`
		FeeType       string `json:"fee_type"`
		ID            string `json:"id"`
		ScheduledTs   string `json:"scheduled_ts"`
		SeriesTicker  string `json:"series_ticker"`
	} `json:"series_fee_change_arr"`
}

func GetSeriesfeeChanges(keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/series/fee_changes", "GET", keyID, keyPem, auth, map[string]string{})
}
