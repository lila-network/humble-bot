# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- tbd

## [v0.1.1] - 2023-06-07

### Added
- shortflag for using different config file
- config option `mastodon.visibility` to set post visibility
- config validation for `mastodon.visibility`
- systemd service file in `examples/`
- `--version` and `-v` cli flags to display the current version

### Removed
- cli option to generate bash completion file

## [v0.1.0] - 2023-06-07

### Added
- Fetch Bundles from humblebundle.com
- Post to Mastodon
- Postgres as Database and queue



[unreleased]: https://git.lauka.net/lauralani/humble-bot/compare/v0.1.1...HEAD
[v0.1.1]: https://git.lauka.net/lauralani/humble-bot/compare/v0.1.0...v0.1.1
[v0.1.0]: https://git.lauka.net/lauralani/humble-bot/releases/tag/v0.1.0