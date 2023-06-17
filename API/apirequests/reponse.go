package apirequests

import (
	"bytes"
	"github.com/mwaurawakati/onlyfans/API/auth"
	"log"
	"net/http"
)

// Response returns *http.Reponse. Created as it appears in almost if not all requests
func Response(url, endpoint, authID, userAgent, cookie, xBC string) (any, error) {
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header = header
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	return resp, nil
}

func PostResponse(r *bytes.Buffer, url, endpoint, authID, userAgent, cookie, xBC string) (any, error) {
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header = header
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	return resp, nil
}

func DeleteResponse(url, endpoint, authID, userAgent, cookie, xBC string) (any, error) {
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header.Set("Content-Type", "application/json")
	req.Header = header
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	return resp, nil
}

func PatchResponse(r *bytes.Buffer, url, endpoint, authID, userAgent, cookie, xBC string) (any, error) {
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, r)
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header.Set("Content-Type", "application/json")
	req.Header = header
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	return resp, nil
}

func PutResponse(r *bytes.Buffer, url, endpoint, authID, userAgent, cookie, xBC string) (any, error) {
	dynamicRules := GetDynamicRules()
	header := auth.CreateHeader(dynamicRules, authID, userAgent, cookie, xBC)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, r)
	if err != nil {
		log.Println("Error Creating New Request: ", err)
		return serverError, err
	}
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0
	header = auth.SignHeader(endpoint, dynamicRules, header)
	req.Header.Set("Content-Type", "application/json")
	req.Header = header
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Println("Error Sending request: ", err)
		return serverError, err1
	}
	return resp, nil
}
