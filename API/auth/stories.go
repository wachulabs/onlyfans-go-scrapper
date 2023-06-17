package auth

import "time"

type (
	ArchivedStoriesResponse struct {
		HasMore bool              `json:"hasMore"`
		List    []ArchivedStories `json:"list"`
		Marker  int               `json:"marker"`
	}
	ArchivedStories struct {
		CreatedAt time.Time              `json:"createdAt"`
		ID        int                    `json:"id"`
		Media     []ArchivedStoriesMedia `json:"media"`
		Question  string                 `json:"question"`
	}
	ArchivedStoriesMedia struct {
		CanView          bool                 `json:"canView"`
		Files            ArchivedStoriesFiles `json:"files"`
		ConvertedToVideo bool                 `json:"convertedToVideo"`
		CreatedAt        time.Time            `json:"createdAt"`
		HasError         bool                 `json:"hasError"`
		ID               int                  `json:"id"`
		Type             string               `json:"type"`
	}
	ArchivedStoriesFiles struct {
		Preview       SquarePreviewAS `json:"preview"`
		Source        SourceAS        `json:"source"`
		SquarePreview SquarePreviewAS `json:"squarePreview"`
		Thumb         ThumbAS         `json:"thumb"`
	}

	ThumbAS struct {
		Height int    `json:"height"`
		Size   int    `json:"size"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	}

	SquarePreviewAS struct {
		Height  int         `json:"height"`
		Size    int         `json:"size"`
		Sources SourcesSPAS `json:"sources"`
		URL     string      `json:"url"`
		Width   int         `json:"width"`
	}

	SourcesSPAS struct {
		W150 string `json:"w150"`
		W480 string `json:"w480"`
	}

	SourceAS struct {
		Duration int        `json:"duration"`
		Height   int        `json:"height"`
		Size     int        `json:"size"`
		Sources  SourcesSAS `json:"sources"`
		URL      string     `json:"url"`
		Width    int        `json:"width"`
	}

	SourcesSAS struct {
		S240 string `json:"240"`
		S720 string `json:"720"`
	}
)
