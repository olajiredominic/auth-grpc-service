package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RequestParams struct {
	Method  string
	URL     string
	Payload map[string]string
	Headers map[string]string
}

func MakeHttpRequest(params RequestParams) (*http.Response, error) {
	var httpReq *http.Request
	var err error

	if params.Method == http.MethodGet {
		// Build query parameters for GET request
		queryParams := url.Values{}
		for key, value := range params.Payload {
			queryParams.Add(key, value)
		}
		params.URL += "?" + queryParams.Encode()
		httpReq, err = http.NewRequest(params.Method, params.URL, nil)
	} else {
		// Marshal the payload for non-GET requests
		jsonPayload, err := json.Marshal(params.Payload)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to marshal request payload: %v", err)
		}
		httpReq, _ = http.NewRequest(params.Method, params.URL, bytes.NewBuffer(jsonPayload))
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create HTTP request: %v", err)
	}
	// Set header for non-GET requests
	if params.Method != http.MethodGet {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	for key, value := range params.Headers {
		httpReq.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to make HTTP request: %v", err)
	}

	return resp, nil
}

func PostRequest(url string, payload map[string]string, headers map[string]string) (*http.Response, error) {
	return MakeHttpRequest(RequestParams{
		Method:  "POST",
		URL:     url,
		Payload: payload,
		Headers: headers,
	})
}
func PatchRequest(url string, payload map[string]string, headers map[string]string) (*http.Response, error) {
	return MakeHttpRequest(RequestParams{
		Method:  "PATCH",
		URL:     url,
		Payload: payload,
		Headers: headers,
	})
}
func PutRequest(url string, payload map[string]string, headers map[string]string) (*http.Response, error) {
	return MakeHttpRequest(RequestParams{
		Method:  "PUT",
		URL:     url,
		Payload: payload,
		Headers: headers,
	})
}
func GetRequest(url string, payload map[string]string, headers map[string]string) (*http.Response, error) {
	return MakeHttpRequest(RequestParams{
		Method:  "GET",
		URL:     url,
		Payload: payload,
		Headers: headers,
	})
}
