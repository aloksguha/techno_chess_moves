package board

import (
	"fmt"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"github.com/aloksguha/techno_chess_moves/utils"
)

type Cell struct {
	Name       string
	Position   []int
	IsOccupied bool
	OccupiedBy pieces.TypePiece
	Color      string
	IsPossible bool
}

func (c *Cell) printCell() {
	if c.IsOccupied {
		fmt.Print(" " + utils.Current(c.Name) + " ")
	} else if c.IsPossible {
		fmt.Print(" " + utils.Possible(c.Name) + " ")

	} else if c.Color == "White" {
		fmt.Print(" " + utils.ColorWhite(c.Name) + " ")
	} else {
		fmt.Print(" " + utils.ColorBlack(c.Name) + " ")
	}
}
