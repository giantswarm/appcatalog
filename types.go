package appcatalog

import "time"

type index struct {
	APIVersion string             `yaml:"apiVersion"`
	Entries    map[string][]entry `yaml:"entries"`
	Generated  time.Time          `yaml:"generated"`
}

type entry struct {
	APIVersion  string    `yaml:"apiVersion"`
	AppVersion  string    `yaml:"appVersion"`
	Created     time.Time `yaml:"created"`
	Description string    `yaml:"description"`
	Digest      string    `yaml:"digest"`
	Name        string    `yaml:"name"`
	Urls        []string  `yaml:"urls"`
	Version     string    `yaml:"version"`
}
