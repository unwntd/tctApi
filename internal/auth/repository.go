package auth

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(userID uint, expiresAt time.Time) (*RefreshToken, error)
	Validate(tokenID string, userID uint) (*RefreshToken, error)
	Revoke(tokenID string, userID uint) error
}

type refreshRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshRepo{db: db}
}

func (r *refreshRepo) Create(userID uint, expiresAt time.Time) (*RefreshToken, error) {
	v := &RefreshToken{}
	r.ValidateByUserID(userID, v)

	rt := &RefreshToken{
		UserID:    userID,
		Token:     uuid.New().String(),
		ExpiredAt: expiresAt,
		Revoked:   false,
	}
	return rt, r.db.Create(rt).Error
}

func (r *refreshRepo) ValidateByUserID(userID uint, v *RefreshToken) {
	err := r.db.Where("user_id = ?", userID).First(&v).Error
	if err != nil {
		r.db.Model(&v).Where("id = ?", v.ID).Update("revoked", true)
	}
}

func (r *refreshRepo) Validate(tokenID string, userID uint) (*RefreshToken, error) {
	var token RefreshToken
	err := r.db.Where("token = ? AND user_id = ? AND revoked = ?", tokenID, userID, false).First(&token).Error
	if err != nil {
		return nil, err
	}
	if time.Now().After(token.ExpiredAt) {
		return nil, gorm.ErrRecordNotFound
	}
	return &token, nil
}

func (r *refreshRepo) Revoke(tokenID string, userID uint) error {
	return r.db.Model(&RefreshToken{}).Where("id = ? AND user_id = ?", tokenID, userID).Update("revoked", true).Error
}
