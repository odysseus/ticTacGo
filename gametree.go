package main

import (
	"fmt"
)

type GameTree struct {
	game         Game
	children     []*GameTree
	depth        int
	subtreeNodes int
	subtreeDepth int
}

func (g *GameTree) String() string {
	return fmt.Sprintf("Max Depth: %v -- Total Nodes: %v",
		g.subtreeDepth, g.subtreeNodes)
}

func GenerateGameTree() *GameTree {
	root := &GameTree{
		game:         NewGame(),
		children:     make([]*GameTree, 0),
		depth:        0,
		subtreeNodes: 0,
		subtreeDepth: 0,
	}

	for i, val := range root.game.state {
		if val == None {
			child := recursiveGenerate(root.game, root.depth+1, i)
			root.children = append(root.children, child)
			root.subtreeNodes += child.subtreeNodes
			root.subtreeDepth = child.subtreeDepth
		}
	}

	return root
}

func recursiveGenerate(g Game, depth, position int) *GameTree {
	gt := &GameTree{
		game:         g,
		children:     make([]*GameTree, 0),
		depth:        depth,
		subtreeNodes: 0,
		subtreeDepth: 0,
	}

	// Fill the position we were specified to fill
	// If the depth is even then it's O's turn
	if depth%2 == 0 {
		gt.game.state[position] = O
	} else {
		gt.game.state[position] = X
	}

	// Check to see if the game is continuing.
	// Unwind recursion if finished
	if gt.game.goalTest(); gt.game.finished {
		gt.subtreeNodes = 1
		gt.subtreeDepth = depth
		return gt
	}

	// If we have blank spaces still then fill them
	for i, val := range gt.game.state {
		if val == None {
			child := recursiveGenerate(gt.game, gt.depth+1, i)
			gt.children = append(gt.children, child)
			gt.subtreeNodes += child.subtreeNodes
			gt.subtreeDepth = child.subtreeDepth
		}
	}

	return gt
}
