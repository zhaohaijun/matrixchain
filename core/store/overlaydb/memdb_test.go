package overlaydb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	db := NewMemDB(0)
	db.Put([]byte("aaa"), []byte("bbb"))
	iter := db.NewIterator(nil)
	assert.Equal(t, iter.First(), true)
	assert.Equal(t, iter.Last(), true)
	db.Delete([]byte("aaa"))
	assert.Equal(t, iter.First(), true)
	assert.Equal(t, len(iter.Value()), 0)
	assert.Equal(t, iter.Last(), true)
	assert.Equal(t, len(iter.Value()), 0)
}
