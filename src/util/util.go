/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package util

func ArraysEqual(a, b []int) bool {
    if (a == nil) != (b == nil) {
      return false
    }
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TwoDimArraysEqual(a, b [][]int) bool {
    if (a == nil) != (b == nil) {
        return false
    }
    if len(a) != len(a) {
        return false
    }
    for i := range a {
        if !ArraysEqual(a[i], b[i]) {
          return false
        }
    }
    return true
}

func IsArraySortedAscending(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            return false
        }
    }
    return true
}

func IsArraySortedDescending(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] < arr[i+1] {
            return false
        }
    }
    return true
}
