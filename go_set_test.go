package go_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s1 := NewSet([]int{1, 1, 2, 2, 3})
	assert.Equal(t, len(s1.Val()), 3)

	s2 := NewSet([]int{1, 4, 5, 4})
	assert.Equal(t, (*s1).Equals(*s2), false)

	s3 := NewSet([]int{1, 2, 3})
	assert.Equal(t, (*s1).Equals(*s3), true)

	s1.Update(*s2)
	assert.Equal(t, len(s1.Val()), 5)
}
