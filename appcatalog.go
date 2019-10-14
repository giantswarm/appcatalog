package appcatalog

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
)

// GetLatestChart returns the latest chart tarball file in the specified catalog.
func GetLatestChart(ctx context.Context, storage, app string) (string, error) {
	index, err := getIndex(storage)
	if err != nil {
		return "", microerror.Mask(err)
	}

	var downloadURL string
	{
		entry, ok := index.Entries[app]
		if !ok {
			return "", microerror.Maskf(notFoundError, fmt.Sprintf("no app %q in index.yaml", app))
		}
		downloadURL = entry[0].Urls[0]
	}

	return downloadURL, nil
}

// GetLatestVersion returns the latest version in the specified catalog.
func GetLatestVersion(ctx context.Context, catalog, app string) (string, error) {
	index, err := getIndex(catalog)
	if err != nil {
		return "", microerror.Mask(err)
	}

	var version string
	{
		entry, ok := index.Entries[app]
		if !ok {
			return "", microerror.Maskf(notFoundError, fmt.Sprintf("no app %#q in index.yaml", app))
		}
		version = entry[0].Version
	}

	return version, nil
}

func getIndex(storage string) (index, error) {
	indexURL := fmt.Sprintf("%s/index.yaml", storage)
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
