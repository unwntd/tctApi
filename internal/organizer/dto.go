package organizer

type OrganizerRequest struct {
	Name           string `json:"name" binding:"required"`
	Address        string `json:"address"`
	PicName        string `json:"pic_name" binding:"required"`
	PicPhoneNumber string `json:"pic_phone_number" binding:"required,numeric"`
	PicEmail       string `json:"pic_email" binding:"required,email"`
	LogoURL        string `json:"logo_url" binding:"url"`
}
