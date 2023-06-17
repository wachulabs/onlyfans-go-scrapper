package handlers

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	ap "github.com/mwaurawakati/onlyfans/API/apirequests"
	"github.com/mwaurawakati/onlyfans/API/auth"
)

// @Summary Gets creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/edit [patch]
func GetTippedMessages() {}

// @Summary Gets creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/visibleonline [patch]
func ClientChangeOnlineVisibility(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	visible, erro := strconv.ParseBool(ctx.Query("visibility"))
	if erro != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: erro.Error(),
			Code:    "500",
		})
	}
	visibility := &auth.Visibility{
		IsVisibleOnline: visible,
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(visibility)
	u, err := ap.ChangeOnlineVisibility(payloadBuf, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Gets creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/socialbuttons [Get]
func ClientGetMySocialButtons(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetMySocialButtons(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Adds creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/socialbuttons [PUT]
func ClientAddSocialButton(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	button := new(auth.ButtonPayload)
	if err := ctx.BodyParser(button); err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Internal Server Error",
			Code:    "500",
		})
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(button)
	u, err := ap.AddSocialButton(payloadBuf, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Adds creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/socialbuttons/:buttonid [Delete]
func ClientDeleteSocialButton(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	buttonid := ctx.Params("buttonid")
	u, err := ap.DeleteSocialButton(buttonid, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Adds creator's tipped messages
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/createhighlight [Post]
func ClientCreateHighlight(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	button := new(auth.CreateHighlightPayload)
	if err := ctx.BodyParser(button); err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Internal Server Error",
			Code:    "500",
		})
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(button)
	u, err := ap.CreateHighlight(payloadBuf, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Get Archived stories
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/stories/archived [Get]
func ClientGetArchivedStories(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	u, err := ap.GetArchivedStories(authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Get highliht by id
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/highligh/:highlighid [Get]
func ClientGetHighlightByID(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	id := ctx.Params("highlightid")
	u, err := ap.GetHighlightByID(id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary edit highlight
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/highligh/:highlighid [Put]
func ClientEditHighlight(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	button := new(auth.CreateHighlightPayload)
	if err := ctx.BodyParser(button); err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Internal Server Error",
			Code:    "500",
		})
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(button)
	id := ctx.Params("highlightid")
	u, err := ap.EditHighlight(payloadBuf, id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Delete highliht by id
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/highligh/:highlighid [Delete]
func ClientDeleteHighlightByID(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	id := ctx.Params("highlightid")
	u, err := ap.DeleteHighlightByID(id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Create labels
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/profile/label [Post]
func ClientCreateLabel(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	button := new(auth.CreateLabelPayload)
	if err := ctx.BodyParser(button); err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Internal Server Error",
			Code:    "500",
		})
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(button)
	u, err := ap.CreateLabel(payloadBuf, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Edit post
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/posts/:id [Put]
func ClientEditPost(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	button := new(auth.EditPostPayload)
	if err := ctx.BodyParser(button); err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: "Internal Server Error",
			Code:    "500",
		})
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(button)
	id := ctx.Params("id")
	u, err := ap.EditPost(payloadBuf, id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}

// @Summary Get post stats
// @Description NB:PATCH requests not currently working
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 500 {object} auth.GenericResponse
// @Router /apis/onlyfans/v1/posts/:id/stats [Get]
func ClientGetPostStats(ctx *fiber.Ctx) error {
	authID := ctx.Query("id")
	userAgent := ctx.Query("useragent")
	cookieClient := ctx.Query("sess")
	xBC := ctx.Query("xbc")
	id := ctx.Params("id")
	u, err := ap.GetPostStats(id, authID, userAgent, cookieClient, xBC)
	if err != nil {
		return ctx.JSON(&auth.GenericResponse{
			Status:  false,
			Message: err.Error(),
			Code:    "500",
		})
	}
	return ctx.JSON(&u)
}
