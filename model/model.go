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
