package core

type CreateShort struct {
	TargetURL string  `form:"target-url" json:"targetUrl"`
	ExpiresAt *string `form:"expires-at" json:"expiresAt" validate:"datetime"`
}
