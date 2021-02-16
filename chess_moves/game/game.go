package game

import "github.com/aloksguha/techno_chess_moves/chess_moves/board"

type Game struct {
	Board *board.Board
}

func NewGame() *Game{
	return &Game{
		Board: board.NewBoard("First Board", 8),
	}
}


