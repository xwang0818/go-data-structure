/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package tree

import "testing"
import "reflect"

func TestTreeConstruction(t *testing.T) {
    a := []int{1,2,3,4,5,6,7}
    root := constructTreeFromArray(a)
    b := levelOrderIterative(root)
    // level order is not exactly the same as
    // array used to construct a tree
    if !reflect.DeepEqual(a, b) {
        t.Errorf("Failed to construct tree")
    }
}
