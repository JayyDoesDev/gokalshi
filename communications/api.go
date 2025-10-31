package communications

import "github.com/jayydoesdev/gokalshi"

func GetCommunicationsId(keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/id", "GET", keyID, keyPem, true, map[string]string{})
}

func GetQuotes(keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes", "GET", keyID, keyPem, true, map[string]string{})
}

func CreateQuote(params CreateQuoteOptions, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes", "POST", keyID, keyPem, true, map[string]string{}, params)
}

func GetQuote(q, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes/"+q, "GET", keyID, keyPem, true, map[string]string{})
}

func DeleteQuote(q, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes/"+q, "DELETE", keyID, keyPem, true, map[string]string{})
}

func AcceptQuote(q string, opts struct {
	AcceptedSide string `json:"accepted_side"`
}, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes"+q+"/accept", "GET", keyID, keyPem, true, map[string]string{}, opts)
}

func ConfirmQuote(q, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/quotes/"+q+"/confirm", "PUT", keyID, keyPem, true, map[string]string{})
}

func GetRFQs(keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/rfqs", "GET", keyID, keyPem, true, map[string]string{})
}

func CreateRFQs(params RFQOptions, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/rfqs", "POST", keyID, keyPem, true, map[string]string{}, params)
}

func GetRFQ(rfq, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/rfqs/"+rfq, "GET", keyID, keyPem, true, map[string]string{})
}

func DeleteRFQ(rfq, keyID, keyPem string) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/rfqs/"+rfq, "DELETE", keyID, keyPem, true, map[string]string{})
}
