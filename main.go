package main

import (
	"fmt"
)

func main() {
	gt := GenerateGameTree()

	for _, val := range gt.children[0].children {
		fmt.Println(&val.game)
	}
}
