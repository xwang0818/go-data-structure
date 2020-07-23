/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package graph

import "testing"
import "util"

func TestGraphStronglyConnectedComp(t *testing.T) {
    /* 0 1 2 3 4 5 6 7 8 9 10
    *
    *   0 <--- 2        4         7 --> 8
    *    \    ^       ^  \        ^     |
    *     v  /       /    v       |     v
    *      1 -----> 3 <--- 5 <--- 6 <-- 9 --> 10
    */
    input := [][]int{[]int{0,1}, []int{1,2}, []int{2,0}, []int{1,3},
                     []int{3,4}, []int{4,5}, []int{5,3}, []int{6,5},
                     []int{6,7}, []int{7,8}, []int{8,9}, []int{9,6},
                     []int{9,10}}
    result := stronglyConnectedComponents(11, input)
    expected := [][]int{[]int{7, 8, 9, 6},
                        []int{10},
                        []int{1, 2, 0},
                        []int{4, 5, 3}}
    if !util.TwoDimArraysEqual(expected, result) {
        t.Errorf("Failed to find strongly connect components")
    }
}
