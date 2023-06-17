package auth

type (
	List struct {
		CanAddUsers         bool   `json:"canAddUsers"`
		CanDeleteUsers      bool   `json:"canDeleteUsers"`
		CanManageUsers      bool   `json:"canManageUsers"`
		CanPinnedToChat     bool   `json:"canPinnedToChat"`
		CanPinnedToFeed     bool   `json:"canPinnedToFeed"`
		CanUpdate           bool   `json:"canUpdate"`
		CustomOrderUsersIds []any  `json:"customOrderUsersIds"`
		Direction           string `json:"direction"`
		ID                  string `json:"id"`
		IsPinnedToChat      bool   `json:"isPinnedToChat"`
		IsPinnedToFeed      bool   `json:"isPinnedToFeed"`
		Name                string `json:"name"`
		Order               string `json:"order"`
		Posts               []any  `json:"posts"`
		PostsCount          int    `json:"postsCount"`
		Type                string `json:"type"`
		Users               []any  `json:"users"`
		UsersCount          int    `json:"usersCount"`
	}
)
