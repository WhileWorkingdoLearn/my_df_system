package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Parse(buff []byte) (int, bool) {
	sep := bytes.Index(buff, []byte{':'})
	end := false
	if sep != -1 {
		if sep < len(buff)-1 && buff[sep+1] == byte(':') {
			end = true
		}
		return sep, end
	}

	return 0, end
}

func TestParse(t *testing.T) {
	b := []byte("aabb::c")
	parsed, end := Parse(b)
	assert.Equal(t, 4, parsed)
	assert.True(t, end)

	b = []byte("aabb:c")
	parsed, end = Parse(b)
	assert.Equal(t, 4, parsed)
	assert.False(t, end)

	b = []byte("aabbc")
	parsed, end = Parse(b)
	assert.Equal(t, 0, parsed)
	assert.False(t, end)

	b = []byte("a:b:bc")
	parsed, end = Parse(b)
	assert.Equal(t, 1, parsed)
	assert.False(t, end)
}
