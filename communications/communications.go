package communications

import "github.com/jayydoesdev/gokalshi"

type Communications struct {
	CommunicationsId string `json:"communications_id"`
}

type Quotes struct {
	Cursor string `json:"cursor"`
	Quotes []map[string]struct {
		AcceptedSide       string `json:"accepted_slide"`
		AcceptedTs         string `json:"accepted_ts"`
		CancellationReason string `json:"cancellation_reason"`
		CancelledTs        string `json:"cancelled_ts"`
	} `json:"quotes"`
}

func GetCommunicationsId(keyID, keyPem string, auth bool) ([]byte, error) {
	return gokalshi.Request[[]byte]("/communications/id", "GET", keyID, keyPem, auth)
}
