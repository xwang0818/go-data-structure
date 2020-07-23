/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package sortutil

func quickSort(arr []int) {
    quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, start, end int) {
    if start >= end {
        return
    }
    pivit := end
    left, right := start, end
    for ; left < right; {
        for ; arr[left] < arr[pivit]; {
            left++
        }
        for ; arr[right] > arr[pivit]; {
            right--
        }
        if left < right {
            temp := arr[left]
            arr[left] = arr[right]
            arr[right] = temp
        }
    }
    temp := arr[left]
    arr[left] = arr[pivit]
    arr[pivit] = temp
    pivit = left
    quickSortHelper(arr, start, pivit-1)
    quickSortHelper(arr, pivit+1, end)
}


func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := (len(arr)+1)/2
    arrL := make([]int, mid)
    arrR := make([]int, len(arr)-mid)
    copy(arrL, arr[:mid])
    copy(arrR, arr[mid:len(arr)])
    left := mergeSort(arrL)
    right := mergeSort(arrR)
    for k, i, j := 0, 0, 0; i < len(left) || j < len(right); {
        if i < len(left) && j < len(right) {
            if left[i] < right[j] {
                arr[k] = left[i]
                i++
                k++
            }else {
                arr[k] = right[j]
                j++
                k++
            }
        } else if i < len(left) {

            arr[k] = left[i]
            i++
            k++
        } else if j < len(right) {
            arr[k] = right[j]
            j++
            k++
        }
    }
    return arr
}
