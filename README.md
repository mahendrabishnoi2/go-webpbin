# WebP Encoder/Decoder for Golang

[![](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/nickalie/go-webpbin)

WebP Encoder/Decoder for Golang based on libwebp tools.

**Requires `cwebp` and `dwebp` to be installed on your system.** This library wraps the command-line tools and does not download binaries automatically.

## Requirements

- Go 1.25 or later
- `cwebp` and `dwebp` command-line tools installed on your system

## Install

```
go get github.com/nickalie/go-webpbin
```

### Installing webp tools

**Debian/Ubuntu:**
```sh
apt-get install webp
```

**Alpine:**
```sh
apk add libwebp-tools
```

**macOS (Homebrew):**
```sh
brew install webp
```

**Windows:**

Download prebuilt binaries from the [official WebP site](https://developers.google.com/speed/webp/docs/precompiled) and add the `bin/` directory to your PATH.

## Environment variables

|Name|Default|Description|
|-----|------|-----------|
|VENDOR_PATH|_(empty)_|Directory containing `cwebp`/`dwebp` binaries. If unset, binaries are resolved via system PATH.|


## Example of usage

```go
package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"github.com/nickalie/go-webpbin"
)

func main() {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.webp")
	if err != nil {
		log.Fatal(err)
	}

	if err := webpbin.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
```

## CWebP

CWebP is a wrapper for *cwebp* command line tool.

Example to convert image.png to image.webp:

```go
err := webpbin.NewCWebP().
		Quality(80).
		InputFile("image.png").
		OutputFile("image.webp").
		Run()
```

## DWebP

DWebP is a wrapper for *dwebp* command line tool.

Example to convert image.webp to image.png:

```go
err := webpbin.NewDWebP().
		InputFile("image.webp").
		OutputFile("image.png").
		Run()
```

## Migrating from earlier versions

This version removes automatic binary downloading. The following breaking changes apply:

- `SetSkipDownload()` removed -- no longer needed since binaries are never downloaded.
- `DetectUnsupportedPlatforms()` removed -- no longer needed since all platforms require system-installed tools.
- `SKIP_DOWNLOAD` env var removed -- no longer recognized.
- `LIBWEBP_VERSION` env var removed -- the library no longer manages webp versions; your system package manager controls the installed version.
- `VENDOR_PATH` default changed from `.bin/webp` to empty (binaries resolve via system PATH).

If you previously relied on automatic downloading, install `cwebp` and `dwebp` via your system package manager before upgrading.

## Building libwebp from source

If your platform doesn't provide prebuilt webp packages, you can build libwebp from source:

```sh
apk add --no-cache --update libpng-dev libjpeg-turbo-dev giflib-dev tiff-dev autoconf automake make gcc g++ wget

wget https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.5.0.tar.gz && \
tar -xvzf libwebp-1.5.0.tar.gz && \
mv libwebp-1.5.0 libwebp && \
rm libwebp-1.5.0.tar.gz && \
cd /libwebp && \
./configure && \
make && \
make install && \
rm -rf libwebp
```
