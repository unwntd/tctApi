package organizer

import "time"

type Organizer struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"type:varchar(50);not null" json:"name"`
	Address        string    `gorm:"type:text;not null" json:"address"`
	PicName        string    `gorm:"type:varchar(50);not null" json:"pic_name"`
	PicPhoneNumber string    `gorm:"type:varchar(20);not null" json:"pic_phone_number"`
	PicEmail       string    `gorm:"type:text;not null" json:"pic_email"`
	LogoURL        string    `gorm:"type:text;not null" json:"logo_url"`
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"not null" json:"updated_at"`
}
