# Update go-webpbin to Go 1.25 and Latest go-binwrapper

## Overview

Update the go-webpbin project from Go 1.14 to Go 1.25 and upgrade go-binwrapper to the latest version, which has removed all binary download functionality. This is a breaking change that removes the auto-download API surface and simplifies the library to require system-installed webp tools.

## Context

- Files involved: `go.mod`, `go.sum`, `webpbin.go`, `cwebp.go`, `dwebp.go`, `cwebp_test.go`, `dwebp_test.go`, `docker/Dockerfile`, `docker/Dockerfile.alpine`, `docker/Dockerfile.arm`
- Related patterns: go-binwrapper now requires Go 1.24+, has removed `Src`, `NewSrc`, `Strip`, `SkipDownload` and all download/archive extraction code
- Dependencies removed: `github.com/mholt/archiver` and all its transitive deps (`dsnet/compress`, `golang/snappy`, `nwaples/rardecode`, `pierrec/lz4`, `ulikunitz/xz`, `xi2/xz`)

## Development Approach

- **Testing approach**: Regular (code first, then tests)
- Complete each task fully before moving to the next
- **CRITICAL: every task MUST include new/updated tests**
- **CRITICAL: all tests must pass before starting next task**

## Implementation Steps

### Task 1: Update go.mod and tidy dependencies

**Files:**
- Modify: `go.mod`
- Regenerate: `go.sum`

- [x] Update `go` directive from `1.14` to `1.25`
- [x] Update `github.com/nickalie/go-binwrapper` to latest (remove pinned commit hash)
- [x] Update `github.com/stretchr/testify` to latest
- [x] Update `golang.org/x/image` to latest
- [x] Run `go mod tidy` to remove stale indirect dependencies (archiver and friends)
- [x] Verify `go mod tidy` completes without errors

### Task 2: Remove download logic and simplify webpbin.go

**Files:**
- Modify: `webpbin.go`

- [x] Remove `skipDownload` package var
- [x] Remove `libwebpVersion` package var
- [x] Change default `dest` from `".bin/webp"` to `""` (empty string, so binaries resolve via system PATH)
- [x] Remove `SetSkipDownload` function
- [x] Remove `DetectUnsupportedPlatforms` function
- [x] Remove `loadDefaultFromENV` function
- [x] Simplify `createBinWrapper`: remove `macVersionMap`, all `Src()` calls, `.Strip(2)` call; keep `AutoExe()`, `Dest()`, and option func application
- [x] Add inline VENDOR_PATH env var support in `createBinWrapper` (check `os.Getenv("VENDOR_PATH")` and set dest if non-empty, before applying option funcs)
- [x] Replace `io/ioutil` import with `os` (ioutil deprecated since Go 1.16) - note: with download code removed, ioutil import is no longer needed at all
- [x] Remove unused imports (`bytes`, `runtime`, `io/ioutil`) that were only used by removed code
- [x] Verify `go build ./...` passes

### Task 3: Update tests

**Files:**
- Modify: `cwebp_test.go`
- Modify: `dwebp_test.go`

- [x] Remove `DetectUnsupportedPlatforms()` call from `init()` in `cwebp_test.go`
- [x] Update `TestVersionCWebP`: remove hardcoded `"1.2.0"` assertion and `DOCKER_ARM_TEST` env check; instead assert version is non-empty
- [x] Update `TestVersionDWebP`: same change - assert version is non-empty instead of hardcoded `"1.2.0"`
- [x] Run `go test ./...` to verify all tests pass (requires system-installed webp tools: `cwebp` and `dwebp`)

### Task 4: Update Dockerfiles

**Files:**
- Modify: `docker/Dockerfile`
- Modify: `docker/Dockerfile.alpine`
- Modify: `docker/Dockerfile.arm`

- [x] `docker/Dockerfile`: Update to use Go modules (remove `go get` commands, use `COPY go.mod go.sum` + `go mod download` pattern), install `webp` package via apt, remove GOPATH-based layout
- [x] `docker/Dockerfile.alpine`: Update to Go modules workflow, keep the from-source libwebp build (but consider updating to a newer libwebp version), remove `go get` and GOPATH layout
- [x] `docker/Dockerfile.arm`: Update to Go modules workflow, keep the from-source build, remove `go get` and GOPATH layout. Consider updating or removing the `resin/raspberry-pi3-alpine-golang:slim` base image (likely very outdated)

### Task 5: Verify acceptance criteria

- [x] Run full test suite: `go test ./...`
- [x] Run vet: `go vet ./...`
- [x] Verify no compilation errors: `go build ./...`
- [x] Verify go.mod has no unnecessary dependencies

### Task 6: Update documentation

- [x] Update README.md if user-facing changes
- [x] Update CLAUDE.md if internal patterns changed
- [x] Move this plan to `docs/plans/completed/`
