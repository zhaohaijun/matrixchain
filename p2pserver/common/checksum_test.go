package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksum(t *testing.T) {
	data := []byte{1, 2, 3}
	cs := Checksum(data)

	writer := NewChecksum()
	writer.Write(data)
	checksum2 := writer.Sum(nil)
	assert.Equal(t, cs[:], checksum2)

}
