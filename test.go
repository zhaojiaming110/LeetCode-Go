// Copyright (c) 2020, beuself. All rights reserved.
// license that can be found in the LICENSE file.
// @Date: 2020/4/1 下午11:00

package main

import "fmt"

func main() {
	fmt.Println(hello(0))
}

func hello(k int) bool {
	k = k % 5
	if k == 0 {
		fmt.Println(k)
		fmt.Println("hello, world")
		return true
	}
	return false
}