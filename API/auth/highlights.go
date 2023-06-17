package auth

import "time"

type (
	Highlights struct {
		Hasmore bool        `json:"hasMore"`
		List    []Highlight `json:"list"`
	}

	Highlight struct {
		Cover        any       `json:"cover"`
		CoverStoryId int       `json:"coverStoryId"`
		CreatedAt    time.Time `json:"createdAt"`
		Id           int       `json:"id"`
		StoriesCount int       `json:"storiesCount"`
		Tittle       string    `json:"stories"`
		UserId       int       `json:"userId"`
	}

	CreateHighlightPayload struct {
		Title        string `json:"title"`
		CoverStoryID int    `json:"coverStoryId"`
		StoryIDs     []int  `json:"storyIds"`
	}
	CreateHighlightResponse struct {
		CoverURL     string    `json:"coverUrl"`
		CoverStoryID int       `json:"coverStoryId"`
		CreatedAt    time.Time `json:"createdAt"`
		ID           int       `json:"id"`
		StoriesCount int       `json:"storiesCount"`
		Title        string    `json:"title"`
		UserID       int       `json:"userId"`
	}

	GetHighlightResponse struct {
		CoverURL     string    `json:"coverUrl"`
		CoverStoryID int       `json:"coverStoryId"`
		CreatedAt    time.Time `json:"createdAt"`
		ID           int       `json:"id"`
		StoriesCount int       `json:"storiesCount"`
		Title        string    `json:"title"`
		UserID       int       `json:"userId"`
		Stories      []Story   `json:"stories"`
	}

	Story struct {
		CanDelete         bool                   `json:"canDelete"`
		CommentCount      int                    `json:"commentsCount"`
		CreatedAt         time.Time              `json:"createdAt"`
		ID                int                    `json:"id"`
		IsHighlighCover   bool                   `json:"isHighlighCover"`
		IsLastInHighlight bool                   `json:"isLastInHighlight"`
		IsReady           bool                   `json:"isReady"`
		IsWatched         bool                   `json:"isWatched"`
		LikesCount        bool                   `json:"likesCount"`
		Question          bool                   `json:"question"`
		TipsAmount        any                    `json:"tipsAmount"`
		TipsAmountRaw     int                    `json:"tipsAmountRaw"`
		TipsCount         int                    `json:"tipsCount"`
		UserID            int                    `json:"userId"`
		ViewersCount      int                    `json:"viewersCount"`
		Media             []ArchivedStoriesMedia `json:"media"`
		Viewers           []StoryViewer          `json:"viewers"`
	}
	StoryViewer struct {
		Avatar string `json:"avatar"`
		AvatarThumbs
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		View     string `json:"view"`
	}
	AvatarThumbs struct {
		C50  string `json:"c50"`
		C144 string `json:"c144"`
	}
)
