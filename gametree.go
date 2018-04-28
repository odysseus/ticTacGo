package main

import (
	"fmt"
)

type GameTree struct {
	game            Game
	children        []*GameTree
	depth           int
	subtreeNodes    int
	subtreeLeaves   int
	subtreeDepth    int
	subtreeOutcomes [3]int
}

func (g *GameTree) String() string {
	return fmt.Sprintf(
		"Max Depth: %v -- Total Nodes: %v -- Leaves: %v -- Outcomes: %v\n",
		g.subtreeDepth, g.subtreeNodes, g.subtreeLeaves, g.subtreeOutcomes)
}

func (gt *GameTree) mergeOutcomes(other *GameTree) {
	for i, val := range other.subtreeOutcomes {
		gt.subtreeOutcomes[i] += val
	}
}
