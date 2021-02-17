package board

import (
	"fmt"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"github.com/aloksguha/techno_chess_moves/utils"
	"strconv"
)

type Board struct {
	Name   string
	Degree int
	Cells  map[string]Cell
	cellmap map[string]string
}

/**
This method creates and returns the new instance of Board
*/
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

/**
This method pretty-prints the current state of board to the terminal
 */
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

/**
This method resets the board to its initial state, where all cells are empty and unoccupied
 */
func (b Board) ResetBoard(){
	for _, cellIndex := range(b.cellmap){
		cell := b.Cells[cellIndex]
		cell.IsPossible = false
		cell.IsOccupied = false
		cell.OccupiedBy = nil
		b.Cells[cellIndex] = cell
	}
}

/**
This method places a Piece to board and performs next action(s)
If it doesnt find matching cell on board, it will panic.
 */
func(b Board) Place(cellkey string, Piece pieces.TypePiece) []string{
	if cellid, ok := b.cellmap[cellkey]; ok {
		cell := b.Cells[cellid]
		cell.IsOccupied = true
		cell.OccupiedBy = Piece
		b.Cells[cellid] = cell
		return b.computePossibleMoves(cell, cell.OccupiedBy.GetPossibleMoves())
	}else{
		panic("This place doesn't exist on board !!")
	}
}


/**
This method calculates the possible move based on type of piece and its current position on board.
 */
func (b Board) computePossibleMoves(cell Cell, possibleMoves [][]int) []string{
	var PossibleMovesSlice []string
	var PossibleMovesSliceLables []string
	for i:=0; i< len(possibleMoves); i++ {
		possibleMoveTop := cell.Position[0]+possibleMoves[i][0]
		possibleMoveRight := cell.Position[1]+possibleMoves[i][1]
		PossibleMovesSlice = append(PossibleMovesSlice, "R"+strconv.Itoa(possibleMoveTop)+"C"+strconv.Itoa(possibleMoveRight))
	}

	for _, possibleMove := range PossibleMovesSlice {
		if cell, ok := b.Cells[possibleMove]; ok {
			PossibleMovesSliceLables = append(PossibleMovesSliceLables, cell.Name)
			cell.IsPossible = true
			b.Cells[possibleMove] = cell
		}
	}
	return PossibleMovesSliceLables
}




