package appcatalog

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
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
	for i, e := range entries {
		if appVersion != "" && !matchesAppVersion(entries[i].Version, appVersion) {
			continue
		}

		if latestCreated == nil || entries[i].Created.After(*latestCreated) {
			latestCreated = &entries[i].Created
			latestEntry = e
			continue
		}
	}

	if latestEntry.Name != "" {
		return latestEntry, nil
	}

	return Entry{}, microerror.Maskf(notFoundError, "no app %#q in index.yaml with given appVersion %#q", app, appVersion)
}

// gitSHA matches a full or abbreviated lowercase git object name.
var gitSHA = regexp.MustCompile(`^[0-9a-f]{7,40}$`)

// devVersionSHA captures the abbreviated SHA that gitsemver appends to a development
// version, for example "ha781825" in
// "2.0.2-dev.my-branch.2026-08-19.13-58-39.ha781825".
var devVersionSHA = regexp.MustCompile(`\.h([0-9a-f]{7,40})$`)

// matchesAppVersion reports whether an index entry version satisfies the requested
// appVersion.
//
// architect used to publish charts as "<version>-<full 40 character SHA>", so a plain
// suffix test matched a caller passing CIRCLE_SHA1. Since architect orb 9.x the format is
// the gitsemver development version shown above, whose trailing SHA is abbreviated to
// seven characters -- a full SHA can never match that by suffix, and every lookup fails
// with notFoundError even though the chart was published correctly. Accept both formats so
// callers can keep passing the full SHA whichever orb published the chart.
//
// appVersion is also legitimately a plain chart version: apptest passes App.Version when
// App.SHA is empty. The abbreviated comparison therefore only applies when appVersion
// actually looks like a git object name, leaving the version path exactly as it was.
func matchesAppVersion(entryVersion, appVersion string) bool {
	// Preserved verbatim so both the full-SHA and the plain-version callers keep working.
	if strings.HasSuffix(entryVersion, appVersion) {
		return true
	}

	if !gitSHA.MatchString(appVersion) {
		return false
	}

	match := devVersionSHA.FindStringSubmatch(entryVersion)
	if match == nil {
		return false
	}

	// The entry carries an abbreviated SHA, so it matches when it abbreviates the SHA
	// the caller asked for.
	return strings.HasPrefix(appVersion, match[1])
}

func getIndex(storageURL string) (index, error) {
	indexURL := fmt.Sprintf("%s/index.yaml", storageURL)

	// We use https in catalog URLs so we can disable the linter in this case.
	resp, err := http.Get(indexURL) // #nosec
	if err != nil {
		return index{}, microerror.Mask(err)
	}
	defer func() { _ = resp.Body.Close() }()

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
