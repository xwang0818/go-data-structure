/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package sortutil

import "testing"
import "util"
import "sort"
import "reflect"

func TestSortQuickSort(t *testing.T) {
    a := []int{9,1,2,0,4,5,3,7}
    b := make([]int, len(a))
    arrLength := copy(b, a)
    if arrLength != len(a) {
        t.Errorf("Error copying array")
    }
    quickSort(a)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j]})
    if !reflect.DeepEqual(a, b) {
        t.Errorf("Failed to QuickSort")
    }
    if !util.IsArraySortedAscending(a) {
        t.Errorf("Failed to QuickSort ascendingly")
    }
}

func TestSortMergeSort(t *testing.T) {
    a := []int{9,1,2,0,4,5,3,7}
    b := make([]int, len(a))
    arrLength := copy(b, a)
    if arrLength != len(a) {
        t.Errorf("Error copying array")
    }
    mergeSort(a)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j]})
    if !reflect.DeepEqual(a, b) {
        t.Errorf("Failed to MergeSort")
    }
    if !util.IsArraySortedAscending(a) {
        t.Errorf("Failed to MergeSort ascendingly")
    }
}
