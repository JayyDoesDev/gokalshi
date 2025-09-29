package gokalshi

type ExchangeSchedule struct {
	Schedule map[string]struct {
		MaintenanceWindows []map[string]struct {
			EndDateTime   string `json:"end_datetime"`
			StartDateTime string `json:"start_datetime"`
		} `json:"maintenance_windows"`
		StandardHours []map[string]struct {
			EndTime   string          `json:"end_time"`
			Friday    []OpenCloseTime `json:"friday"`
			Monday    []OpenCloseTime `json:"monday"`
			Saturday  []OpenCloseTime `json:"saturday"`
			StartTime string          `json:"start_time"`
			Sunday    []OpenCloseTime `json:"sunday"`
			Thursday  []OpenCloseTime `json:"thursday"`
			Tuesday   []OpenCloseTime `json:"tuesday"`
			Wednesday []OpenCloseTime `json:"wednesday"`
		}
	} `json:"schedule"`
}

type OpenCloseTime struct {
	CloseTime string `json:"close_time"`
	OpenTime  string `json:"open_time"`
}

func GetExchangeSchedule(keyID, keyPem string, auth bool) ([]byte, error) {
	return Request[[]byte]("/exchange/schedule", "GET", keyID, keyPem, auth)
}
