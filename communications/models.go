package communications

type Communications struct {
	CommunicationsId string `json:"communications_id"`
}

type Quotes struct {
	Cursor string `json:"cursor"`
	Quotes []map[string]struct {
		AcceptedSide            string `json:"accepted_slide"`
		AcceptedTs              string `json:"accepted_ts"`
		CancellationReason      string `json:"cancellation_reason"`
		CancelledTs             string `json:"cancelled_ts"`
		ConfirmedTs             string `json:"confirmed_ts"`
		Contracts               int    `json:"contacts"`
		CreatedTs               string `json:"created_ts"`
		CreatorId               string `json:"creator_id"`
		CreatorOrderId          string `json:"creator_order_id"`
		CreatorUserId           string `json:"creator_user_id"`
		ExecutedTs              string `json:"executed_ts"`
		Id                      string `json:"id"`
		MarketTicker            string `json:"market_ticker"`
		NoBid                   int    `json:"no_bid"`
		NoContractsOffered      int    `json:"no_contracts_offered"`
		RestRemainder           bool   `json:"rest_remainder"`
		RfqCreatorId            string `json:"rfq_creator_id"`
		RfqCreatorOrderId       string `json:"rfq_creator_order_id"`
		RfqCreatorUserId        string `json:"rfq_creator_user_id"`
		RfqId                   string `json:"rfq_id"`
		RfqTargetCostCentiCents int    `json:"rfq_target_cost_centi_cents"`
		Status                  string `json:"status"`
		UpdatedTs               string `json:"updated_ts"`
		YesBid                  int    `json:"yes_bid"`
		YesContractsOffered     int    `json:"yes_contracts_offered"`
	} `json:"quotes"`
}

type CreateQuoteOptions struct {
	NoBid         string `json:"no_bid"`
	RestRemainder bool   `json:"rest_remainder"`
	RfqId         string `json:"rfq_id"`
	YesBid        string `json:"yes_bid"`
}

type RFQ struct {
	CancellationReason  string `json:"cancellation_reason"`
	CanncelledTs        string `json:"cancelled_ts"`
	Contracts           int    `json:"contracts"`
	CreatedTs           string `json:"created_ts"`
	CreatorId           string `json:"creator_id"`
	CreatorUserId       string `json:"creator_user_id"`
	Id                  string `json:"id"`
	MarketTicker        string `json:"market_ticker"`
	MveCollectionTicker string `json:"mve_collection_ticker"`
	MveSelectedLegs     []map[string]struct {
		EventTicker  string `json:"event_ticker"`
		MarketTicker string `json:"market_ticker"`
		Side         string `json:"side"`
	} `json:"mve_selected_legs"`
	RestRemainder        bool   `json:"rest_remainder"`
	Status               string `json:"status"`
	TargetCostCentiCents int    `json:"target_cost_centi_cents"`
	UpdateTs             string `json:"updated_ts"`
}

type RFQOptions struct {
	Contracts            int    `json:"contracts"`
	MarketTicker         string `json:"market_ticker"`
	RestRemainder        bool   `json:"rest_remainder"`
	TargetCostCentiCents int    `json:"target_cost_centi_cents"`
}
