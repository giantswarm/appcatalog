package appcatalog

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/giantswarm/microerror"
	"sigs.k8s.io/yaml"
)

// GetLatestChart returns the latest chart tarball file for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestChart(ctx context.Context, storageURL, app, appVersion string) (string, error) {
	entry, err := getLatestEntry(ctx, storageURL, app, appVersion)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return entry.Urls[0], nil
}

// GetLatestVersion returns the latest app version for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestVersion(ctx context.Context, storageURL, app, appVersion string) (string, error) {
	entry, err := getLatestEntry(ctx, storageURL, app, appVersion)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return entry.Version, nil
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

func getLatestEntry(ctx context.Context, storageURL, app, appVersion string) (entry, error) {
	index, err := getIndex(storageURL)
	if err != nil {
		return entry{}, microerror.Mask(err)
	}

	entries, ok := index.Entries[app]
	if !ok {
		return entry{}, microerror.Maskf(notFoundError, "no app %#q in index.yaml", app)
	}

	var latestCreated *time.Time
	var latestEntry entry
	for _, e := range entries {
		if appVersion == "" {
			continue
		}

		if !strings.HasSuffix(e.AppVersion, appVersion) {
			continue
		}

		if !strings.HasSuffix(e.Version, appVersion) {
			continue
		}

		t, err := parseTime(e.Created)
		if err != nil {
			return entry{}, microerror.Mask(err)
		}

		if latestCreated == nil || t.After(*latestCreated) {
			latestCreated = t
			latestEntry = e
			continue
		}
	}

	if latestEntry.Name != "" {
		return latestEntry, nil
	}

	return entry{}, microerror.Maskf(notFoundError, "no app %#q in index.yaml with given appVersion %#q", app, appVersion)
}

func getIndex(storageURL string) (index, error) {
	indexURL := fmt.Sprintf("%s/index.yaml", storageURL)

	// We use https in catalog URLs so we can disable the linter in this case.
	resp, err := http.Get(indexURL) // #nosec
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
