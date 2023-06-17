//Client handlers are aimed to specifically handle requests from the client.
//These requests have query parameters.

package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	ap "github.com/mwaurawakati/onlyfans/API/apirequests"
	"github.com/mwaurawakati/onlyfans/API/auth"
)

// @Summary Sends a message on behalf of the client.
// @Description Takes message string and giphyID as a string.
// @Description failure if provided details are no accepted by onlyfans
// @Description the message and giphyID are parsed are url query parameters
// @Description payload for the message to be implemented later
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/messages [post]
func ClientSendMessage(ctx *fiber.Ctx) error {
	username := ctx.Params("usermame")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	message := &ap.Message{
		Text: ctx.Query("message"),
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(message)
	u, err := ap.SendMessage(payloadBuf, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets messages on behalf of the client with a given username.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/messages [get]
func ClientGetMessages(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.ClientGetMessages(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets the client messages with a given user that has media.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/media [get]
func ClientGetChatMedia(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	log.Println(username)
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetChatMedia(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets client messages that are unlocked(already paid and free).
// @Description Gets photo media.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/media/unlocked [get]
func ClientGetUnlockedChatMedia(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	log.Println(username)
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetUnlockedChatMedia(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets messages with photos.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/media/photos [get]
func ClientGetPhotoChatMedia(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	log.Println(username)
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetPhotoChatMedia(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets client messages with videos
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/media/videos [get]
func ClientGetVideoChatMedia(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	log.Println(username)
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetVideoChatMedia(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets client messages with audios
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/media/audio [get]
func ClientGetAudioChatMedia(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	log.Println(username)
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetAudioChatMedia(limit, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Deletes a message.
// @Summary The message must have been sent not more than 120 sec before sending the request.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid [delete]
func ClientDeleteMessage(ctx *fiber.Ctx) error {
	messageid := ctx.Params("messageid")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.DeleteMessage(messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)

}

// @Summary Likes a given message.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid/like [post]
func ClientLikeMessage(ctx *fiber.Ctx) error {
	messageid := ctx.Params("messageid")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.LikeMessage(messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Marks chats as read.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/markasread [post]
func ClientMarkChatAsRead(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.MarkChatAsRead(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Unlikes a given message.
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid/unlike [delete]
func ClientUnlikeMessage(ctx *fiber.Ctx) error {
	messageid := ctx.Params("messageid")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.UnlikeMessage(messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Subscribes the client to a given user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/subscribe [post]
func ClientSubscribe(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.Subscribe(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary unsubscribes the client from a given user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/unsubscribe [delete]
func ClientUnsubscribe(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.Unsubscribe(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Unsubscribing resons
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/reasons [get]
func Reasons(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.UnsubscribeReasons(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets the client's profile information
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile [get]
func ClientGetProfile(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetProfile(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Edits the client's profile
// @Description NB:PATCH requests not currently working
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/edit [patch]
func ClientEditProfile(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	payloadBuf := new(bytes.Buffer)
	profile := &ap.Profile{
		Bio:         ctx.Query("bio"),
		Website:     ctx.Query("website"),
		Location:    ctx.Query("location"),
		Username:    ctx.Query("username"),
		DisplayName: ctx.Query("name"),
		Wishlist:    ctx.Query("wishlist"),
	}
	json.NewEncoder(payloadBuf).Encode(profile)
	u, err := ap.EditProfile(payloadBuf, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary pins a given message
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid/pin [post]
func ClientPinMessage(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Query("username")
	messageid := ctx.Params("messageid")
	u, err := ap.PinMessage(username, messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Unpins a pinned message
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid/unpin [delete]
func ClientUnpinMessage(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Query("username")
	messageid := ctx.Params("messageid")
	u, err := ap.UnpinMessage(username, messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Hides a given message
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/messages/:messageid/hide [put]
func ClientHideMessage(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	messageid := ctx.Params("messageid")
	u, err := ap.HideMessage(messageid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary replies to a given message. The message and message id should be passed in query params
// @Summary payloads to be implemented later
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/:username/messages/reply [post]
func ClientReplyMessage(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	msgid, errm := strconv.Atoi(ctx.Query("messageid"))
	if errm != nil {
		ctx.Status(406)
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Invalid Message Id" + errm.Error(),
			Code:    "406",
		})

	}
	message := &ap.Message{
		Text:             ctx.Query("message"),
		ReplyToMessageID: msgid,
		GiphyID:          ctx.Query("giphyId"),
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(message)
	u, err := ap.ReplyMessage(payloadBuf, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets client notifications
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications [get]
func ClientGetNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client bookmarks
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks [get]
func ClientGetBookmarks(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarks(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets bookmarked photos details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks/photos [get]
func ClientGetBookmarksPhotos(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarksPhotos(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client's bookmarked videos details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks/videos [get]
func ClientGetBookmarksVideos(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarksVideos(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets other bookmarks
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks/others [get]
func ClientGetBookmarksOthers(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarksOthers(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets audio bookmarks details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks/audios [get]
func ClientGetBookmarksAudio(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarksAudio(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets archived bookmarks details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/bookmarks/archived [get]
func ClientGetBookmarksArchived(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetBookmarksArchived(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Clients get another user profile details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username [get]
func ClientGetUser(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUser(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets posts of a given user on behalf of the client
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/posts [get]
func ClientGetUserPosts(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	limit := ctx.Query("limit")
	u, err := ap.GetUserPosts(username, limit, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Get the clients posts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/posts [get]
func ClientGetMyPosts(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetPosts("100", authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets a given post given the post id
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/posts/id/:id [get]
func ClientGetPostByID(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	id := ctx.Params("id")
	u, err := ap.GetPost(id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets a user's social buttons
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/social/buttons [get]
func ClientGetSocialButtons(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Params("username")
	u, err := ap.GetSocialButtons(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets a given user labels
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/labels [get]
func ClientGetUserLabels(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Params("username")
	u, err := ap.GetUserLabels(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets a given user highlights
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/highlights [get]
func ClientGetUserHighlights(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Params("username")
	u, err := ap.GetHighlights(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary get client lists
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/lists [get]
func ClientGetLists(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetLists(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client chats
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats [get]
func ClientGetChats(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserChats(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets the notofication count of the client
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/count [get]
func ClientGetNotificationsCount(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserNotificationsCount(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets the clients tabs orders
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notofications/tabsorder [get]
func ClientGetTabsOrder(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserTabsOrder(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets commented notofications
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/commented [get]
func ClientGetCommentedNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserCommentedNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets mentioned notifications for the client
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/mentioned [get]
func ClientGetMentionedNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserMentionedNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets subscription notifications
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/subscriptions [get]
func ClientGetSubscriptionsNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserSubscriptionNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets promotion notifications
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/promotion [get]
func ClientGetPromotionsNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserPromotionsNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary marks notifications as read
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/read [post]
func ClientReadNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.ReadNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary marks commented notifications only as read
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/commented/read [post]
func ClientReadCommentedNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.ReadCommentedNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary marks mentioned notifications only as read
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/mentioned/read [post]
func ClientReadMentionedNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.ReadMentionedNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Marks subscription messages as read
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/subscriptions/read [post]
func ClientReadSubscriptionsNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.ReadSubscriptionsNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary marks promotional messages as read
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/notifications/promotions/read [post]
func ClientReadPromotionsNotifications(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.ReadPromotionsNotifications(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary counts the number of subscriptions for the client
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions/count [get]
func ClientSubscriptionsCount(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserSubscriptionsCount(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets the list of client subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions [get]
func ClientSubscriptions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserSubscriptions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client's active subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions/active [get]
func ClientActiveSubscriptions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserActiveSubscriptions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client's expired subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions/expired [get]
func ClientExpiredSubscriptions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserExpiredSubscriptions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets client's subscriptions that require attention
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions/attentionrequired [get]
func ClientAttentionRequiredSubscriptions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserAttentionRequiredSubscriptions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary the clents searches through their subscriptions
// @Summary the query param should be passed as url query
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/subscriptions/query [get]
func ClientQuerySubscriptions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	query := ctx.Query("query")
	limit := ctx.Query("limit")
	u, err := ap.QuerySubscriptions(limit, query, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary client query chats. Query param should be passed as url query
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/chats/query [get]
func ClientQueryChats(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	query := ctx.Query("query")
	limit := ctx.Query("limit")
	u, err := ap.QueryChats(limit, query, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets a given user stories
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/stories [get]
func ClientGetStories(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	username := ctx.Params("username")
	u, err := ap.GetUserStories(username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns the client settings
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/settings [post]
func ClientGetSettings(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserSettings(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns clients notification settings
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/settings/notifications [get]
func ClientGetNotificationsSettings(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserNotificationsSettings(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Returns client account payout details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/payout/account [get]
func ClientGetPayoutAccountDetails(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetPayoutAccountDetails(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns clients payout 	details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/payout/legal [get]
func ClientGetPayoutLegalDetails(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetPayoutLegalDetails(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Returns a list of clients transactions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions [get]
func ClientGetTransactions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserTransactions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns a list of clients payment cards
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/cards [get]
func ClientGetCards(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserCards(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns true if client has transactions and false otherwise
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/has [get]
func ClientGetHasTransactions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserHasTransactions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary gets transactions that didnt go through due to an error
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/errored [get]
func ClientGetErroredTransactions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserErroredTransactions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns a list of client's topup transactions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/topup [get]
func ClientGetTopupTransactions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserTopupTransactions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns a list of payment transactions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/payments [get]
func ClientGetPaymentTransactions(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserPaymentTransactions(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary makes a given card default card
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/cards/:cardid/makedefault [put]
func ClientMakeCardDefault(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	cardid := ctx.Params("cardid")
	u, err := ap.MakeCardDefault(cardid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary removes card from onlyfans
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/transactions/card/:cardid/delete [delete]
func ClientDeleteCard(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	cardid := ctx.Params("cardid")
	u, err := ap.DeleteCard(cardid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary creates a list
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/lists/create/:listname [post]
func ClientCreateList(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	listName := ctx.Params("listname")
	u, err := ap.CreateList(listName, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary removes a given list
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/lists/remove/:listid [delete]
func ClientRemoveList(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	listID := ctx.Params("listid")
	u, err := ap.RemoveList(listID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary adds a given user to a list
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/lists/:listid/users/:userid [post]
func ClientAddUserToList(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	listID := ctx.Params("listid")
	userID := ctx.Params("userid")
	u, err := ap.AddUserToList(listID, userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary removes a given user from a list
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/lists/:listid/users/:userid/remove [delete]
func ClientRemoveUserFromList(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	listID := ctx.Params("listid")
	userID := ctx.Params("userid")
	u, err := ap.RemoveUserFromList(listID, userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary restricts a user from sending a message the client messages
// @Similar to blocking
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/restrict [post]
func ClientRestrictUser(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.RestrictUser(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary unrestricts a restricted user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/unrestrict [delete]
func ClientUnrestrictUser(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.UnrestrictUser(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary block the given user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/block [post]
func ClientBlockUser(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.BlockUser(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Unblocks a given user
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/unblock [delete]
func ClientUnblockUser(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.UnblockUser(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary this functions enables posts of a given user to appear in client's feed
// @Summary the posts must have intially be hidden
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/user/:username/posts/unhide [delete]
func ClientShowPosts(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.ShowPosts(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary hides users posts from client feeds
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/user/:username/posts/hide [put]
func ClientHidePosts(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	u, err := ap.HidePosts(userID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Returns a given user posts with photos
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/posts/photos [get]
func ClientGetUserPhotos(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	limit := ctx.Query("limit")
	u, err := ap.GetUserPhotos(userID, limit, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns a given user's video posts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/posts.videos [get]
func ClientGetUserVideos(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	limit := ctx.Query("limit")
	u, err := ap.GetUserVideos(userID, limit, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns a given user's media
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/users/:username/posts/medias [get]
func ClientGetUserMedias(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	userID := ctx.Params("username")
	limit := ctx.Query("limit")
	u, err := ap.GetUserMedias(userID, limit, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary returns general(ALL) stories for the client's subscriptions
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/stories [get]
func ClientGetGeneralStories(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetGeneralStories(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary marks a story as watched
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/stories/:storyid/watched [put]
func ClientWatchStory(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	storyID := ctx.Params("storyid")
	u, err := ap.WatchStory(storyID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary sends a reply to a given story
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/stories/reply [post]
func ClientReplyStory(ctx *fiber.Ctx) error {
	username := ctx.Query("username")
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	msgid, errm := strconv.Atoi(ctx.Query("storyid"))
	if errm != nil {
		return ctx.JSON(&auth.GenericResponse{
			Message: "Invalid story Id" + errm.Error(),
			Status:  false,
			Code:    "406",
		})
	}
	message := &ap.Message{
		Text:    ctx.Query("message"),
		StoryID: msgid,
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(message)
	u, err := ap.ReplyStory(payloadBuf, username, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary likes a given story
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/stories/:storyid/like [post]
func ClientLikeStory(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	storyID := ctx.Params("storyid")
	u, err := ap.LikeStory(storyID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary unlikes a liked story
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/stories/:storyid/unlike [delete]
func ClientUnlikeStory(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	storyID := ctx.Params("storyid")
	u, err := ap.UnlikeStory(storyID, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary get recommends
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/recommends [get]
func ClientGetRecommends(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetUserRecommends(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary get client feeds
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/feeds [get]
func ClientGetFeed(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetFeed(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)

}

// @Summary get hints
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/hints [get]
func ClientGetHints(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetHints(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)

}

// @Summary get client ip details
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/ip [get]
func ClientGetIP(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetIP(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)

}
