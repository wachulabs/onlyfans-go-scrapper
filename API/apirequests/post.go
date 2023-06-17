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

func GetUserPosts(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 auth.PostResponse
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.PostsByUsernameEndpoint(id, limit)
	hasMore := true
	for hasMore {
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
			var p auth.PostResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.PostResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		}
		url, endpoint = el.PostsByUsernameEndpoint(id, strconv.Itoa(limitint), date)
	}
	return p1, nil

}

func GetPosts(limit, authID, userAgent, cookie, xBC string) (any, error) {
	var p1 auth.PostResponse
	var date string

	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.PostsEndpoint([]string{limit})
	hasMore := true
	for hasMore {
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
			var p auth.PostResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.PostResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		}
		url, endpoint = el.PostsEndpoint([]string{strconv.Itoa(limitint), date})
	}
	return p1, nil
}

func GetPost(id, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PostByIdEndpoint(id)
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
		var p auth.Post
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

func GetUserPinnedPosts(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 auth.PostResponse
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.PinnedPostsByUsernameEndpoint(id, limit)
	hasMore := true
	for hasMore {
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
			var p auth.PostResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.PostResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		}
		url, endpoint = el.PinnedPostsByUsernameEndpoint(id, strconv.Itoa(limitint), date)
	}
	return p1, nil

}

func GetUserPostsWithMediaCount(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 auth.PostResponse
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.PostsWithMediaCounterByUsernameEndpoint(id, limit)
	hasMore := true
	for hasMore {
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
			var p auth.UserPostsResponse
			json.Unmarshal(gBytes1, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p auth.UserPostsResponse
			json.Unmarshal(gBytes, &p)
			p1.List = append(p1.List, p.List...)
			hasMore = p.HasMore
			limitint = limitint - len(p.List)
			if limitint <= 0 {
				hasMore = false
			}
			date = fmt.Sprintf("%v", p.TailMarker)
		}
		url, endpoint = el.PostsWithMediaCounterByUsernameEndpoint(id, strconv.Itoa(limitint), date)
	}
	return p1, nil

}

func CreateLabel(payloadBuf *bytes.Buffer, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.CreateHighlightEndpoint()
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
		var p auth.CreateLabelResponse
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

func EditPost(payloadBuf *bytes.Buffer, id, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PostByIdEndpoint(id)
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
		var p auth.Post
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

func GetPostStats(id, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PostStatsEndpoint(id)
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
		var p auth.PostStats
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
