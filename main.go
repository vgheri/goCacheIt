package main

import (
	"fmt"
	"github.com/vgheri/goCacheIt/splay"
)

func main() {
	var tree = splay.New()
	_, err := tree.Insert("test", "myTestValue")
	if err != nil {
		panic(err)
	}
	fmt.Print("Node inserted correctly. Exiting...")
}
