package auth

type NotificationsCount struct {
	All              int `json:"all"`
	Commented        int `json:"commented"`
	Mentioned        int `json:"mentioned"`
	Subscribed       int `json:"subscribed"`
	Message          int `json:"message"`
	Sytem            int `json:"system"`
	DeactivatedMedia int `json:"deactivated_media"`
}

type TabsOrder struct {
	All        any `json:"all"`
	Commented  any `json:"commented"`
	Mentioned  any `json:"mentioned"`
	Subscribed any `json:"subscribed"`
	Message    any `json:"message"`
}
