package gokalshi

type UserDataTimestamp struct {
	AsOfTime string `json:"as_of_time"`
}

func GetUserDataTimestamp(keyID, keyPath string) (UserDataTimestamp, error) {
	data, err := Request[UserDataTimestamp]("/exchange/user_data_timestamp", "GET", keyID, keyPath)
	if err != nil {
		return UserDataTimestamp{}, err
	}

	return data, nil
}
