package pieces_test

import (
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"testing"
)

func TestQueen(t *testing.T) {
	const PossibleMoves = 56                 // Queen can have maximum possible 28 moves
	const PossibleMoveCoordinatesLength = 2 //2d board
	genericPiece := pieces.Piece {
		IsActive: true,
	}

	genericPiece.Name = "Queen"
	piece := pieces.Queen {Piece: genericPiece}

	if piece.Name != genericPiece.Name {
		t.Errorf("Expected name should be %s, but got %s",genericPiece.Name, piece.Name)
	}

	moves := piece.GetPossibleMoves()
	if len(moves) != PossibleMoves {
		t.Errorf("Expected name should be %d, but got %d", PossibleMoves, len(moves))
	}

	for i:=0; i<len(moves); i++ {
		move := moves[i]
		if len(move) != PossibleMoveCoordinatesLength {
			t.Errorf("Expected name should be %d, but got %d", PossibleMoveCoordinatesLength, len(move))
		}

		if move[0] == 0 && move[1] == 0 {
			t.Errorf("Piece should move from original place when asked to move")
		}

		if move[0] > 7 || move[0] < -7 {
			t.Errorf("Queen should not move than seven cell in one move")
		}

		if move[1] > 7 || move[1] < -7 {
			t.Errorf("Queen should not move than seven cell in one move")
		}

	}
}
