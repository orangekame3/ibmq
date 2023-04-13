package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orangekame3/ibmq/model"
)

func PostRequest(token, endpointPath string, body []byte) (*http.Response, error) {
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %v", err)
	}

	endpoint.Path += endpointPath

	req, err := http.NewRequest("POST", endpoint.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating HTTP request: %v", err)
	}

	client := &http.Client{}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending HTTP request: %v", err)
	}

	return resp, nil
}

func Authenticate(token string) (string, error) {
	data := []byte(`{"api_key":"` + token + `"}`)
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %v", err)
	}
	endpoint.Path += "/users/loginWithToken"
	req, err := http.NewRequest("POST", endpoint.String(), bytes.NewBuffer(data))

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	credentials := &model.Credentials{}
	err = json.NewDecoder(resp.Body).Decode(credentials)
	if err != nil {
		return "", err
	}

	return credentials.Token, nil
}
