package go_set

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSet(t *testing.T) {
  	s1 := Set{[]int{1, 1, 2, 2, 3}}
	assert.Equal(t, len(s1.Elements()), 3)
}
