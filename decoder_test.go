package webpbin

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/webp"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("source.webp")
	require.NoError(t, err)
	defer f.Close()
	imgSource, err := Decode(f)
	assert.Nil(t, err)
	f.Seek(0, 0)
	imgTarget, err := webp.Decode(f)
	assert.Nil(t, err)
	assert.Equal(t, imgSource.Bounds(), imgTarget.Bounds())
}
