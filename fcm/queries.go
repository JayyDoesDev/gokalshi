package fcm

import "github.com/jayydoesdev/gokalshi"

type FCMOrdersQuery struct {
	Ticker      string
	EventTicker string
	MinTs       int
	MaxTs       int
	Status      string
	Limit       int
	Cursor      string
	SubTraderId string
}

type FCMPositionsQuery struct {
	Ticker           string
	EventTicker      string
	CountFilter      string
	SettlementStatus string
	Limit            int
	Cursor           string
	SubTraderId      string
}

func (q FCMOrdersQuery) ToMap() map[string]string {
	subtrader_id, _ := gokalshi.SprintfWithError(q.SubTraderId, true)

	return map[string]string{
		"ticker":       gokalshi.Sprintf(q.Ticker),
		"event_ticker": gokalshi.Sprintf(q.EventTicker),
		"min_ts":       gokalshi.Sprintf(q.MinTs),
		"max_ts":       gokalshi.Sprintf(q.MaxTs),
		"status":       gokalshi.Sprintf(q.Status),
		"limit":        gokalshi.Sprintf(q.Limit),
		"cursor":       gokalshi.Sprintf(q.Cursor),
		"subtrader_id": subtrader_id,
	}
}

func (q FCMPositionsQuery) ToMap() map[string]string {
	subtrader_id, _ := gokalshi.SprintfWithError(q.SubTraderId, true)

	return map[string]string{
		"ticker":            gokalshi.Sprintf(q.Ticker),
		"event_ticker":      gokalshi.Sprintf(q.EventTicker),
		"count_filter":      gokalshi.Sprintf(q.CountFilter),
		"settlement_status": gokalshi.Sprintf(q.SettlementStatus),
		"limit":             gokalshi.Sprintf(q.Limit),
		"cursor":            gokalshi.Sprintf(q.CountFilter),
		"subtrader_id":      subtrader_id,
	}
}
