# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.2] - 2021-05-12

### Fixed

- Fixed the problem of errors that satisfy the net.Error interface but aren't actually from the net package getting picked up.

## [0.0.1] - 2020-11-23

### Added

- Add GetNetError which recursively checks an error to see if it fulfills the net.Error interface.