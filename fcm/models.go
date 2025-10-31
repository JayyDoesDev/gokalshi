package fcm

type FCMOrders struct {
	Orders []FMCOrder `json:"orders"`
	Cursor string     `json:"cursor"`
}

type FMCOrder struct {
	OrderId                 string `json:"order_id"`
	UserId                  string `json:"user_id"`
	ClientOrderId           string `json:"client_order_id"`
	Ticker                  string `json:"ticker"`
	Side                    string `json:"side"`
	Action                  string `json:"action"`
	Type                    string `json:"type"`
	Status                  string `json:"status"`
	YesPrice                int    `json:"yes_price"`
	NoPrice                 int    `json:"no_price"`
	YesPriceDollars         string `json:"yes_price_dollars"`
	NoPriceDollars          string `json:"no_price_dollars"`
	FillCount               int    `json:"fill_count"`
	RemainingCount          int    `json:"remaining_count"`
	InitialCount            int    `json:"initial_count"`
	TakerFees               int    `json:"taker_fees"`
	MakerFees               int    `json:"maker_fees"`
	TakerFillCost           int    `json:"taker_fill_cost"`
	MakerFillCost           int    `json:"maker_fill_cost"`
	TakerFillCostDollars    string `json:"taker_fill_cost_dollars"`
	MakerFillCostDollars    string `json:"maker_filler_cost_dollars"`
	QueuePosition           int    `json:"queue_position"`
	TakerFeeDollars         string `json:"taker_fee_dollars"`
	MakerFeeDollars         string `json:"maker_fees_dollars"`
	ExpirationTime          string `json:"expiration_time"`
	CreatedTime             string `json:"created_time"`
	LastUpdateTime          string `json:"last_update_time"`
	SelfTradePreventionType string `json:"self_trade_prevention_type"`
	OrderGroupId            string `json:"order_group_id"`
	CancelOrderOnPause      string `json:"cancel_order_on_pause"`
}

type FCMPositions struct {
	Cursor          string            `json:"cursor"`
	MarketPositions []MarketPositions `json:"market_positions"`
	EventPositions  []EventPositions  `json:"event_positions"`
}

type MarketPositions struct {
	Ticker                string `json:"ticker"`
	TotalTraded           int    `json:"total_traded"`
	TotalTradedDollars    string `json:"total_traded_dollars"`
	Position              int    `json:"positions"`
	MarketExposure        int    `json:"market_exposure"`
	MarketExposureDollars string `json:"market_exposure_dollars"`
	RealizedPNL           int    `json:"realized_pnl"`
	RealizedPNLDollars    int    `json:"realized_pnl_dollars"`
	RestingOrdersCount    int    `json:"resting_orders_count"`
	FeesPaid              int    `json:"fees_paid"`
	FeesPaidDollars       string `json:"fees_paid_dollars"`
	LastUpdatedTS         string `json:"last_updated_ts"`
}

type EventPositions struct {
	EventTicker          string `json:"event_ticker"`
	TotalCost            int    `json:"total_cost"`
	TotalCostDollars     string `json:"total_cost_dollars"`
	EventExposure        int    `json:"event_exposure"`
	EventExposureDollars string `json:"event_exposure_dollars"`
	RealizedPNL          int    `json:"realized_pnl"`
	RealizedPNLDollars   int    `json:"realized_pnl_dollars"`
	RestingOrdersCount   int    `json:"resting_orders_count"`
	FeesPaid             int    `json:"fees_paid"`
	FeesPaidDollars      string `json:"fees_paid_dollars"`
}
