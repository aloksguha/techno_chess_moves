package game

import "github.com/aloksguha/techno_chess_moves/chess_moves/board"

type Game struct {
	Board *board.Board
}

/**
This method creates and returns the new instance of Game
 */
func NewGame() *Game{
	return &Game{
		Board: board.NewBoard("First Board", 8),
	}
}


