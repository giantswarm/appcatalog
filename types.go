package appcatalog

type index struct {
	Entries map[string][]entry `json:"entries"`
}

type entry struct {
	AppVersion string   `json:"appVersion"`
	Name       string   `json:"name"`
	Urls       []string `json:"urls"`
	Version    string   `json:"version"`
}
