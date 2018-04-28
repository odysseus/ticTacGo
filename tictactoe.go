package main

import (
	"fmt"
)

type Mark int

const (
	None Mark = iota
	X
	O
)

func (m Mark) String() string {
	switch m {
	case 0:
		return " "
	case 1:
		return "X"
	case 2:
		return "O"
	default:
		return fmt.Sprintf("%v", int64(m))
	}
}

type Game struct {
	state    [9]Mark
	finished bool
	winner   Mark
}

func NewGame() Game {
	return Game{
		state:    [9]Mark{None, None, None, None, None, None, None, None, None},
		finished: false,
		winner:   None,
	}
}

func (g *Game) String() string {
	var result string
	if g.goalTest(); g.finished {
		if g.winner == None {
			result = "GAME OVER: Draw!"
		} else {
			result = fmt.Sprintf("GAME OVER: %v is the winner!", g.winner)
		}
	} else {
		result = "STATUS: The game is not finished"
	}

	return fmt.Sprintf(
		"\n     |     |     \n"+
			"  %v  |  %v  |  %v  \n"+
			"     |     |     \n"+
			"-----------------\n"+
			"     |     |     \n"+
			"  %v  |  %v  |  %v  \n"+
			"     |     |     \n"+
			"-----------------\n"+
			"     |     |     \n"+
			"  %v  |  %v  |  %v  \n"+
			"     |     |     \n"+
			"\n%v\n",
		g.state[0], g.state[1], g.state[2],
		g.state[3], g.state[4], g.state[5],
		g.state[6], g.state[7], g.state[8],
		result)
}

func (g *Game) move(player Mark, index int) {
	g.state[index-1] = player
}

func (g *Game) hasEmptySpaces() bool {
	for _, val := range g.state {
		if val == None {
			return true
		}
	}

	return false
}

func (g *Game) goalTest() {
	rows := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {6, 4, 2}}

	for _, row := range rows {
		if g.state[row[0]] == None {
			break
		} else if g.state[row[0]] == g.state[row[1]] &&
			g.state[row[1]] == g.state[row[2]] {
			g.finished = true
			g.winner = g.state[row[0]]
			return
		}
	}

	if !g.hasEmptySpaces() {
		g.finished = true
	}
}
