/**
@author: Administrator
@Date: 2022-11-10-0010
@Note
**/
package main

import (
	"go_test/findPath"
	"math/rand"
)

func main() {
	closeChan := make(chan bool)
	for i := 0; i < 100; i++ {
		go find()
	}
	<-closeChan
}

func find() {
	start := &findPath.Cell{X: rand.Int63n(999) + 1, Y: rand.Int63n(999) + 1}
	end := &findPath.Cell{X: rand.Int63n(999) + 1, Y: rand.Int63n(999) + 1}
	findPath.QuadFindPath(*start, *end)
}
