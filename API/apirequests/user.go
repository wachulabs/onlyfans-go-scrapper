package apirequests

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/mwaurawakati/onlyfans/API/auth"
	el "github.com/mwaurawakati/onlyfans/API/endpointlinks"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	serverError = "Only fans APi Server Error"
)

// GetUser returns a auth.User struct if the user exists and auth.Error if the user does not exits
// It returns serverError incase of any other error
func GetUser(username, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ProfileEndpoint(username)
	resp1, err := Response(url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp1, err
	}
	resp := resp1.(*http.Response)
	if resp.StatusCode != http.StatusOK {
		s := strconv.Itoa(resp.StatusCode)
		e := errors.New(s)
		body, err2 := ioutil.ReadAll(resp.Body)
		var er auth.ErrorRespose
		err = json.Unmarshal(body, &er)
		if err2 != nil {
			return serverError, err
		}
		return er, errors.Wrap(e, "Status code not 200:Ok!. Got")
	} else {
		defer resp.Body.Close()
		enc := resp.Header.Get("Content-encoding")
		if enc == "gzip" {
			var u auth.User
			gReader, err := gzip.NewReader(resp.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Could not extract response body")
			}
			defer gReader.Close()
			gBytes, err := ioutil.ReadAll(gReader)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			err = json.Unmarshal(gBytes, &u)
			return u, nil
		} else {
			gBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return serverError, errors.Wrap(err, "Unable to read uncompressed content")
			}
			return string(gBytes), nil
		}
	}
}

func RestrictUser(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.RestrictUserEndpoint(strconv.Itoa(id))
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

func UnrestrictUser(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.RestrictUserEndpoint(strconv.Itoa(id))
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

func BlockUser(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.BlockUserEndpoint(strconv.Itoa(id))
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

func UnblockUser(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.BlockUserEndpoint(strconv.Itoa(id))
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

func ShowPosts(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.ShowPostsEndpoint(strconv.Itoa(id))
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
		var p any //p []auth.Post
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func HidePosts(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.ShowPostsEndpoint(strconv.Itoa(id))
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
		var p any //p []auth.Post
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p any //p []auth.Post
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}
