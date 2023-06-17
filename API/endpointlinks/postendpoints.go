package endpointlinks

import (
	"strconv"
)

func EmailResendEndpoint(id int) (string, string) {
	endpoint := Apiurl + "emails/resend"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func EmailChangeEndpoint(id int) (string, string) {
	endpoint := Apiurl + "emails/change"
	fulllink := Url + endpoint
	return fulllink, endpoint

}

func SendMessageEndpoint(id int) (string, string) {
	endpoint := Apiurl + "chats/" + strconv.Itoa(id) + "/messages"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func LikeMessageEndpoint(id string) (string, string) {
	endpoint := Apiurl + "messages/" + id + "/like"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func MarkAsReadEndpoint(id int) (string, string) {
	endpoint := Apiurl + "chats/" + strconv.Itoa(id) + "/mark-as-read"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func SubscriptionEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users/" + strconv.Itoa(id) + "/subscribe"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func LogEndpoint() (string, string) {
	endpoint := Apiurl + "log"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PinMessageEndpoint(messageID, userID string) (string, string) {
	endpoint := Apiurl + "messages/" + messageID + "/pin/user/" + userID
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func HideMessageEndpoint(messageID string) (string, string) {
	endpoint := Apiurl + "messages/" + messageID + "/hide"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ReadNotificationsEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/read"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ReadCommentedNotificationsEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/read/commented"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ReadMentionedNotificationsEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/read/mentioned"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ReadSubscriptionsNotificationsEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/read/subscribed"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func ReadPromotionsNotificationsEndpoint() (string, string) {
	endpoint := Apiurl + "users/notifications/read/message"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func MakeCardDefaultEndpoint(id string) (string, string) {
	endpoint := Apiurl + "payments/cards/" + id + "/default"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func FinishPromotionEndpoint(id string) (string, string) {
	endpoint := Apiurl + "promotions/" + id + "/finish"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func PromotionsEndpoint() (string, string) {
	endpoint := Apiurl + "promotions"
	fulllink := Url + endpoint
	return fulllink, endpoint
}
