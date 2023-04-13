package model

type BackendList struct {
	Devices []string `json:"devices"`
}

type BackendDetails struct {
	State          bool   `json:"state"`
	Status         string `json:"status"`
	Message        string `json:"message"`
	LengthQueue    int    `json:"length_queue"`
	BackendVersion string `json:"backend_version"`
}

type JobRequest struct {
	QASM string `json:"qasm"`
}

type JobStatus struct {
	Status string `json:"status"`
	QobjID string `json:"qobj_id"`
}

type JobResult struct {
	Results []struct {
		Counts struct {
			Values map[string]int `json:"values"`
		} `json:"counts"`
	} `json:"results"`
}

type Credentials struct {
	Token string `json:"access_token"`
}
