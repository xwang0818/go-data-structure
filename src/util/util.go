package util

func ArraysEqual(arrOne []int, arrTwo []int) bool {
    if len(arrOne) != len(arrTwo) {
        return false
    }
    for i := 0; i < len(arrOne); i++ {
        if arrOne[i] != arrTwo[i] {
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
