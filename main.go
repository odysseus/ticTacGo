package main

import (
	"fmt"
)

func main() {
	g := NewGame()

	g.move(X, 1)
	g.move(X, 2)
	g.move(O, 3)
	g.move(O, 4)
	g.move(X, 5)
	g.move(X, 6)
	g.move(O, 7)
	g.move(O, 8)
	g.move(O, 8)
	g.move(X, 9)

	fmt.Println(g)
}
