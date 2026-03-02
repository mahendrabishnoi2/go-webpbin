package webpbin

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
)

func init() {
	createTestJPEG("source.jpg")
	createTestWebP("source.webp")
}

func createTestJPEG(path string) {
	if _, err := os.Stat(path); err == nil {
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x * 2), G: uint8(y * 2), B: 128, A: 255})
		}
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 90}); err != nil {
		panic(err)
	}
}

func createTestWebP(path string) {
	if _, err := os.Stat(path); err == nil {
		return
	}

	createTestJPEG("source.jpg")

	c := NewCWebP()
	c.InputFile("source.jpg")
	c.OutputFile(path)
	if err := c.Run(); err != nil {
		panic(err)
	}
}

func TestEncodeImage(t *testing.T) {
	c := NewCWebP()
	f, err := os.Open("source.jpg")
	require.NoError(t, err)
	defer f.Close()
	img, err := jpeg.Decode(f)
	require.NoError(t, err)
	c.InputImage(img)
	c.OutputFile("target.webp")
	err = c.Run()
	assert.Nil(t, err)
	validateWebp(t)
}

func TestEncodeReader(t *testing.T) {
	c := NewCWebP()
	f, err := os.Open("source.jpg")
	require.NoError(t, err)
	defer f.Close()
	c.Input(f)
	c.OutputFile("target.webp")
	err = c.Run()
	assert.Nil(t, err)
	validateWebp(t)
}

func TestEncodeFile(t *testing.T) {
	c := NewCWebP()
	c.InputFile("source.jpg")
	c.OutputFile("target.webp")
	err := c.Run()
	assert.Nil(t, err)
	validateWebp(t)
}

func TestEncodeWriter(t *testing.T) {
	f, err := os.Create("target.webp")
	require.NoError(t, err)
	defer f.Close()

	c := NewCWebP()
	c.InputFile("source.jpg")
	c.Output(f)
	err = c.Run()
	assert.Nil(t, err)
	f.Close()
	validateWebp(t)
}

func TestVersionCWebP(t *testing.T) {
	c := NewCWebP()
	r, err := c.Version()
	assert.Nil(t, err)
	assert.NotEmpty(t, r)
}

func validateWebp(t *testing.T) {
	defer os.Remove("target.webp")
	fSource, err := os.Open("source.jpg")
	require.NoError(t, err)
	defer fSource.Close()
	imgSource, err := jpeg.Decode(fSource)
	require.NoError(t, err)
	fTarget, err := os.Open("target.webp")
	require.NoError(t, err)
	defer fTarget.Close()
	imgTarget, err := webp.Decode(fTarget)
	require.NoError(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
