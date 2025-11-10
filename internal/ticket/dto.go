package ticket

import (
	"encoding/json"
	"errors"
	"time"
)

type TicketRequest struct {
	TierID       uint      `json:"tier_id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        float64   `json:"price" binding:"required"`
	Tax          *float64  `json:"tax"`
	Limit        *int64    `json:"limit"`
	AvailableQty *int64    `json:"available_qty"`
	PendingQty   *int64    `json:"pending_qty"`
	StartPeriod  time.Time `json:"start_period" binding:"required" time_format:"2006-01-02" time_utc:"true"`
	EndPeriod    time.Time `json:"end_period" binding:"required" time_format:"2006-01-02"  time_utc:"true"`
	IsRefundable *bool     `json:"is_refundable"` // default false if nil
}

func (r *TicketRequest) UnmarshalJSON(b []byte) error {
	type Alias TicketRequest
	aux := &struct {
		StartPeriod string `json:"start_period"`
		EndPeriod   string `json:"end_period"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	const layout = "2006-01-02"

	if aux.StartPeriod != "" {
		t, err := time.Parse(layout, aux.StartPeriod)
		if err != nil {
			return err
		}
		r.StartPeriod = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	}

	if aux.EndPeriod != "" {
		t, err := time.Parse(layout, aux.EndPeriod)
		if err != nil {
			return err
		}
		r.EndPeriod = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	}

	if r.EndPeriod.Before(r.StartPeriod) {
		return errors.New("end_period must be on or after start_period")
	}

	return nil
}
