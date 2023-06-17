package auth

import (
	"time"
)

type (
	MarkAsRead struct {
		Success bool `json:"success"`
	}

	LikeMessage struct {
		Success bool `json:"success"`
		IsLiked bool `json:"isLiked"`
	}

	Messages struct {
		HasMore bool      `json:"hasMore"`
		List    []Message `json:"list"`
	}

	Message struct {
		ResponseType        string      `json:"responseType"`
		Text                string      `json:"text"`
		GiphyID             int         `json:"giphyId"`
		LockedText          bool        `json:"lockedText"`
		IsFree              bool        `json:"isFree"`
		Price               float64     `json:"price"`
		IsMediaReady        bool        `json:"isMediaReady"`
		MediaCount          int         `json:"mediaCount"`
		Media               []ChatMedia `json:"media"`
		Previews            []any       `json:"previews"`
		IsTip               bool        `json:"isTip"`
		IsReportedByMe      bool        `json:"isReportedByMe"`
		IsCouplePeopleMedia bool        `json:"isCouplePeopleMedia"`
		QueueID             int         `json:"queueId"`
		FromUser            FromUser    `json:"fromUser"`
		IsFromQueue         bool        `json:"isFromQueue"`
		ID                  int         `json:"id"`
		IsOpen              bool        `json:"isOpen"`
		IsNew               bool        `json:"isNew"`
		CreatedAt           time.Time   `json:"createdAt"`
		ChangedAt           time.Time   `json:"changedAt"`
		CancelSeconds       int         `json:"cancelSeconds"`
		IsLiked             bool        `json:"isLiked"`
		CanPurchase         bool        `json:"canPurchase"`
		CanPurchaseReason   string      `json:"canPurchaseReason"`
		CanReport           bool        `json:"canReport"`
		CanBePinned         bool        `json:"canBePinned"`
		IsPinned            bool        `json:"isPinned"`
		CanUnsendQueue      bool        `json:"canUnsendQueue"`
		UnsendSecondsQueue  int64       `json:"unsendSecondsQueue"`
		ReleaseForms        []any       `json:"releaseForms"`
	}

	ChatListResponse struct {
		HasMore bool   `json:"hasMore"`
		List    []List `json:"list"`
	}

	ChatResponse struct {
		HasMore    bool   `json:"hasMore"`
		List       []Chat `json:"list"`
		NextOffset int    `json:"nextOffset"`
	}

	Chat struct {
		CanGoToProfile       bool     `json:"canGoToProfile"`
		CanNotSendReason     bool     `json:"canNotSendReason"`
		CanSendMessage       bool     `json:"canSendMessage"`
		CountPinnedMessages  int      `json:"countPinnedMessages"`
		HasPurchasedFeed     bool     `json:"hasPurchasedFeed"`
		HasUnreadTips        bool     `json:"hasUnreadTips"`
		IsMutedNotifications bool     `json:"isMutedNotifications"`
		LastMessage          Message  `json:"lastMessage"`
		LastReadMessageID    int      `json:"lastReadMessageId"`
		UnreadMessagesCount  int      `json:"unreadMessagesCount"`
		WithUser             FromUser `json:"withUser"`
	}

	ChatMediaResponse struct {
		HasMore    bool      `json:"hasMore"`
		List       []Message `json:"list"`
		NextLastID int       `json:"nextLastId"`
	}

	ChatMedia struct {
		CanView       bool  `json:"canView"`
		Files         Files `json:"files"`
		Duration      int
		HasError      bool            `json:"hasError"`
		ID            int             `json:"id"`
		Info          Info            `json:"info"`
		Locked        any             `json:"locked"`
		Preview       string          `json:"preview"`
		Source        ChatMediaSource `json:"source"`
		SquarePreview string          `json:"squarePreview"`
		SRC           string          `json:"scr"`
		Thumb         string          `json:"thumb"`
		Type          string          `json:"type"`
		VideoSources  VideoSources    `json:"videoSources"`
	}

	Info struct {
		Preview Preview `json:"preview"`
		Source  Source  `json:"source"`
	}
	Preview struct {
		Height int `json:"height"`
		Size   int `json:"size"`
		Width  int `json:"width"`
	}
	Source struct {
		Height int `json:"height"`
		Size   int `json:"size"`
		Width  int `json:"width"`
	}
	ChatMediaSource struct {
		Source string `json:"source"`
	}
	VideoSources struct {
		V240 any `json:"240"`
		V720 any `json:"720"`
	}
	FromUser struct {
		ID   int    `json:"id"`
		View string `json:"_view"`
	}
	Files struct {
		Preview FilesPreview `json:"preview"`
	}

	FilesPreview struct {
		Options []Options `json:"options"`
	}
	Options struct {
		ID       int    `json:"id"`
		Selected string `json:"selected"`
		URL      string `json:"url"`
	}
)
