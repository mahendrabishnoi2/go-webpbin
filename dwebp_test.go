package webpbin

import (
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
)

func TestVersionDWebP(t *testing.T) {
	c := NewDWebP()
	r, err := c.Version()
	assert.Nil(t, err)
	assert.NotEmpty(t, r)
}

func TestDecodeReader(t *testing.T) {
	c := NewDWebP()
	f, err := os.Open("source.webp")
	require.NoError(t, err)
	defer f.Close()
	c.Input(f)
	c.OutputFile("target.png")
	img, err := c.Run()
	assert.Nil(t, err)
	assert.Nil(t, img)
	validatePng(t)
}

func TestDecodeFile(t *testing.T) {
	c := NewDWebP()
	c.InputFile("source.webp")
	c.OutputFile("target.png")
	img, err := c.Run()
	assert.Nil(t, err)
	assert.Nil(t, img)
	validatePng(t)
}

func TestDecodeImage(t *testing.T) {
	c := NewDWebP()
	f, err := os.Open("source.webp")
	require.NoError(t, err)
	defer f.Close()
	imgSource, err := webp.Decode(f)
	require.NoError(t, err)
	f.Seek(0, 0)
	c.Input(f)
	imgTarget, err := c.Run()
	assert.Nil(t, err)
	assert.NotNil(t, imgTarget)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}

func TestDecodeWriter(t *testing.T) {
	f, err := os.Create("target.png")
	require.NoError(t, err)
	defer f.Close()
	c := NewDWebP()
	c.InputFile("source.webp")
	c.Output(f)
	img, err := c.Run()
	assert.Nil(t, err)
	assert.Nil(t, img)
	f.Close()
	validatePng(t)
}

func validatePng(t *testing.T) {
	defer os.Remove("target.png")
	fSource, err := os.Open("source.webp")
	require.NoError(t, err)
	defer fSource.Close()
	imgSource, err := webp.Decode(fSource)
	require.NoError(t, err)
	fTarget, err := os.Open("target.png")
	require.NoError(t, err)
	defer fTarget.Close()
	imgTarget, err := png.Decode(fTarget)
	require.NoError(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
