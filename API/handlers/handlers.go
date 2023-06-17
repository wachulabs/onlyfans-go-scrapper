package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	ap "github.com/mwaurawakati/onlyfans/API/apirequests"
)

const (
	authid    = "293038355"
	cookie    = "sess=rsoas308gr2cl30ehuouss9fq6"
	xbc       = "5c81ee1042be3c2f0f06054041b53e9d9536ff33"
	useragent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
)

// GetUser gets the user details given the username
func GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user from onlyfans: ", username)
	u, err := ap.GetUser(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user's posts from onlyfans: ", username)
	u, err := ap.GetUserPosts(username, "10", authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting feautured posts")
	u, err := ap.GetPosts("100", authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("Getting posts by id: ", id)
	u, err := ap.GetPost(id, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)
}

func GetSocialButtons(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting social buttons for user: ", username)
	u, err := ap.GetSocialButtons(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)
}

func GetUserLabels(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user labelsfrom onlyfans: ", username)
	u, err := ap.GetUserLabels(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetHighlights(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user highligts from onlyfans: ", username)
	u, err := ap.GetHighlights(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetLists(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user lists from onlyfans: ")
	u, err := ap.GetLists(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetChats(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user chats from onlyfans: ")
	u, err := ap.GetUserChats(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user notifications from onlyfans: ")
	u, err := ap.GetUserNotifications(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetNotificationsCount(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user notifications count from onlyfans: ")
	u, err := ap.GetUserNotificationsCount(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetTabsOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user tabs order count from onlyfans: ")
	u, err := ap.GetUserTabsOrder(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetSubscribes(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user subscribes count from onlyfans: ")
	u, err := ap.GetUserSubscriptions(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetStories(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user stories count from onlyfans: ")
	u, err := ap.GetUserStories(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user suettings count from onlyfans: ")
	u, err := ap.GetUserSettings(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetRecommends(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user recommends count from onlyfans: ")
	u, err := ap.GetUserRecommends(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)
}

func GetGeneralStories(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user general stories count from onlyfans: ")
	u, err := ap.GetGeneralStories(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetFeed(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user feed count from onlyfans: ")
	u, err := ap.GetFeed(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetHints(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user Hints count from onlyfans: ")
	u, err := ap.GetHints(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	log.Println("Getting user photos count from onlyfans: ")
	u, err := ap.GetUserPhotos(username, "10", authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user transactions count from onlyfans: ")
	u, err := ap.GetUserTransactions(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetHasTransactions(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user has transactions count from onlyfans: ")
	u, err := ap.GetUserHasTransactions(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetCards(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user IP from onlyfans: ")
	u, err := ap.GetUserCards(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetIP(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting user IP from onlyfans: ")
	u, err := ap.GetIP(authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	log.Println(username)
	log.Println("Getting user messages from onlyfans with : ", username)
	u, err := ap.GetMessages(username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	log.Println("Sending user messages to onlyfans with : ", username)
	u, err := ap.SendMessage(&bytes.Buffer{}, username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetMessageById(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	messageid := mux.Vars(r)["messageid"]
	log.Println(username)
	log.Println("Getting user messages from onlyfans with : ", username)
	u, err := ap.GetMessageById(messageid, username, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}

func GetMessageById2(w http.ResponseWriter, r *http.Request) {
	messageid := mux.Vars(r)["messageid"]
	log.Println("Getting message from onlyfans with : ", messageid)
	u, err := ap.GetMessageById2(messageid, authid, useragent, cookie, xbc)
	if err != nil {
		log.Println(err)
	}
	ind, _ := json.MarshalIndent(u, " ", "   ")
	w.Write(ind)

}
