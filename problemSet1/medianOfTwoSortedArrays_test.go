package problemSet1

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	{
		t.Log("medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2}, 1)")
		got := medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2}, 1)

		td.CmpLax(t, got, 2)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2, 4}, 2)")
		got := medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2, 4}, 2)

		td.CmpLax(t, got, 2.5)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{1}, 1, []int{2, 4}, 2)")
		got := medianOfTwoSortedArrays([]int{1}, 1, []int{2, 4}, 2)

		td.CmpLax(t, got, 2)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{2, 4}, 2, []int{2, 4}, 2)")
		got := medianOfTwoSortedArrays([]int{2, 4}, 2, []int{2, 4}, 2)

		td.CmpLax(t, got, 3)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{2, 4}, 2, []int{}, 0)")
		got := medianOfTwoSortedArrays([]int{2, 4}, 2, []int{}, 0)

		td.CmpLax(t, got, 3)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{2}, 1, []int{}, 0)")
		got := medianOfTwoSortedArrays([]int{2}, 1, []int{}, 0)

		td.CmpLax(t, got, 2)
	}
	{
		t.Log("medianOfTwoSortedArrays([]int{}, 0, []int{}, 0)")
		got := medianOfTwoSortedArrays([]int{}, 0, []int{}, 0)

		td.CmpLax(t, got, 0)
	}
}
