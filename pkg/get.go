package pkg

import (
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "https://us-east.quantum-computing.cloud.ibm.com"

func GetRequest(token, endpointPath string) (*http.Response, error) {
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %v", err)
	}

	endpoint.Path += endpointPath

	req, err := http.NewRequest("GET", endpoint.String(), nil)
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
