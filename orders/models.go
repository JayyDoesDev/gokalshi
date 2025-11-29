package orders

type OrderGroups struct {
	OrderGroups []OrderGroup `json:"order_groups"`
}

type OrderGroup struct {
	Id                  string `json:"id"`
	IsAutoCancelEnabled bool   `json:"is_auto_cancel_enabled"`
}
