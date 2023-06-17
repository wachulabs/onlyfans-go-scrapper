package auth

import (
	"time"
)

type (
	Post struct {
		ResposeType           string                `json:"responseType"`
		Id                    int                   `json:"id"`
		PostedAt              time.Time             `json:"postedAt"`
		PostedAtPrecise       any                   `json:"postedAtPrecise"`
		ExpiredAt             time.Time             `json:"expiredAt"`
		Author                User                  `json:"author"`
		Text                  string                `json:"text"`
		RawText               string                `json:"rawText"`
		LockedText            bool                  `json:"lockedText"`
		IsFavorite            bool                  `json:"isFavorite"`
		CanReport             bool                  `json:"canReport"`
		CanDelete             bool                  `json:"canDelete"`
		CanComment            bool                  `json:"canComment"`
		CanEdit               bool                  `json:"canEdit"`
		IsPinned              bool                  `json:"isPinned"`
		FavouritesCount       int                   `json:"favoritesCount"`
		MediaCount            int                   `json:"mediaCount"`
		IsMediaReady          bool                  `json:"isMediaReady"`
		Voting                []any                 `json:"voting"`
		IsOpened              bool                  `json:"isOpened"`
		CanToogleFavourite    bool                  `json:"canToggleFavorite"`
		StreamId              any                   `json:"streamId"`
		Price                 any                   `json:"price"`
		HasVoting             bool                  `json:"hasVoting"`
		IsAddedToBookmarks    bool                  `json:"isAddedToBookmarks"`
		IsArchived            bool                  `json:"isArchived"`
		IsPrivateArchived     bool                  `json:"isPrivateArchived"`
		IsDeleted             bool                  `json:"isDeleted"`
		HasUrl                bool                  `json:"hasUrl"`
		IsCouplePeopleMedia   bool                  `json:"isCouplePeopleMedia"`
		CantCommentReason     string                `json:"cantCommentReason"`
		CommentsCount         int                   `json:"commentsCount"`
		MentionedUsers        []User                `json:"mentionedUsers"`
		LinkedUsers           []User                `json:"linkedUsers"`
		LinkedPosts           []Post                `json:"linkedPosts"`
		Medias                []Media               `json:"media"`
		CanViewMedia          bool                  `json:"canViewMedia"`
		Preview               []any                 `json:"preview"`
		TipsAmount            int                   `json:"tipsAmount"`
		TipsAmountRaw         int                   `json:"tipsAmountRaw"`
		LabelStates           []CreateLabelResponse `json:"labelStates"`
		HasPurchases          bool                  `json:"hasPurchases"`
		IsPublishedWithPeriod bool                  `json:"isPublishedWithPeriod"`
		ReleaseForms          []any                 `json:"releaseForms"`
	}

	Media struct {
		Id               int                     `json:"id"`
		Type             string                  `json:"type"`
		ConvertedTovideo bool                    `json:"convertedToVideo"`
		CanView          bool                    `json:"canView"`
		HasError         bool                    `json:"hasError"`
		CreatedAt        time.Time               `json:"createdAt"`
		Info             map[string]SourceStruct `json:"info"`
		Source           SourceStruct            `json:"source"`
		SquarePreview    any                     `json:"squarePreview"`
		Full             any                     `json:"full"`
		Preview          any                     `json:"preview"`
		Thumb            any                     `json:"thumb"`
	}

	SourceStruct struct {
		Source   any `json:"source"`
		Width    int `json:"width"`
		Height   int `json:"height"`
		Size     int `json:"size"`
		Duration any `json:"duration"`
	}

	Feed struct {
		List   []any `json:"list"`
		Marker any   `json:"marker"`
	}

	PostResponse struct {
		HasMore    bool   `json:"hasMore"`
		HeadMarker any    `json:"headMarker"`
		List       []Post `json:"list"`
		TailMarker any    `json:"tailMarker"`
	}

	UserPostsResponse struct {
		HasMore    bool    `json:"hasMore"`
		HeadMarker any     `json:"headMarker"`
		List       []Post  `json:"list"`
		TailMarker any     `json:"tailMarker"`
		Counters   Counter `json:"counters"`
	}

	Counter struct {
		ArchivedPostsCount int `json:"archivedPostsCount"`
		AudiosCount        int `json:"audiosCount"`
		PhotosCount        int `json:"photosCount"`
		PostsCount         int `json:"postsCount"`
		MediasCount        int `json:"mediasCount"`
		VideosCount        int `json:"videosCount"`
	}

	CreateLabelPayload struct {
		Name string `json:"name"`
	}
	CreateLabelResponse struct {
		ID                string `json:"id"`
		IsClearInProgress bool   `json:"isClearInProgress"`
		Name              string `json:"name"`
		Posts             string `json:"posts"`
		PostsCount        int    `json:"postsCount"`
		Type              string `json:"type"`
	}

	EditPostPayload struct {
		Text                string `json:"text"`
		Preview             []any  `json:"preview"`
		MediaFiles          []int  `json:"mediaFiles"`
		LabelIDs            []int  `json:"labelIds"`
		IsCouplePeopleMedia int    `json:"isCendantPeopleMedium"`
	}

	PostStats struct {
		CommentChart        []any `json:"commentChart"`
		CommentCount        int   `json:"commentCount"`
		HasStats            bool  `json:"hasStats"`
		HasVideo            bool  `json:"hasVideo"`
		IsAvailable         bool  `json:"isAvailable"`
		LikeChart           []any `json:"likeChart"`
		LikeCount           int   `json:"likeCount"`
		LookChart           []any `json:"lookChart"`
		LookCount           int   `json:"lookCount"`
		LookDuration        int   `json:"lookDuration"`
		LookDurationAvarage int   `json:"lookDurationAverage"`
		PurchasedCount      int   `json:"purchasedCount"`
		PurchasedSumm       int   `json:"purchasedSumm"`
		PurchasesChart      []any `json:"purchasesChart"`
		TipChart            []any `json:"tipChart"`
		TipCount            int   `json:"tipsCount"`
		TipSum              int   `json:"tipsSum"`
		TipSumChart         []any `json:"tipsSumChart"`
		UniqueLookChart     []any `json:"uniqueLookChart"`
		UniqueLookCount     int   `json:"uniqueLookCount"`
	}
)
