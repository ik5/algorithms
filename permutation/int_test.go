// +build ignore

package permutation

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestZeroPermutations(t *testing.T) {
	in := []int{1, 2, 3, 4}

	expected := [][]int{}

	out := IntSlice(0, in)

	if !cmp.Equal(out, expected) {
		t.Errorf("Out is not equal to expected. Returned: %+v | Expected: %+v", out, expected)
	}

}

func TestOnePermutationIntSlice(t *testing.T) {
	in := []int{1, 2, 3, 4}

	expected := [][]int{
		{1, 2, 3, 4},
	}

	out := IntSlice(1, in)

	if !cmp.Equal(out, expected) {
		t.Errorf("Out is not equal to expected. Returned: %+v | Expected: %+v", out, expected)
	}
}
