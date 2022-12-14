package core

import "time"

type CreateShort struct {
	TargetURL string  `form:"target-url" json:"targetUrl" binding:"required"`
	ExpiresAt *string `form:"expires-at" json:"expiresAt"`
	MaxUses   *uint   `form:"max-uses" json:"maxUses"`
}

type CreateUser struct {
	Username        string `form:"username" json:"username" binding:"required,printascii"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password-confirm" json:"passwordConfirm" binding:"required,eqcsfield=Password"`
}

type LoginUser struct {
	Username string `form:"username" json:"username" binding:"required,printascii"`
	Password string `form:"password" json:"password" binding:"required"`
}

type APIShort struct {
	ShortID    string
	TargetURL  string     `json:"targetUrl"`
	VisitCount uint       `json:"visitCount,omitempty"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
	MaxUses    *uint      `json:"maxUses,omitempty"`
	OwnerID    *uint      `json:"ownerId,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
}
