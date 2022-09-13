package core

type CreateShort struct {
	TargetURL string `form:"target-url" json:"targetUrl"`
}

type Short struct {
	TargetURL string `json:"targetUrl"`
	ShortID   string `json:"shortId"`
}
