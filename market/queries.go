package market

import "github.com/jayydoesdev/gokalshi"

type EventsQuery struct {
	Limit             int
	Cursor            string
	WithNestedMarkets bool
	Status            string
	SeriesTracker     string
	MinCloseTs        int
}

func (q EventsQuery) ToMap() map[string]string {
	return map[string]string{
		"limit":               gokalshi.Sprintf(q.Limit),
		"cursor":              gokalshi.Sprintf(q.Cursor),
		"with_nested_markets": gokalshi.Sprintf(q.WithNestedMarkets),
		"status":              gokalshi.Sprintf(q.Status),
		"series_tracker":      gokalshi.Sprintf(q.SeriesTracker),
		"min_close_ts":        gokalshi.Sprintf(q.MinCloseTs),
	}
}

type EventQuery struct {
	WithNestedMarkets bool
}

func (q EventQuery) ToMap() map[string]string {
	return map[string]string{
		"with_nested_markets": gokalshi.Sprintf(q.WithNestedMarkets),
	}
}

type MarketsQuery struct {
	Limit        int
	Cursor       string
	EventTicker  string
	SeriesTicker string
	MaxCloseTs   int
	MinCloseTs   int
	Status       string
	Tickers      string
}

func (q MarketsQuery) ToMap() map[string]string {
	limit, _ := gokalshi.SprintfWithError(q.Limit, true)

	return map[string]string{
		"limit":         limit,
		"cursor":        gokalshi.Sprintf(q.Cursor),
		"event_ticker":  gokalshi.Sprintf(q.EventTicker),
		"series_ticker": gokalshi.Sprintf(q.SeriesTicker),
		"max_close_ts":  gokalshi.Sprintf(q.MaxCloseTs),
		"min_close_ts":  gokalshi.Sprintf(q.MinCloseTs),
		"status":        gokalshi.Sprintf(q.Status),
		"tickers":       gokalshi.Sprintf(q.Tickers),
	}
}

type TradesQuery struct {
	Limit  int
	Cursor string
	Ticker string
	MinTs  int
	MaxTs  int
}

func (q TradesQuery) ToMap() map[string]string {
	return map[string]string{
		"limit":  gokalshi.Sprintf(q.Limit),
		"cursor": gokalshi.Sprintf(q.Cursor),
		"ticker": gokalshi.Sprintf(q.Ticker),
		"min_ts": gokalshi.Sprintf(q.MinTs),
		"max_ts": gokalshi.Sprintf(q.MaxTs),
	}
}

type MarketOrderBookQuery struct {
	Depth int
}

func (q MarketOrderBookQuery) ToMap() map[string]string {
	return map[string]string{
		"depth": gokalshi.Sprintf(q.Depth),
	}
}

type SeriesQuery struct {
	Category               string
	IncludeProductMetadata bool
	Tags                   string
}

func (q SeriesQuery) ToMap() map[string]string {
	category, _ := gokalshi.SprintfWithError(q.Category, true)

	return map[string]string{
		"category":                 category,
		"include_product_metadata": gokalshi.Sprintf(q.IncludeProductMetadata),
		"tags":                     gokalshi.Sprintf(q.Tags),
	}
}

type MarketCandlesticksQuery struct {
	StartTs        int
	EndTs          int
	PeriodInterval int
}

func (q MarketCandlesticksQuery) ToMap() map[string]string {
	start_ts, _ := gokalshi.SprintfWithError(q.StartTs, true)
	end_ts, _ := gokalshi.SprintfWithError(q.EndTs, true)
	period_interval, _ := gokalshi.SprintfWithError(q.PeriodInterval, true)

	return map[string]string{
		"start_ts":        start_ts,
		"end_ts":          end_ts,
		"period_interval": period_interval,
	}
}
