package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/orangekame3/ibmq/model"
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

func GetBackendStatus(token, device string) (model.BackendDetails, error) {
	resp, err := GetRequest(token, fmt.Sprintf("/backends/%s/status", device))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	var backendDetails model.BackendDetails
	err = json.Unmarshal(body, &backendDetails)
	if err != nil {
		fmt.Println("Error parsing JSON for backend details:", err)
		os.Exit(1)
	}
	return backendDetails, nil
}

func GetJobStatus(token, jobID string) (*model.JobStatus, error) {
	resp, err := GetRequest(token, fmt.Sprintf("/jobs/%s/status", jobID))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	jobStatus := &model.JobStatus{}
	err = json.NewDecoder(resp.Body).Decode(jobStatus)
	if err != nil {
		return nil, err
	}
	return jobStatus, nil
}

func GetJobResult(token, jobID string) (*model.JobResult, error) {
	resp, err := GetRequest(token, fmt.Sprintf("/jobs/%s/result", jobID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jobResult := &model.JobResult{}
	err = json.NewDecoder(resp.Body).Decode(jobResult)
	if err != nil {
		return nil, err
	}

	return jobResult, nil
}
