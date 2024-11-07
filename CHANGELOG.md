# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.3.0] - 2023-11-07
### Added
- Added `list` alias to `queue show`
- Added `delete` alias to `queue remove`
- config option to toggle colored output
- Docker build for armv7, arm64 and amd64

### Changed
- Changed table output style
- various default config values (see examples/config.yaml)

### Fixed
- missing error handling

## [v0.2.1] - 2023-06-11
### Changed
- when starting the app in daemon mode, the app will now wait one complete interval until it posts the next post 
instead of immediately posting one bundle

## [v0.2.0] - 2023-06-11
### Added
- Added cli option `queue show` to list currently queued bundles
[#2](https://code.lila.network/adoralaura/humble-bot/issues/2)
- Added cli option `queue fill` to query humblebundle.com for new bundles and then fill the queue 
[#2](https://code.lila.network/adoralaura/humble-bot/issues/2)
- Added cli option `queue remove` to remove a queued up bundle [#2](https://code.lila.network/adoralaura/humble-bot/issues/2)

### Fixed
- [[BUG] Error parsing mastodon response field time](https://code.lila.network/adoralaura/humble-bot/issues/1)

### Removed
- BREAKING: Removed action when executing `./humble-bot` without arguments, moved to `queue fill`

## [v0.1.2] - 2023-06-07
### Added
- logging for new mastodon post now also logs the post url

### Fixed
- example config file description of `mastodon.visibility`

### Changed
- Logging will not show timestamps when app is started via systemd

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



[unreleased]: https://code.lila.network/adoralaura/humble-bot/compare/v0.3.0...HEAD
[v0.3.0]: https://code.lila.network/adoralaura/humble-bot/compare/v0.2.1...v0.3.0
[v0.2.1]: https://code.lila.network/adoralaura/humble-bot/compare/v0.2.0...v0.2.1
[v0.2.0]: https://code.lila.network/adoralaura/humble-bot/compare/v0.1.2...v0.2.0
[v0.1.2]: https://code.lila.network/adoralaura/humble-bot/compare/v0.1.1...v0.1.2
[v0.1.1]: https://code.lila.network/adoralaura/humble-bot/compare/v0.1.0...v0.1.1
[v0.1.0]: https://code.lila.network/adoralaura/humble-bot/releases/tag/v0.1.0
