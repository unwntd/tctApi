package tickettier

import "time"

type TicketTier struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID      uint      `gorm:"not null;index:idx_event_id" json:"event_id"` // foreign key to event table
	Name         string    `gorm:"type:varchar(50);not null" json:"name"`
	Description  string    `gorm:"type:text;not null" json:"description"`
	Limit        uint      `gorm:"not null" json:"limit"`
	AvailableQty uint      `gorm:"not null" json:"available_qty"`
	PendingQty   uint      `gorm:"not null" json:"pending_qty"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
}
