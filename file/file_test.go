/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package file

import "testing"
import "reflect"

func TestFileRead(t *testing.T) {
    a := [][]string{{"one"}, {"two", "three"}, {""}, {"four", "five", "six"}}
    b := ReadFile("test.txt")
    if !reflect.DeepEqual(a, b) {
        t.Errorf("Failed to read from file")
    }
}
