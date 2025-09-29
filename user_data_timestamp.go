package gokalshi

type UserDataTimestamp struct {
	AsOfTime string `json:"as_of_time"`
}

func GetUserDataTimestamp(keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/exchange/user_data_timestamp", "GET", keyID, keyPem, auth)
}
