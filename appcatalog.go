package appcatalog

import (
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

// GetLatestEntry returns the latest app entry for the specified storage URL and app
// and returns notFoundError when it can't find a specified app.
func GetLatestEntry(storageURL, app, appVersion string) (Entry, error) {
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
			// appVersion could be the SHA string which is followed by the chart version.
			// if this SHA is neither the suffix of appVersion or the suffix of version in appcatalog Entry, we skip it.
			if !strings.HasSuffix(e.AppVersion, appVersion) && !strings.HasSuffix(e.Version, appVersion) {
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
