package apirequests

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mwaurawakati/onlyfans/API/auth"
	el "github.com/mwaurawakati/onlyfans/API/endpointlinks"
	"github.com/pkg/errors"
)

func GetUserChats(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ChatsEndpoint()
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

type Messages struct {
	HasMore bool  `json:"hasMore"`
	List    []any `json:"list"`
}

func GetMessages(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.MessageEndpoint(id, 100)

	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

type Message struct {
	Text string `json:"text"`
	//MediaFiles   []int `json:"mediaFiles"`
	GiphyID string `json:"giphyId"`
	//Preview []any  `json:"preview"`
	//LockedText bool `json:"lockedText"`
	//Price int `json:"price"`
	//IsCouplePeopleMedia bool `json:"isCouplePeopleMedia"`
	ReplyToMessageID int `json:"replyToMessageId"`
	StoryID          int `json:"storyId"`
}

func SendMessage(payloadBuf *bytes.Buffer, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO:Currenctly, this functions is not sending media. That should be solved
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.SendMessageEndpoint(id)

	resp, err := PostResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func GetMessageById(messageid, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO:This function is not getting the message needed
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.MessageByIdEndpoint(strconv.Itoa(id), messageid)

	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

// Uses the client details to get messages
func ClientGetMessages(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.Messages
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.MessageEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.Messages
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			fmt.Println(len(p.List))
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.MessageEndpoint(id, lim, p.List[len(p.List)-1].ID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.Messages
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			url, endpoint = el.MessageEndpoint(id, lim, p.List[len(p.List)-1].ID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	fmt.Println(len(p1.List))
	return p1, nil

}

func DeleteMessage(messageid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.DeleteMessageEndpoint(messageid)
	resp, err := DeleteResponse(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func LikeMessage(messageid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.LikeMessageEndpoint(messageid)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("")
	resp, err := PostResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.LikeMessage
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.LikeMessage
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func GetMessageById2(messageid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.MessageByIdEndpoint2(messageid)
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func MarkChatAsRead(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.MarkAsReadEndpoint(id)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("")
	resp, err := PostResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.MarkAsRead //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.MarkAsRead
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func UnlikeMessage(messageid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.LikeMessageEndpoint(messageid)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("")
	resp, err := DeleteResponse(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.LikeMessage
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.LikeMessage
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func PinMessage(username, messageID, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.PinMessageEndpoint(messageID, strconv.Itoa(id))
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("")
	resp, err := PostResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func UnpinMessage(username, messageID, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.PinMessageEndpoint(messageID, strconv.Itoa(id))
	resp, err := DeleteResponse(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func HideMessage(messageID, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.HideMessageEndpoint(messageID)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode("")
	resp, err := PutResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func ReplyMessage(payloadBuf *bytes.Buffer, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO:Currenctly, this functions is not sending media. That should be solved
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.SendMessageEndpoint(id)
	resp, err := PostResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func QueryChats(limit, query, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.QueryChatsEndpoint(query, limit)
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

// Uses the client details to get messages
func GetChatMedia(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.ChatMediaResponse
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.ChatMediaEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.ChatMediaEndpoint(id, lim, p.NextLastID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			url, endpoint = el.ChatMediaEndpoint(id, lim, p.NextLastID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	return p1, nil

}

// Uses the client details to get messages
func GetUnlockedChatMedia(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.ChatMediaResponse
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.UnlockedChatMediaEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.UnlockedChatMediaEndpoint(id, lim, p.NextLastID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			url, endpoint = el.UnlockedChatMediaEndpoint(id, lim, p.NextLastID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	return p1, nil

}

// Uses the client details to get messages
func GetPhotoChatMedia(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.ChatMediaResponse
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.PhotoChatMediaEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			fmt.Println(len(p.List))
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.PhotoChatMediaEndpoint(id, lim, p.NextLastID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			url, endpoint = el.PhotoChatMediaEndpoint(id, lim, p.NextLastID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	return p1, nil

}

// Uses the client details to get messages
func GetVideoChatMedia(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.ChatMediaResponse
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.VideoChatMediaEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.VideoChatMediaEndpoint(id, lim, p.NextLastID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			url, endpoint = el.VideoChatMediaEndpoint(id, lim, p.NextLastID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	return p1, nil

}

// Uses the client details to get messages
func GetAudioChatMedia(limit, username, authID, userAgent, cookie, xBC string) (any, error) {
	//TODO: the limit is not wroking past 100 messages
	var p1 auth.ChatMediaResponse
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	HasMore := true
	lim, _ := strconv.Atoi(limit)
	url, endpoint := el.AudioChatMediaEndpoint(id, lim)
	for HasMore == true {
		resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
		if err != nil {
			return resp, err
		}
		resp1 := resp.(*http.Response)
		enc := resp1.Header.Get("Content-encoding")
		if resp1.StatusCode != http.StatusOK {
			s := strconv.Itoa(resp1.StatusCode)
			e := errors.New(s)
			body, err2 := ioutil.ReadAll(resp1.Body)
			var er auth.ErrorRespose
			err = json.Unmarshal(body, &er)
			if err2 != nil {
				return serverError, err
			}
			return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
		}
		defer resp1.Body.Close()
		if enc == "gzip" {
			gReader1, err1 := gzip.NewReader(resp1.Body)
			if err1 != nil {
				return serverError, errors.Wrap(err1, "Could not extract response body")
			}
			defer gReader1.Close()
			gBytes1, _ := ioutil.ReadAll(gReader1)
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes1, &p)
			if len(p.List) == 0 {
				break
			}
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			p1.HasMore = p.HasMore
			if len(p.List) < 100 {
				HasMore = false
			}
			lim = lim - len(p.List)
			url, endpoint = el.AudioChatMediaEndpoint(id, lim, p.NextLastID)

		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.ChatMediaResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			HasMore = p.HasMore
			lim = lim - len(p.List)
			if len(p.List) == 0 {
				break
			}
			url, endpoint = el.AudioChatMediaEndpoint(id, lim, p.NextLastID)
			p1.HasMore = p.HasMore

		}
		if lim <= 0 {
			HasMore = false
		}
	}
	return p1, nil

}

func GetChatsLists(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ChatsListsEndpoint()
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatListResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func GetUserUnreadChats(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.UnreadMessageEndpoint()
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func GetUserTippedChats(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.TippedMessagesEndpoint()
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func GetPriorityUnreadChats(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PriorityMessagesEndpoint()
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}

func GetListUnreadChats(listid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ListMessagesEndpoint(listid)
	resp, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp, err
	}
	resp1 := resp.(*http.Response)
	enc := resp1.Header.Get("Content-encoding")
	if resp1.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp1.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp1.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	}
	defer resp1.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp1.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p auth.ChatResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		return string(gBytes), nil
	}

}
