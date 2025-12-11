# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Migrate Chart.yaml annotations to new format as per https://docs.giantswarm.io/reference/platform-api/chart-metadata/

## [1.0.0] - 2024-03-25

### Changed

- Long deprecated `AppCatalog` CRs are now not created by default. Can be controlled by new Helm value `.vintage.appCatalog.create` which defaults to `false`.

## [0.10.1] - 2023-03-01

### Fixed

- Fix JSON schema for Helm values to actually reflect the desired data structure

## [0.10.0] - 2023-03-01

### Added

- Add `catalogNamespace` Helm value to support fixing the catalog to a given namespace instead of purely depending on `catalogVisibility`

## [0.9.1] - 2022-07-01

## [0.9.0] - 2022-07-01

- Add support for merging Secret's values into ConfigMap.

## [0.8.0] - 2022-06-13

## [0.7.0] - 2022-04-11

- Add support for OCI tarball URLs.

## [0.6.0] - 2021-07-28

### Added

- Add `Keywords` into Entry struct.

## [0.5.0] - 2021-06-28

### Added

- Add `Catalog` CRs.

## [0.4.2] - 2021-04-16

### Fixed

- Change catalog base url from giantswarm.github.com to giantswarm.github.io.

## [0.4.1] - 2021-03-09

### Fixed

- Do not compare AppVersion when getting the latest entry as it no longer
contains the Git commit SHA.

## [0.4.0] - 2021-02-02

### Added

- Added `GetLatestEntry` function and expose `Entry` struct.

## [0.3.2] - 2020-12-02

### Fixed

- Stop skipping entries when appVersion parameter is not presented

## [0.3.1] - 2020-11-23

### Changed

- Comparing `appVersion` with version, too.

## [0.3.0] - 2020-11-17

### Changed

- Comparing `appVersion` as substring.

## [0.2.7] - 2020-07-23

### Changed

- Use sigs.k8s.io/yaml not github.com/ghodss/yaml.

## [0.2.6] - 2020-07-09

### Changed

- Push chart to control plane catalog.

## [v0.2.5] 2020-06-22

### Changed

- Add appVersion parameter to GetLatestChart (#28)

## [v0.2.4] 2020-05-29

### Changed

- Add -catalog suffix to default ConfigMap and Secret names (#26)
- Use giantswarm as default namespace for catalog ConfigMap and Secret (#26)
- Support catalog-visibility annotation (#27)

- Flattening operator release structure.

[Unreleased]: https://github.com/giantswarm/appcatalog/compare/v1.0.0...HEAD

### Changed

- Migrate Chart.yaml annotations to new format as per https://docs.giantswarm.io/reference/platform-api/chart-metadata/
[1.0.0]: https://github.com/giantswarm/appcatalog/compare/v0.10.1...v1.0.0
[0.10.1]: https://github.com/giantswarm/appcatalog/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/giantswarm/appcatalog/compare/v0.9.1...v0.10.0
[0.9.1]: https://github.com/giantswarm/appcatalog/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/giantswarm/appcatalog/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/appcatalog/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/giantswarm/appcatalog/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/appcatalog/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/appcatalog/compare/v0.4.2...v0.5.0
[0.4.2]: https://github.com/giantswarm/appcatalog/compare/v0.4.1...v0.4.2
[0.4.1]: https://github.com/giantswarm/appcatalog/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/appcatalog/compare/v0.3.2...v0.4.0
[0.3.2]: https://github.com/giantswarm/appcatalog/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/appcatalog/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/appcatalog/compare/v0.2.7...v0.3.0
[0.2.7]: https://github.com/giantswarm/appcatalog/compare/v0.2.6...v0.2.7
[0.2.6]: https://github.com/giantswarm/appcatalog/compare/v0.2.5...v0.2.6
[v0.2.5]: https://github.com/giantswarm/appcatalog/compare/v0.2.4...v0.2.5
[v0.2.4]: https://github.com/giantswarm/app-operator/releases/tag/v0.2.4
