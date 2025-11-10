package tickettier

type TicketTierRequest struct {
	EventID      uint   ` json:"event_id" binding:"required"`
	Name         string ` json:"name" binding:"required"`
	Description  string `json:"description"`
	Limit        *uint  `json:"limit" binding:"numeric"`
	AvailableQty *uint  `json:"available_qty" binding:"numeric"`
	PendingQty   *uint  `json:"pending_qty" binding:"numeric"`
}
