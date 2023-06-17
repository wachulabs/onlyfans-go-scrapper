package apirequests

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/mwaurawakati/onlyfans/API/auth"
	el "github.com/mwaurawakati/onlyfans/API/endpointlinks"
	"github.com/pkg/errors"
)

func GetUserSubscriptions(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.SubscriptionsEndpoint()
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
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p []any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func Subscribe(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.SubscriptionEndpoint(id)
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

type Reason struct {
	UReason int `json:"reason"`
}

func Unsubscribe(username, authID, userAgent, cookie, xBC string) (any, error) {
	user, err := GetUser(username, authID, userAgent, cookie, xBC)
	if err != nil {
		return user, err
	}
	id := user.(auth.User).Id
	url, endpoint := el.UnSubscriptionEndpoint(id)
	reason := &Reason{5}
	payloadBuf, _ := json.Marshal(reason)
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(payloadBuf))
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header.Set("Content-Type", "application/json")
	req.Header = header
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	enc := resp.Header.Get("Content-encoding")
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
	}
	defer resp.Body.Close()
	if enc == "gzip" {
		gReader1, err1 := gzip.NewReader(resp.Body)
		if err1 != nil {
			return serverError, errors.Wrap(err1, "Could not extract response body")
		}
		defer gReader1.Close()
		gBytes1, _ := ioutil.ReadAll(gReader1)
		var p any
		json.Unmarshal(gBytes1, &p)
		return p, nil
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

func UnsubscribeReasons(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ReasonsEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func GetUserSubscriptionsCount(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.SubscriptionsCountEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func GetUserActiveSubscriptions(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ActiveSubscriptionsEndpoint()
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
			return serverError + "error reading gzip", errors.Wrap(err1, "Could not extract response body")
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

func GetUserExpiredSubscriptions(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.ExpiredSubscriptionsEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func GetUserAttentionRequiredSubscriptions(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.AttentionRequiredSubscriptionsEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func QuerySubscriptions(limit, query, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.QuerySubscriptionsEndpoint(query, limit)
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}

}

func GetTrials(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.TrialsEndpoint()
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
		var p auth.GetTrialsResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.GetTrialsResponse
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}

func CreateTrials(payloadBuf *bytes.Buffer, query, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.TrialsEndpoint()
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
		var p auth.TrialsResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.TrialsResponse
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}

func DeleteTrial(trialid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.DeleteTrialsEndpoint(trialid)
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

func CreatePromotion(payloadBuf *bytes.Buffer, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PromotionsEndpoint()
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
		var p auth.PromotionsResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.PromotionsResponse
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}

func FinishPromotions(promotionid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.FinishPromotionEndpoint(promotionid)
	resp, err := PostResponse(&bytes.Buffer{}, url, endpoint, authID, userAgent, cookie, xBC)
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

func ClearPromotions(promotionid, authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.DeletePromotionEndpoint(promotionid)
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
		var p auth.PromotionsResponse
		json.Unmarshal(gBytes1, &p)
		return p, nil
	} else {
		gBytes, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			return serverError, errors.Wrap(err, "Unable to read uncompressed content")
		}
		var p auth.PromotionsResponse
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}

func GetSubscribersCount(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.SubscribersCountEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}

func GetPromotions(authID, userAgent, cookie, xBC string) (any, error) {
	url, endpoint := el.PromotionsEndpoint()
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
		var p any
		json.Unmarshal(gBytes, &p)
		return p, nil
	}
}
