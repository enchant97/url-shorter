package core

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
