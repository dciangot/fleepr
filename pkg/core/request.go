package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request input struct
type Request struct {
	URL         string
	RequestType string
	Headers     map[string]string
	AuthUser    string
	AuthPwd     string
	Content     []byte
	Timeout     time.Duration
}

func validateRequest(r Request) (Request, error) {

	validatedRequest := r

	// TODO: implemente timeout from config
	//if &r.Timeout == nil {
	validatedRequest.Timeout = 5 * time.Minute
	//}

	if r.URL == "" {
		return Request{}, fmt.Errorf("URL not specified")
	}

	if r.RequestType == "" {
		validatedRequest.RequestType = "GET"
	}

	return validatedRequest, nil
}

// PrepareAuthHeaders ..
func PrepareAuthHeaders(header map[string]string) string {

	var authHeaderList []string

	for k, v := range header {

		keyTemp := fmt.Sprintf("%s = %s", k, v)
		authHeaderList = append(authHeaderList, keyTemp)
	}

	authHeader := strings.Join(authHeaderList, ";")

	//fmt.Printf(authHeader)

	return authHeader
}

// MakeRequest function based on inputs
func MakeRequest(request Request) (body []byte, statusCode int, err error) {

	var req *http.Request

	r, err := validateRequest(request)
	if err != nil {
		return nil, -1, fmt.Errorf("Failed to validate request inputs %s", err)
	}

	client := &http.Client{
		Timeout: r.Timeout,
	}

	switch r.RequestType {
	case "POST":
		req, err = http.NewRequest(r.RequestType, r.URL, strings.NewReader(string(r.Content)))
		if err != nil {
			return nil, -1, fmt.Errorf("Failed to create POST http request: %s", err)
		}
	default:
		req, err = http.NewRequest(r.RequestType, r.URL, nil)
		if err != nil {
			return nil, -1, fmt.Errorf("Failed to create %s http request: %s", r.RequestType, err)
		}
	}

	if request.AuthUser != "" && request.AuthPwd != "" {
		req.SetBasicAuth(url.QueryEscape(request.AuthUser), url.QueryEscape(request.AuthPwd))
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	fmt.Println(req.Header.Get("grant_type"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, fmt.Errorf("Remote request failed: %s", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, fmt.Errorf("Failed to read the response: %s", err)
	}

	return body, resp.StatusCode, nil
}
