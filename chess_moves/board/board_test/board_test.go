package board_test

import (
	"github.com/aloksguha/techno_chess_moves/chess_moves/board"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
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


	//testing for pawn
	pawn := pieces.GetSpecificPiece("Pawn")
	moves := testBoard.Place("D4", pawn) //placing pawn at middle of board so we get all possible moves
	if len(moves) != len(pawn.GetPossibleMoves()){
		t.Errorf("Expected moves should  be %d, but got %d", pawn.GetPossibleMoves(), len(moves))
	}

	possibleMarkedCells := filterPossibleMoveMarkedCells(testBoard.Cells)
	if len(possibleMarkedCells) != len(pawn.GetPossibleMoves()){
		t.Errorf("Expected number of cells should  be %d, but got %d", pawn.GetPossibleMoves(), len(possibleMarkedCells))
	}

	testBoard.ResetBoard()


	//testing for king
	king := pieces.GetSpecificPiece("King")
	movesOfKing := testBoard.Place("D4", king) //placing pawn at middle of board so we get all possible moves
	if len(movesOfKing) != len(king.GetPossibleMoves()){
		t.Errorf("Expected moves should  be %d, but got %d", king.GetPossibleMoves(), len(movesOfKing))
	}

	possibleMarkedCells = filterPossibleMoveMarkedCells(testBoard.Cells)
	if len(possibleMarkedCells) != len(king.GetPossibleMoves()){
		t.Errorf("Expected number of cells should  be %d, but got %d", king.GetPossibleMoves(), len(possibleMarkedCells))
	}

	testBoard.ResetBoard()


	//testing for knight
	knight := pieces.GetSpecificPiece("knight")
	movesOfknight:= testBoard.Place("D4", knight) //placing pawn at middle of board so we get all possible moves
	if len(movesOfknight) != len(king.GetPossibleMoves()){
		t.Errorf("Expected moves should  be %d, but got %d", king.GetPossibleMoves(), len(movesOfknight))
	}

	possibleMarkedCells = filterPossibleMoveMarkedCells(testBoard.Cells)
	if len(possibleMarkedCells) != len(knight.GetPossibleMoves()){
		t.Errorf("Expected number of cells should  be %d, but got %d", knight.GetPossibleMoves(), len(possibleMarkedCells))
	}
	testBoard.ResetBoard()

	anotherTestBoard := board.NewBoard(nameOfBoard, degreeOfBoard)
	anotherPawn := pieces.GetSpecificPiece("Pawn")
	anotherMoves := anotherTestBoard.Place("A4", anotherPawn) // placing at corner row, so moves should be 3
	if len(anotherMoves) == len(anotherPawn.GetPossibleMoves()){
		t.Errorf("Expected moves should  not be %d, but got %d", pawn.GetPossibleMoves(), len(moves))
	}

	newPossibleMarkedCells := filterPossibleMoveMarkedCells(anotherTestBoard.Cells)
	if len(newPossibleMarkedCells) != len(anotherMoves){
		t.Errorf("Expected number of cells should  be %d, but got %d", len(anotherMoves), len(newPossibleMarkedCells))
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