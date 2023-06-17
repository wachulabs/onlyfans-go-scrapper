package apirequests

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/mwaurawakati/onlyfans/API/auth"
	el "github.com/mwaurawakati/onlyfans/API/endpointlinks"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUserPhotos(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 []any
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.PhotosPostsEndpoint(id, limit)
	hasMore := true
	for hasMore {
		fmt.Println(url)
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
			var p []any
			json.Unmarshal(gBytes1, &p)
			fmt.Println(len(p))
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p)
			if limitint <= 0 {
				hasMore = false
			}

			date = p[len(p)-1].(map[string]any)["postedAtPrecise"].(string)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p []any
			json.Unmarshal(gBytes, &p)
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p1)
			if limitint <= 0 {
				hasMore = false
			}
		}

		url, endpoint = el.PhotosPostsEndpoint(id, strconv.Itoa(limitint), date)
	}
	fmt.Println(len(p1), hasMore, date)
	return p1, nil

}

func GetUserMedias(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 []any
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.MediasPostsEndpoint(id, limit)
	hasMore := true
	for hasMore {
		fmt.Println(url)
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
			var p []any
			json.Unmarshal(gBytes1, &p)
			fmt.Println(len(p))
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p)
			if limitint <= 0 {
				hasMore = false
			}

			date = p[len(p)-1].(map[string]any)["postedAtPrecise"].(string)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p []any
			json.Unmarshal(gBytes, &p)
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p1)
			if limitint <= 0 {
				hasMore = false
			}
		}

		url, endpoint = el.MediasPostsEndpoint(id, strconv.Itoa(limitint), date)
	}
	fmt.Println(len(p1), hasMore, date)
	return p1, nil

}

func GetUserVideos(username, limit, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	var p1 []any
	var date string
	if err != nil {
		return user, err
	}
	id := strconv.Itoa(user.(auth.User).Id)
	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return "Limit must be a number", err
	}
	url, endpoint := el.VideosPostsEndpoint(id, limit)
	hasMore := true
	for hasMore {
		fmt.Println(url)
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
			var p []any
			json.Unmarshal(gBytes1, &p)
			fmt.Println(len(p))
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p)
			if limitint <= 0 {
				hasMore = false
			}

			date = p[len(p)-1].(map[string]any)["postedAtPrecise"].(string)
		} else {
			gBytes, err := ioutil.ReadAll(resp1.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			var p []any
			json.Unmarshal(gBytes, &p)
			p1 = append(p1, p...)
			if len(p) < 50 {
				hasMore = false
			}
			limitint = limitint - len(p1)
			if limitint <= 0 {
				hasMore = false
			}
		}

		url, endpoint = el.VideosPostsEndpoint(id, strconv.Itoa(limitint), date)
	}
	fmt.Println(len(p1), hasMore, date)
	return p1, nil

}
