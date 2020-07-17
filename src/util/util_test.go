package util

import "testing"

func TestUtilArraySortedAsc(t *testing.T) {
    arr := []int{0,1,2,3,4,5,7}
    if !IsArraySortedAscending(arr) {
        t.Errorf("Failed array is not sorted ascending")
    }
}

func TestUtilArraySortedDesc(t *testing.T) {
    arr := []int{7,4,3,0}
    if !IsArraySortedDescending(arr) {
        t.Errorf("Failed array is not sorted descending")
    }
}
