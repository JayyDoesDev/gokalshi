package gokalshi

type ExchangeStatus struct {
	ExchangeActive              bool   `json:"exchange_active"`
	ExchangeEstimatedResumeTime string `json:"exchange_estimated_resume_time"`
	TradingActive               bool   `json:"trading_active"`
}

func GetExchangeStatus(keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/exchange/status", "GET", keyID, keyPem, auth, map[string]string{})
}
