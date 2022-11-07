package appcatalog

import (
	"context"
	"fmt"
	"io"
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
	entry, err := GetLatestEntry(ctx, storageURL, app, appVersion)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return entry.Urls[0], nil
}

// GetLatestVersion returns the latest app version for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestVersion(ctx context.Context, storageURL, app, appVersion string) (string, error) {
	entry, err := GetLatestEntry(ctx, storageURL, app, appVersion)
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
	if u.Scheme == "oci" {
		u.Path = path.Join(u.Path, fmt.Sprintf("%s:%s", appName, version))
	} else {
		u.Path = path.Join(u.Path, fmt.Sprintf("%s-%s.tgz", appName, version))
	}
	return u.String(), nil
}

// GetLatestEntry returns the latest app entry for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestEntry(ctx context.Context, storageURL, app, appVersion string) (Entry, error) {
	index, err := getIndex(storageURL)
	if err != nil {
		return Entry{}, microerror.Mask(err)
	}

	entries, ok := index.Entries[app]
	if !ok {
		return Entry{}, microerror.Maskf(notFoundError, "no app %#q in index.yaml", app)
	}

	var latestCreated *time.Time
	var latestEntry Entry
	for _, e := range entries {
		if appVersion != "" {
			// If the version suffix doesn't match the SHA string we skip it.
			if !strings.HasSuffix(e.Version, appVersion) {
				continue
			}
		}

		if latestCreated == nil || e.Created.After(*latestCreated) {
			latestCreated = &e.Created
			latestEntry = e
			continue
		}
	}

	if latestEntry.Name != "" {
		return latestEntry, nil
	}

	return Entry{}, microerror.Maskf(notFoundError, "no app %#q in index.yaml with given appVersion %#q", app, appVersion)
}

func getIndex(storageURL string) (index, error) {
	indexURL := fmt.Sprintf("%s/index.yaml", storageURL)

	// We use https in catalog URLs so we can disable the linter in this case.
	resp, err := http.Get(indexURL) // #nosec
	if err != nil {
		return index{}, microerror.Mask(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
