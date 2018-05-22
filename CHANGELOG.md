# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## 1.1.3 - 2019-05-22
### Fixed
- Fix `dep` manifest to avoid constraining versions of users' build tools.
- Drop external dependency on GDDO, which fixes an incompatibility with newer
  versions of github.com/google/protobuf.

## 1.1.2 - 2019-04-24
### Fixed
- Handle empty `Accept` header correctly.

## 1.1.1 - 2019-04-23
### Fixed
- Rely on on the standard library in unit tests.

## 1.1.0 - 2019-04-23
### Added
- Add `IsNoMatch`, which checks whether errors are caused by a lack of
  matching offers.

## 1.0.0 - 2019-04-20
### Added
- Initial stable release, exposing only the `ContentType` function for content
  type negotiation.

