package main

func GenerateGameTree() *GameTree {
	return recursiveTreeGenerate(NewGame(), 0, -1)
}

func recursiveTreeGenerate(g Game, depth, position int) *GameTree {
	gt := &GameTree{
		game:            g,
		children:        make([]*GameTree, 0),
		depth:           depth,
		subtreeNodes:    1,
		subtreeLeaves:   0,
		subtreeDepth:    0,
		subtreeOutcomes: [3]int{0, 0, 0},
	}

	// Fill the position we were specified to fill
	// If the depth is even then it's O's turn
	// Position of -1 is special case for the root
	if position != -1 {
		if depth%2 == 0 {
			gt.game.state[position] = O
		} else {
			gt.game.state[position] = X
		}
	}

	// Check to see if the game is continuing.
	// Unwind recursion if finished
	if gt.game.goalTest(); gt.game.finished {
		gt.subtreeLeaves = 1
		gt.subtreeDepth = depth

		switch gt.game.winner {
		case X:
			gt.subtreeOutcomes = [3]int{1, 0, 0}
		case O:
			gt.subtreeOutcomes = [3]int{0, 1, 0}
		default:
			gt.subtreeOutcomes = [3]int{0, 0, 1}
		}

		return gt
	}

	// If we have blank spaces still then fill them
	for i, val := range gt.game.state {
		if val == None {
			child := recursiveTreeGenerate(gt.game, gt.depth+1, i)
			gt.children = append(gt.children, child)
			gt.subtreeNodes += child.subtreeNodes
			gt.subtreeLeaves += child.subtreeLeaves
			if child.subtreeDepth > gt.subtreeDepth {
				gt.subtreeDepth = child.subtreeDepth
			}
			gt.mergeOutcomes(child)
		}
	}

	return gt
}
