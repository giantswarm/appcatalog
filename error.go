package appcatalog

import "github.com/giantswarm/microerror"

var invalidTarballError = &microerror.Error{
	Kind: "executionFailed",
}

// IsInvalidTarballError asserts invalidTarballError.
func IsInvalidTarballError(err error) bool {
	return microerror.Cause(err) == invalidTarballError
}

var notFoundError = &microerror.Error{
	Kind: "notFoundError",
}

// IsNotFound asserts notFoundError.
func IsNotFound(err error) bool {
	return microerror.Cause(err) == notFoundError
}
