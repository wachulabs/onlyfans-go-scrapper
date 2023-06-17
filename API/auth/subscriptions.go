package auth

import "time"

type (
	TrialsResponse struct {
		ClaimsCount     int       `json:"claimsCount"`
		CreatedAt       time.Time `json:"createdAt"`
		ExpiredAt       time.Time `json:"expiredAt"`
		ID              int       `json:"id"`
		IsFinished      bool      `json:"isFinished"`
		SubscribeCounts int       `json:"subscribesCounts"`
		SubscribesDays  int       `json:"subscribesDays"`
		TrialLinkName   string    `json:"trialLinkName"`
		Url             string    `json:"url"`
	}

	TrialPayload struct {
		ExpireDays      int    `json:"expireDays"`
		SubscribeCounts int    `json:"subscribesCounts"`
		SubscribesDays  int    `json:"subscribesDays"`
		TrialLinkName   string `json:"trialLinkName"`
	}

	GetTrialsResponse struct {
		List    []TrialsResponse `json:"list"`
		HasMore bool             `json:"hasMore"`
	}

	PromotionPayload struct {
		Discount        int      `json:"discount"`
		FinishDays      int      `json:"finishDays"`
		Message         string   `json:"message"`
		SubscribeCounts int      `json:"subscriptionCounts"`
		SubscribesDays  int      `json:"subscribesDays"`
		Type            []string `json:"type"`
	}

	PromotionsResponse struct {
		ID              int       `json:"id"`
		CanClaim        bool      `json:"canClaim"`
		ClaimsCount     int       `json:"claimsCount"`
		CreatedAt       time.Time `json:"createdAt"`
		FinishedAt      time.Time `json:"finishAt"`
		HasRelatedPromo bool      `json:"hasRelatedPromo"`
		IsFinished      bool      `json:"isFinished"`
		Message         string    `json:"message"`
		Price           float64   `json:"price"`
		RawMessage      string    `json:"rawMessage"`
		SubscribeCounts int       `json:"subscriptionCounts"`
		SubscribesDays  int       `json:"subscribesDays"`
		Type            string    `json:"type"`
	}
)
