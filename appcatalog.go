package appcatalog

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
)

// GetLatestChart returns the latest chart tarball file for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestChart(ctx context.Context, storageURL, app, appVersion string) (string, error) {
	index, err := getIndex(storageURL)
	if err != nil {
		return "", microerror.Mask(err)
	}

	entries, ok := index.Entries[app]
	if !ok {
		return "", microerror.Maskf(notFoundError, "no app %#q in index.yaml", app)
	}

	var latestCreated *time.Time
	var latestChart string
	for _, entry := range entries {
		if appVersion != "" && entry.AppVersion != appVersion {
			continue
		}

		t, err := parseTime(entry.Created)
		if err != nil {
			return "", microerror.Mask(err)
		}

		if latestCreated == nil || t.After(*latestCreated) {
			latestCreated = t
			latestChart = entry.Urls[0]
			continue
		}
	}

	if latestChart != "" {
		return latestChart, nil
	}

	return "", microerror.Maskf(notFoundError, "no app %#q in index.yaml with given appVersion %#q", app, appVersion)
}

// GetLatestVersion returns the latest app version for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestVersion(ctx context.Context, storageURL, app string) (string, error) {
	index, err := getIndex(storageURL)
	if err != nil {
		return "", microerror.Mask(err)
	}

	var version string
	{
		entry, ok := index.Entries[app]
		if !ok {
			return "", microerror.Maskf(notFoundError, "no app %#q in index.yaml", app)
		}
		version = entry[0].Version
	}

	return version, nil
}

// NewTarballURL returns the chart tarball URL for the specified app and version.
func NewTarballURL(baseURL string, appName string, version string) (string, error) {
	if baseURL == "" || appName == "" || version == "" {
		return "", microerror.Maskf(executionFailedError, "baseURL %#q, appName %#q, release %#q should not be empty", baseURL, appName, version)
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", microerror.Mask(err)
	}
	u.Path = path.Join(u.Path, fmt.Sprintf("%s-%s.tgz", appName, version))
	return u.String(), nil
}

func getIndex(storageURL string) (index, error) {
	indexURL := fmt.Sprintf("%s/index.yaml", storageURL)
	resp, err := http.Get(indexURL)
	if err != nil {
		return index{}, microerror.Mask(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return index{}, microerror.Mask(err)
	}

	var i index
	err = yaml.Unmarshal(body, &i)
	if err != nil {
		return i, microerror.Mask(err)
	}

	return i, nil
}

func parseTime(created string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, created)
	if err != nil {
		return nil, microerror.Maskf(executionFailedError, "wrong timestamp format %#q", created)
	}
	return &t, nil
}
