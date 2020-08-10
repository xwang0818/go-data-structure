/* Copyright (C) Xiang Wang - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Xiang Wang <xwang1314@gmail.com>, July 2020
 */

package file

import (
    "bufio"
    "os"
    "log"
    "strings"
)

func ReadFile(filename string) [][]string {
    // os.O_RDONLY, os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND
    file, err := os.OpenFile(filename, os.O_RDONLY, 0770)
    if err != nil {
      log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    arr := [][]string{}
    for scanner.Scan() {
        str := scanner.Text()
        arr = append(arr, strings.Split(str, " "))
    }
    return arr
}
