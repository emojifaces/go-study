package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {

	sl := NewSkipList(10)

	sl.Insert(4)
	sl.Insert(8)
	sl.Insert(5)
	sl.Insert(9)
	sl.Insert(2)

	assert.Equal(t, 10, sl.maxLevel)
	assert.Equal(t, 5, sl.length)

	assert.Equal(t, false, sl.Search(10)) // false
	assert.Equal(t, true, sl.Search(5))   // true

	sl.Delete(5)
	sl.Delete(4)

	assert.Equal(t, 10, sl.maxLevel)
	assert.Equal(t, 3, sl.length)

	assert.Equal(t, false, sl.Search(5))  // false
	assert.Equal(t, false, sl.Search(10)) // false
	assert.Equal(t, true, sl.Search(2))   // true

}
