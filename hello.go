package main

import (
	"fmt"
	"my-algorithms/binaryTree"
)

func test() string {
	return "test"
}

func hello() string {
	return "hello"
}

type MyMap map[string]int

func (m *MyMap) tt() int {
	return (*m)["wocao"]
}


func main() {
	root := binaryTree.New(func(c1 interface{}, c2 interface{}) bool {
		return c1.(int) < c2.(int)
	})
	root.Insert(1)
	root.Insert(2)
	root.Insert(3)
	tmp := root.Search(3)
	fmt.Println(tmp.Node)


}
