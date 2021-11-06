// +build go1.7

package is

import (
	"testing"
)

// TestSubtests ensures subtests work as expected.
// https://github.com/matryer/is/issues/1
func TestSubtests(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {
		is := New(t)
		is.Equal(1+1, 2)
	})
}
