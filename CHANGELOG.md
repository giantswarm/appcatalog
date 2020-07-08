# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/appcatalog/compare/v0.2.5...HEAD
[v0.2.5]: https://github.com/giantswarm/appcatalog/compare/v0.2.4...v0.2.5
[v0.2.4]: https://github.com/giantswarm/app-operator/releases/tag/v0.2.4
