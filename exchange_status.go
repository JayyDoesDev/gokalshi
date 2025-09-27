package gokalshi

type ExchangeStatus struct {
	ExchangeActive              bool   `json:"exchange_active"`
	ExchangeEstimatedResumeTime string `json:"exchange_estimated_resume_time"`
	TradingActive               bool   `json:"trading_active"`
}

func GetExchangeStatus(keyID, keyPath string) (ExchangeStatus, error) {
	data, err := Request[ExchangeStatus]("/exchange/status", "GET", keyID, keyPath)
	if err != nil {
		return ExchangeStatus{}, err
	}

	return data, nil
}
