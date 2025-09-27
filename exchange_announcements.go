package gokalshi

type ExchangeAnnoucements struct {
	Announcements []map[string]struct {
		DeliveryTime string `json:"delivery_time"`
		Message      string `json:"message"`
		Status       string `json:"status"`
		Type         string `json:"type"`
	} `json:"announcements"`
}

func GetExchangeAnnoucements(keyID, keyPath string) (ExchangeAnnoucements, error) {
	data, err := Request[ExchangeAnnoucements]("/exchange/annoucements", "GET", keyID, keyPath)
	if err != nil {
		return ExchangeAnnoucements{}, err
	}

	return data, nil
}
