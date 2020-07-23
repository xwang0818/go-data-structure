/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package tree

import "testing"
import "util"

func TestTreeConstruction(t *testing.T) {
    arr1 := []int{1,2,3,4,5,6,7}
    root := constructTreeFromArray(arr1)
    arr2 := levelOrderIterative(root)
    // level order is not exactly the same as
    // array used to construct a tree
    if !util.ArraysEqual(arr1, arr2) {
        t.Errorf("Failed to construct tree")
    }
}
