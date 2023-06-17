package endpointlinks

import "testing"

func TestNewEndpointLinks(t *testing.T) {

	expects := endpointLinks{
		Customer:        "https://onlyfans.com/api2/v2/users/me",
		Users:           "https://onlyfans.com/api2/v2/users/mwaurawakati",
		Subscriptions:   "https://onlyfans.com/api2/v2/subscriptions/subscribes?limit=100&offset=100&type=active",
		Lists:           "https://onlyfans.com/api2/v2/lists?limit=100&offset=0",
		ListsUsers:      "https://onlyfans.com/api2/v2/lists/mwaurawakati/users?limit=100&offset=100&query=",
		ListChats:       "https://onlyfans.com/api2/v2/chats?limit=100&offset=100&order=desc",
		PostById:        "https://onlyfans.com/api2/v2/posts/mwaurawakati",
		MessageById:     "https://onlyfans.com/api2/v2/chats/mwaurawakati/messages?limit=10&offset=0&firstId=mwaurawakati&order=desc&skip_users=all&skip_users_dups=1",
		SearchChat:      "https://onlyfans.com/api2/v2/chats/mwaurawakati/messages/search?query=Good Morning.",
		MessageApi:      "https://onlyfans.com/api2/v2/chats/mwaurawakati/messages?limit=100&offset=100&order=desc",
		SearchMessages:  "https://onlyfans.com/api2/v2/chats/mwaurawakati?limit=10&offset=0&filter=&order=activity&query=Good Morning.",
		MassMessagesApi: "https://onlyfans.com/api2/v2/messages/queue/stats?limit=100&offset=0&format=infinite",
		StoriesApi:      "https://onlyfans.com/api2/v2/users/mwaurawakati/stories?limit=100&offset=0&order=desc",
		ListHighlights:  "https://onlyfans.com/api2/v2/users/mwaurawakati/stories/highlights?limit=100&offset=0&order=desc",
		Highlight:       "https://onlyfans.com/api2/v2/stories/highlights/mwaurawakati",
		PostApi:         "https://onlyfans.com/api2/v2/users/mwaurawakati/posts?limit=100&offset=100&order=publish_date_desc&skip_users_dups=0",
		ArchivedPosts:   "https://onlyfans.com/api2/v2/users/mwaurawakati/posts/archived?limit=100&offset=100&order=publish_date_desc",
		ArchivedStories: "https://onlyfans.com/api2/v2/stories/archive/?limit=100&offset=0&order=publish_date_desc",
		PaidApi:         "https://onlyfans.com/api2/v2/posts/paid?100&offset=100",
		Pay:             "https://onlyfans.com/api2/v2/payments/pay",
		Subscribe:       "https://onlyfans.com/api2/v2/users/mwaurawakati/subscribe",
		Like:            "https://onlyfans.com/api2/v2/mwaurawakati/mwaurawakati/like",
		Favorite:        "https://onlyfans.com/api2/v2/mwaurawakati/mwaurawakati/favorites/mwaurawakati",
		Transactions:    "https://onlyfans.com/api2/v2/payments/all/transactions?limit=10&offset=0",
		TwoFactor:       "https://onlyfans.com/api2/v2/users/otp/check",
	}

	got := NewEndpointLinks(100, 100, "Good Morning.", "mwaurawakati", "mwaurawakati", "mwaurawakati")
	if got != expects {
		t.Errorf("TestFailed")
	}
}
