package board

import (
	"fmt"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"github.com/aloksguha/techno_chess_moves/utils"
	"strconv"
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


type Board struct {
	Name   string
	Degree int
	Cells  map[string]Cell
	cellmap map[string]string
}


func NewBoard(name string, degree int) *Board {
	if degree > 8 || degree < 1 {
		panic("Degree of board should between 2 and 8 !! Please try again")
	}

	var b Board
	b.Name = name
	b.Degree = degree
	b.Cells = make(map[string]Cell)
	b.cellmap = make(map[string]string)

	for i := 0; i < b.Degree; i++ {
		for j := 0; j < b.Degree; j++ {
			alphaIndex := string('A'-1+i+1) + strconv.Itoa(j+1)
			cell := Cell{Name: alphaIndex, Position: []int{i, j}, IsOccupied: false}
			if ((i*b.Degree)+j+(i%2))%2 == 0 {
				cell.Color = "White"
			} else {
				cell.Color = "Black"
			}
			b.Cells["R"+strconv.Itoa(i)+"C"+strconv.Itoa(j)] = cell
			b.cellmap[alphaIndex] = "R"+strconv.Itoa(i)+"C"+strconv.Itoa(j)
		}
	}
	return &b
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.Degree; i++ {
		if i == 0 {
			fmt.Print(utils.Boundry("\n————————————————————————————————\n"))
		}
		for j := 0; j < b.Degree; j++ {
			cell := b.Cells["R"+strconv.Itoa(i)+"C"+strconv.Itoa(j)]
			cell.printCell()
		}
		fmt.Print("\n")
		if i == b.Degree-1 {
			fmt.Print(utils.Boundry("—————————————————————————————————\n"))
		}
	}
}

