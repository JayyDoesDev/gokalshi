package gokalshi

type ExchangeAnnoucements struct {
	Announcements []map[string]struct {
		DeliveryTime string `json:"delivery_time"`
		Message      string `json:"message"`
		Status       string `json:"status"`
		Type         string `json:"type"`
	} `json:"announcements"`
}

func GetExchangeAnnoucements(keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/exchange/annoucements", "GET", keyID, keyPem, auth, map[string]string{})
}
