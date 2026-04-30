package webpbin

import (
	"bytes"
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
)

func TestEncode(t *testing.T) {
	f, err := os.Open("source.jpg")
	require.NoError(t, err)
	defer f.Close()
	imgSource, err := jpeg.Decode(f)
	require.NoError(t, err)
	var b bytes.Buffer
	err = Encode(&b, imgSource)
	assert.Nil(t, err)
	imgTarget, err := webp.Decode(bytes.NewReader(b.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
