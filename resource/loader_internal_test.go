package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_ShouldHashStringWithMd5(t *testing.T) {
	date := "Fri, 02 Feb 1996 03:04:05 GMT"
	str := "123.jpg"

	hashed := hash(str, date)

	assert.Equal(t, "ff686fed68bb8f5d8e679fbe8f794bba", hashed)
}

func TestFilename_ShouldFindLastPieceOfUrlstring(t *testing.T) {
	str := "https://is2-ssl.mzstatic.com/image/thumb/Features4/v4/e7/f0/71/e7f07160-bb0f-496b-fc11-e1c29e36162a/source/3200x600w.png"
	piece := filename(str)

	assert.Equal(t, "3200x600w.png", piece)
}
