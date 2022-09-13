package routes

type newShortFormValues struct {
	TargetUrl string `form:"target-url" json:"targetUrl"`
}

type createdShort struct {
	TargetURL string `json:"targetUrl"`
	ShortID   string `json:"shortId"`
}
