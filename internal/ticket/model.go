package ticket

import (
	tickettier "tctApi/internal/ticket-tier"
	"time"
)

type Ticket struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	TierID       uint      `gorm:"not null;index" json:"tier_id"` // FK to ticket_tiers table
	Name         string    `gorm:"type:text;not null" json:"name"`
	Description  string    `gorm:"type:text;not null" json:"description"`
	Price        float64   `gorm:"type:decimal(15,2);not null" json:"price"`
	Tax          float64   `gorm:"type:decimal(15,2);not null;default:0" json:"tax"`
	Limit        int64     `gorm:"not null" json:"limit"`
	AvailableQty int64     `gorm:"not null" json:"available_qty"`
	PendingQty   int64     `gorm:"not null" json:"pending_qty"`
	StartPeriod  time.Time `gorm:"not null" json:"start_period"`
	EndPeriod    time.Time `gorm:"not null" json:"end_period"`
	IsRefundable bool      `gorm:"not null; default:false" json:"is_refundable"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`

	// Relation
	Tier tickettier.TicketTier `gorm:"foreignKey:TierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tier"`
}
