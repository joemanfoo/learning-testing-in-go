package services

import (
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{4, 5, 3, 2, 1, 6, 0, 9, 8, 7}
	Sort(elements)

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
