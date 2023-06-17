package apirequests

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mwaurawakati/onlyfans/API/auth"
	el "github.com/mwaurawakati/onlyfans/API/endpointlinks"
	"github.com/pkg/errors"
)

func GetProfile(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.UserMeEndpoint()
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

func GetMeFromConfigFile(path string) string {
	header, _ := auth.CreateHeaderFromFile(path, "/api2/v2/users/me", GetDynamicRules())
	url := "https://onlyfans.com/api2/v2/users/me"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	req.Header = header
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	str, _ := auth.RespToJSON(resp, false)
	return str
}

func GetLists(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ListEndpoint()
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
		var p any
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

type Profile struct {
	Bio         string `json:"about"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Username    string `json:"username"`
	DisplayName string `json:"name"`
	Wishlist    string `json:"wishlist"`
}

func EditProfile(payloadBuf *bytes.Buffer, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.UserMeEndpoint()
	resp1, err := PatchResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
	if err != nil {
		return resp1, errors.Wrap(err, "Could not Patch")
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
			var u any
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
			var p any
			json.Unmarshal(gBytes, &p)
			return p, nil
		}
	}
}

func CreateList(listName, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ListEndpoint()
	name := struct {
		Name string `json:"name"`
	}{
		Name: listName,
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(name)
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
		return string(gBytes), nil
	}

}

func RemoveList(listID, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.RemoveListEndpoint(listID)
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

func AddUserToList(listID, userID, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(userID, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.AddUserToListEndpoint(listID, strconv.Itoa(id))
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

func RemoveUserFromList(listID, userID, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(userID, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.AddUserToListEndpoint(listID, strconv.Itoa(id))
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

func OnlyfansInit(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.InitEndpoint()
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
		var p any
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

func GetMyStories() {}

func ChangeOnlineVisibility(payloadBuf *bytes.Buffer, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.UserMeEndpoint()
	resp, err := PatchResponse(payloadBuf, url, endpoint, authID, userAgent, cookie, xBC)
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
		return string(gBytes), nil
	}
}
