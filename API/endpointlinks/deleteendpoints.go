package endpointlinks

import "strconv"

func DeleteMessageEndpoint(messageid string) (string, string) {
	endpoint := Apiurl + "messages/" + messageid
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func UnSubscriptionEndpoint(id int) (string, string) {
	endpoint := Apiurl + "users/" + strconv.Itoa(id) + "/subscribe"
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func DeleteCardEndpoint(id string) (string, string) {
	endpoint := Apiurl + "payments/cards/" + id
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func DeleteTrialsEndpoint(id string) (string, string) {
	endpoint := Apiurl + "trials/" + id
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func DeletePromotionEndpoint(id string) (string, string) {
	endpoint := Apiurl + "promotions/" + id
	fulllink := Url + endpoint
	return fulllink, endpoint
}

func DeleteSocialButtonEndpoint(id string) (string, string) {
	endpoint := Apiurl + "users/social/buttons/" + id
	fulllink := Url + endpoint
	return fulllink, endpoint
}
