//endpoint package holds the onlyfans website api endpoints

package endpointlinks

import (
	"strconv"
	//"fmt"
)

const (
	Url    = "https://onlyfans.com"
	Apiurl = "/api2/v2/"
)

// endpointLinks contains  all the links to onlyfans endpoint
type endpointLinks struct {
	Customer        string //
	Users           string
	Subscriptions   string
	Lists           string
	ListsUsers      string
	ListChats       string
	PostById        string
	MessageById     string
	SearchChat      string
	MessageApi      string
	SearchMessages  string
	MassMessagesApi string
	StoriesApi      string
	ListHighlights  string
	Highlight       string
	PostApi         string
	ArchivedPosts   string
	ArchivedStories string
	PaidApi         string
	Pay             string
	Subscribe       string
	Like            string
	Favorite        string
	Transactions    string
	TwoFactor       string
}

func NewEndpointLinks(GlobalLimit int, GlobalOffset int, text string, Identifiers ...string) endpointLinks {
	var identifier1, identifier2, identifier3, apiurl, strGlobalLimit, strGlobalOffset string
	apiurl = Url + Apiurl
	strGlobalOffset = strconv.Itoa(GlobalOffset)
	strGlobalLimit = strconv.Itoa(GlobalLimit)
	if len(Identifiers) >= 3 {
		identifier1 = Identifiers[0]
		identifier2 = Identifiers[1]
		identifier3 = Identifiers[2]
	} else if len(Identifiers) == 2 {
		identifier3 = ""
		identifier2 = Identifiers[1]
		identifier1 = Identifiers[0]
	} else if len(Identifiers) == 1 {
		identifier1 = Identifiers[0]
		identifier2 = ""
		identifier3 = ""
	} else {
		identifier1 = ""
		identifier2 = ""
		identifier3 = ""
	}
	Customer := apiurl + "users/me"
	Users := apiurl + "users/" + identifier1
	Subscriptions := apiurl + "subscriptions/subscribes?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&type=active"
	Lists := apiurl + "lists?limit=100&offset=0"
	ListsUsers := apiurl + "lists/" + identifier1 + "/users?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&query="
	ListChats := apiurl + "chats?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&order=desc"
	PostById := apiurl + "posts/" + identifier1 + ""
	MessageById := apiurl + "chats/" + identifier1 + "/messages?limit=10&offset=0&firstId=" + identifier2 + "&order=desc&skip_users=all&skip_users_dups=1"
	SearchChat := apiurl + "chats/" + identifier1 + "/messages/search?query=" + text
	MessageApi := apiurl + "chats/" + identifier1 + "/messages?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&order=desc"
	SearchMessages := apiurl + "chats/" + identifier1 + "?limit=10&offset=0&filter=&order=activity&query=" + text
	MassMessagesApi := apiurl + "messages/queue/stats?limit=100&offset=0&format=infinite"
	StoriesApi := apiurl + "users/" + identifier1 + "/stories?limit=100&offset=0&order=desc"
	ListHighlights := apiurl + "users/" + identifier1 + "/stories/highlights?limit=100&offset=0&order=desc"
	Highlight := apiurl + "stories/highlights/" + identifier1 + ""
	PostApi := apiurl + "users/" + identifier1 + "/posts?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&order=publish_date_desc&skip_users_dups=0"
	ArchivedPosts := apiurl + "users/" + identifier1 + "/posts/archived?limit=" + strGlobalLimit + "&offset=" + strGlobalOffset + "&order=publish_date_desc"
	ArchivedStories := apiurl + "stories/archive/?limit=100&offset=0&order=publish_date_desc"
	PaidApi := apiurl + "posts/paid?" + strGlobalLimit + "&offset=" + strGlobalOffset + ""
	Pay := apiurl + "payments/pay"
	Subscribe := apiurl + "users/" + identifier1 + "/subscribe"
	Like := apiurl + "" + identifier1 + "/" + identifier2 + "/like"
	Favorite := apiurl + "" + identifier1 + "/" + identifier2 + "/favorites/" + identifier3
	Transactions := apiurl + "payments/all/transactions?limit=10&offset=0"
	TwoFactor := apiurl + "users/otp/check"

	return endpointLinks{
		Customer,
		Users,
		Subscriptions,
		Lists,
		ListsUsers,
		ListChats,
		PostById,
		MessageById,
		SearchChat,
		MessageApi,
		SearchMessages,
		MassMessagesApi,
		StoriesApi,
		ListHighlights,
		Highlight,
		PostApi,
		ArchivedPosts,
		ArchivedStories,
		PaidApi,
		Pay,
		Subscribe,
		Like,
		Favorite,
		Transactions,
		TwoFactor,
	}

}

func ProfileEndpoint(username string) (string, string) {
	endpoint := Apiurl + "users/" + username
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func RestrictUserEndpoint(username string) (string, string) {
	endpoint := Apiurl + "users/" + username + "/restrict"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func BlockUserEndpoint(username string) (string, string) {
	endpoint := Apiurl + "users/" + username + "/block"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ShowPostsEndpoint(username string) (string, string) {
	endpoint := Apiurl + "subscriptions/" + username + "/hide-posts"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PostsEndpoint(specs ...[]string) (string, string) {
	var endpoint string
	if len(specs) > 0 {
		if len(specs[0]) == 1 {
			endpoint = Apiurl + "posts?limit=" + specs[0][0] + "&skip_users=all&format=infinite"
		} else {
			endpoint = Apiurl + "posts?limit=" + specs[0][0] + "&skip_users=all&format=infinite&beforePublishTime=" + specs[0][1]
		}
	} else {
		endpoint = Apiurl + "posts?limit=100&skip_users=all&format=infinite"
	}
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PostByIdEndpoint(id string) (string, string) {
	endpoint := Apiurl + "posts/" + id
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PostStatsEndpoint(id string) (string, string) {
	endpoint := Apiurl + "posts/" + id + "/stats"
	fulllink := Url + endpoint
	return fulllink, endpoint

}
func SocialButtonEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users/" + strconv.Itoa(id) + "/social/buttons"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func MySocialButtonEndpoint() (string, string) {
	endpoint := Apiurl + "users/social/buttons"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func HighlightsEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users/" + strconv.Itoa(id) + "/stories/highlights"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ListEndpoint() (string, string) {
	endpoint := Apiurl + "lists"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func RemoveListEndpoint(listID string) (string, string) {
	endpoint := Apiurl + "lists/" + listID
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func AddUserToListEndpoint(listID, userID string) (string, string) {
	endpoint := Apiurl + "lists/" + listID + "/users/" + userID
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func ChatsEndpoint() (string, string) {
	endpoint := Apiurl + "chats?limit=1000"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func NotificationsEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications?limit=" + limit + "&offset=0&skip_users=all&format=infinite"

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func CommentedNotificationsEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications?limit=" + limit + "&offset=0&type=commented&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func MentionedNotificationsEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications?limit=" + limit + "&offset=0&type=mentioned&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func SubscribedNotificationsEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications?limit=" + limit + "&offset=0&type=subscribed&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PromotionsNotificationsEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications?limit=" + limit + "&offset=0&type=message&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func NotificationsCountEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "users/notifications/count"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func TabsOrderEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/settings/tabs-order"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func SubscriptionsEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/subscribes?limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func ActiveSubscriptionsEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/subscribes?offset=0&type=active&sort=desc&field=expire_date&limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func ExpiredSubscriptionsEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/subscribes?offset=0&type=expired&sort=desc&field=expire_date&limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func AttentionRequiredSubscriptionsEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/subscribes?offset=0&type=attention&sort=desc&field=expire_date&limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint

}
func UserStoriesEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users/" + strconv.Itoa(id) + "/stories?limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func CommentedNotificationsCount() (string, string) {
	endpoint := Apiurl + "users/notifications/count?type=commented"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func SettingsEndpoint() (string, string) {
	endpoint := Apiurl + "users/me/settings"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func NotificationsSettingsEndpoint() (string, string) {
	endpoint := Apiurl + "users/me/settings/notifications"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func RecommendsEndpoint() (string, string) {
	endpoint := Apiurl + "users/recommends"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func GeneralStoriesEndpoint() (string, string) {
	endpoint := Apiurl + "stories"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func WatchStoryEndpoint(storyID string) (string, string) {
	endpoint := Apiurl + "stories/" + storyID + "/watched"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func LikeStoryEndpoint(storyID string) (string, string) {
	endpoint := Apiurl + "stories/" + storyID + "/like"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func FeedEndpoint() (string, string) {
	endpoint := Apiurl + "streams/feed"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func UsersEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func HintsEndpoint() (string, string) {
	endpoint := Apiurl + "users/hints"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func IPEndpoint() (string, string) {
	endpoint := Apiurl + "ip"
	fulllink := Url + endpoint
	return fulllink, endpoint

}
func ErroredTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/all/transactions?limit=100&type=all"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func TopupTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/all/transactions?limit=100&type=credit"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PaymentTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/all/transactions?limit=100&type=payment"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func AllTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/all/transactions?limit=100&type=error"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func HasTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/all/has-transactions"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func AccountPayoutDetailsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/account"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func LegalPayoutDetailsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/legal"
	fulllink := Url + endpoint
	return fulllink, endpoint

}
func CardsEndpoint() (string, string) {
	endpoint := Apiurl + "payments/cards"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PostsByUsernameEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&counters=0&pinned=0&format=infinite"
	} else {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&beforePublishTime=" + date[0] + "&counters=0&pinned=0&format=infinite"
	}
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PinnedPostsByUsernameEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&counters=0&pinned=0&format=infinite"
	} else {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&beforePublishTime=" + date[0] + "&counters=0&pinned=1&format=infinite"
	}
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PostsWithMediaCounterByUsernameEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&counters=1&pinned=0&format=infinite"
	} else {
		endpoint = Apiurl + "users/" + id + "/posts?limit=" + limit + "&order=publish_date_desc&skip_users=all&beforePublishTime=" + date[0] + "&counters=1&pinned=0&format=infinite"
	}
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PhotosPostsEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts/photos?limit=" + limit
	} else {
		endpoint = Apiurl + "users/" + id + "/posts/photos?limit=" + limit + "&order=publish_date_desc&beforePublishTime=" + date[0]
	}
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func MediasPostsEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts/medias?limit=" + limit
	} else {
		endpoint = Apiurl + "users/" + id + "/posts/medias?limit=" + limit + "&order=publish_date_desc&beforePublishTime=" + date[0]
	}
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func VideosPostsEndpoint(id, limit string, date ...string) (string, string) {
	var endpoint string
	if len(date) == 0 {
		endpoint = Apiurl + "users/" + id + "/posts/videos?limit=" + limit
	} else {
		endpoint = Apiurl + "users/" + id + "/posts/videos?limit=" + limit + "&order=publish_date_desc&beforePublishTime=" + date[0]
	}
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func LabelsEndpoint(id string) (string, string) {
	endpoint := "/api2/v2/users/" + id + "/labels"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func MessageEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/messages?limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/messages?limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&order=desc&skip_users=all&offset=0"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func ChatMediaEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media?skip_users=all&limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media?limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&skip_users=all"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func VideoChatMediaEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/videos?skip_users=all&limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/videos?limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&skip_users=all"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func AudioChatMediaEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/audios?skip_users=all&limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/audios?limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&skip_users=all"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func PhotoChatMediaEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/photos?skip_users=all&limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media/photos?limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&skip_users=all"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func UnlockedChatMediaEndpoint(id, limit int, messageid ...int) (string, string) {
	var endpoint string
	if len(messageid) == 0 {
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media?skip_users=all&opened=1&limit=" + strconv.Itoa(limit)
	} else {
		msgid := messageid[0]
		endpoint = "/api2/v2/chats/" + strconv.Itoa(id) + "/media?opened=1&limit=" + strconv.Itoa(limit) + "&id=" + strconv.Itoa(msgid) + "&skip_users=all"
	}

	fulllink := Url + endpoint
	return fulllink, endpoint

}

func MessageByIdEndpoint(identifier1, identifier2 string) (string, string) {
	endpoint := Apiurl + "chats/" + identifier1 + "/messages?limit=1&offset=-1&id=" + identifier2 + "&order=desc&skip_users=all&skip_users_dups=1"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func MessageByIdEndpoint2(identifier1 string) (string, string) {
	endpoint := Apiurl + "messages/" + identifier1
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func ReasonsEndpoint() (string, string) {
	endpoint := Apiurl + "unsubscribe/reasons"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func UserMeEndpoint() (string, string) {
	endpoint := Apiurl + "users/me"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func BookmarksEndpoint(limit string) (string, string) {
	//fmt.Println(limit)
	endpoint := Apiurl + "posts/bookmarks?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PhotoBookmarksEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "posts/bookmarks/photos?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func VideoBookmarksEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "posts/bookmarks/videos?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func AudioBookmarksEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "posts/bookmarks/audios?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func OtherBookmarksEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "posts/bookmarks/other?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ArchivedBookmarksEndpoint(limit string) (string, string) {
	endpoint := Apiurl + "posts/bookmarks/locked?limit=" + limit + "&offset=0&skip_users=all&format=infinite"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func SubscriptionsCountEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/count/all"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func QuerySubscriptionsEndpoint(query, limit string) (string, string) {
	endpoint := Apiurl + "subscriptions/subscribes?offset=0&type=all&query=" + query + "&sort=desc&field=expire_date&limit=" + limit
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func QueryChatsEndpoint(query, limit string) (string, string) {
	endpoint := Apiurl + "chats?limit=" + limit + "&offset=0&order=recent&query=" + query + "&skip_users=all"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func InitEndpoint() (string, string) {
	endpoint := Apiurl + "init"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func UnreadMessageEndpoint() (string, string) {
	endpoint := Apiurl + "chats?limit=100&offset=0&filter=unread&order=recent&skip_users=all"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PriorityMessagesEndpoint() (string, string) {
	endpoint := Apiurl + "chats?limit=100&offset=0&filter=priority&order=recent&skip_users=all"
	fulllink := Url + endpoint
	return fulllink, endpoint
}
func PinnedMessagesEndpoint() (string, string) {
	endpoint := Apiurl + "chats?limit=100&offset=0&filter=pinned&order=recent&skip_users=all"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ListMessagesEndpoint(listid string) (string, string) {
	endpoint := Apiurl + "chats?limit=10&offset=0&order=recent&skip_users=all&list_id=" + listid
	fulllink := Url + endpoint
	return fulllink, endpoint
}
func TippedMessagesEndpoint() (string, string) {
	endpoint := Apiurl + "chats?limit=10&offset=0&filter=who_tipped&order=recent&skip_users=all"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMyStoriesEndpoint() (string, string) {
	endpoint := Apiurl + "users/me/stories"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMediaFromVaultEndpoint() (string, string) {
	endpoint := Apiurl + "vault/media?limit=24&offset=24&field=recent&sort=desc"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMyPhotosFromVaultEndpoint() (string, string) {
	endpoint := Apiurl + "vault/media?limit=24&offset=0&field=recent&sort=desc&type=photo"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMyVideosFromVaultEndpoint() (string, string) {
	endpoint := Apiurl + "vault/media?limit=24&offset=0&field=recent&sort=desc&type=video"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMyAudioFromVaultEndpoint() (string, string) {
	endpoint := Apiurl + "vault/media?limit=24&offset=0&field=recent&sort=desc&type=audio"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func GetMyGIFsFromVaultEndpoint() (string, string) {
	endpoint := Apiurl + "vault/media?limit=24&offset=0&field=recent&sort=desc&type=gif"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutRequestsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/requests?limit=100&offset=0"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutBalancesEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/balances"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutCheckReceiveEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/check-receive"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutTransactionsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/transactions?limit=100"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutStatsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/stats"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutsReferralBalanceEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/referrals/balance"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PayoutsReferralRequestsEndpoint() (string, string) {
	endpoint := Apiurl + "payouts/requests/referral?limit=100&offset=0"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func MyReferralsCountndpoint() (string, string) {
	endpoint := Apiurl + "users/referrals/count"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func TrialsEndpoint() (string, string) {
	endpoint := Apiurl + "trials"
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}

func SubscribersCountEndpoint() (string, string) {
	endpoint := Apiurl + "subscriptions/subscribers/count"
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}

func ChatsListsEndpoint() (string, string) {
	endpoint := Apiurl + "lists?filter=chat&limit=100&offset=0&skip_users=all&format=infinite"
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}

func ArchivedStoriesEndpoint(params ...string) (string, string) {
	endpoint := Apiurl + "stories/archive?limit=120&offset=0&format=infinite"
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}

func CreateHighlightEndpoint(params ...string) (string, string) {
	endpoint := Apiurl + "stories/highlights"
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}

func GetHighlightByIDEndpoint(id string) (string, string) {
	endpoint := Apiurl + "stories/highlights/" + id
	fulllinl := Url + endpoint
	return fulllinl, endpoint
}
