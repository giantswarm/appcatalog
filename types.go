package appcatalog

import "time"

type index struct {
	Entries map[string][]Entry `json:"entries"`
}

type Entry struct {
	AppVersion  string    `json:"appVersion"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Home        string    `json:"home"`
	Icon        string    `json:"icon"`
	Name        string    `json:"name"`
	Urls        []string  `json:"urls"`
	Version     string    `json:"version"`
}
