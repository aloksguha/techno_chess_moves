package board_test

import (
	"github.com/aloksguha/chess_test/chess_first/board"
	"github.com/aloksguha/chess_test/chess_first/pieces"
	"testing"
)

func TestBoard (t *testing.T) {
	const nameOfBoard = "TestBoard"
	const degreeOfBoard = 8
	testBoard := board.NewBoard(nameOfBoard, degreeOfBoard)

	if testBoard.Name != nameOfBoard {
		t.Errorf("Expected Board name should be %s, but got %s", nameOfBoard, testBoard.Name)
	}

	if len(testBoard.Cells) != degreeOfBoard*degreeOfBoard {
		t.Errorf("Boards Cells should be exact square of its degree. Should %d, but got %d", degreeOfBoard*degreeOfBoard, len(testBoard.Cells))
	}


	pawn := pieces.GetSpecificPiece("Pawn")
	moves := testBoard.Place("D4", pawn) //placing pawn at middle of board so we get all possible moves
	if len(moves) != len(pawn.GetPossibleMoves()){
		t.Errorf("Expected moves should  be %d, but got %d", pawn.GetPossibleMoves(), len(moves))
	}

	possibleMarkedCells := filterPossibleMoveMarkedCells(testBoard.Cells)
	if len(possibleMarkedCells) != len(pawn.GetPossibleMoves()){
		t.Errorf("Expected number of cells should  be %d, but got %d", pawn.GetPossibleMoves(), len(possibleMarkedCells))
	}

	testBoard1 := board.NewBoard(nameOfBoard, degreeOfBoard)
	newpawn := pieces.GetSpecificPiece("Pawn")
	newmoves := testBoard1.Place("A4", newpawn) // placing at corner row, so moves should be 3
	if len(newmoves) == len(newpawn.GetPossibleMoves()){
		t.Errorf("Expected moves should  not be %d, but got %d", pawn.GetPossibleMoves(), len(moves))
	}

	newPossibleMarkedCells := filterPossibleMoveMarkedCells(testBoard1.Cells)
	if len(newPossibleMarkedCells) != len(newmoves){
		t.Errorf("Expected number of cells should  be %d, but got %d", len(newmoves), len(newPossibleMarkedCells))
	}
}


func filterPossibleMoveMarkedCells(cellsMap map[string]board.Cell ) []board.Cell {
	possibleMoveCells := make([]board.Cell, 0)
	for _, cell := range cellsMap {
		if cell.IsPossible {
			possibleMoveCells = append(possibleMoveCells, cell)
		}
	}
	return possibleMoveCells
}