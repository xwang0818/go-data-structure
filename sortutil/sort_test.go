package sortutil

import "testing"
import "util"
import "sort"

func TestSortQuickSort(t *testing.T) {
    arr1 := []int{9,1,2,0,4,5,3,7}
    arr2 := make([]int, len(arr1))
    arrLength := copy(arr2, arr1)
    if arrLength != len(arr1) {
        t.Errorf("Error copying array")
    }
    quickSort(arr1)
    sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j]})
    if !util.ArraysEqual(arr1, arr2) {
        t.Errorf("Failed to QuickSort")
    }
    if !util.IsArraySortedAscending(arr1) {
        t.Errorf("Failed to QuickSort")
    }
}

func TestSortMergeSort(t *testing.T) {
    arr1 := []int{9,1,2,0,4,5,3,7}
    arr2 := make([]int, len(arr1))
    arrLength := copy(arr2, arr1)
    if arrLength != len(arr1) {
        t.Errorf("Error copying array")
    }
    mergeSort(arr1)
    sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j]})
    if !util.ArraysEqual(arr1, arr2) {
        t.Errorf("Failed to MergeSort")
    }
    if !util.ArraysEqual(arr1, arr2) {
        t.Errorf("Failed to MergeSort")
    }
    if !util.IsArraySortedAscending(arr1) {
        t.Errorf("Failed to MergeSort")
    }
}
