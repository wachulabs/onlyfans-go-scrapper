package auth

type (
	SocialButton struct {
		Id          string `json:"id"`
		Sort        int    `json:"sort"`
		Label       string `json:"label"`
		Username    string `json:"username"`
		SocialMedia string `json:"socialMedia"`
		Link        string `json:"link"`
		Clicks      int    `json:"clicks"`
		Url         string `json:"url"`
		IsValid     bool   `json:"isValid"`
	}

	ButtonPayload struct {
		IsValid     bool   `json:"isValid"`
		Label       string `json:"label"`
		Link        string `json:"link"`
		SocialMedia string `json:"socialMedia"`
		Username    string `json:"username"`
	}
)
