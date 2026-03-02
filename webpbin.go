package webpbin

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/nickalie/go-binwrapper"
)

// OptionFunc configures a BinWrapper during construction.
type OptionFunc func(binWrapper *binwrapper.BinWrapper) error

// SetVendorPath sets the directory where the binary is located.
func SetVendorPath(path string) OptionFunc {
	return func(binWrapper *binwrapper.BinWrapper) error {
		binWrapper.Dest(path)
		return nil
	}
}

func createBinWrapper(optionFuncs ...OptionFunc) *binwrapper.BinWrapper {
	b := binwrapper.NewBinWrapper().AutoExe()

	if path := os.Getenv("VENDOR_PATH"); path != "" {
		b.Dest(path)
	}

	for _, optionFunc := range optionFuncs {
		if err := optionFunc(b); err != nil {
			panic("webpbin: option function failed: " + err.Error())
		}
	}

	return b
}

func createReaderFromImage(img image.Image) (io.Reader, error) {
	enc := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	var buffer bytes.Buffer
	err := enc.Encode(&buffer, img)

	if err != nil {
		return nil, err
	}

	return &buffer, nil
}

func version(b *binwrapper.BinWrapper) (string, error) {
	b.Reset()
	err := b.Run("-version")

	if err != nil {
		return "", err
	}

	version := string(b.StdOut())
	version = strings.Replace(version, "\n", "", -1)
	version = strings.Replace(version, "\r", "", -1)
	return version, nil
}
