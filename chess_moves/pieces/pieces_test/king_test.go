package pieces_test

import (
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"testing"
)

func TestKing(t *testing.T) {
	const PossibleMoves = 8                 // King can have maximum possible 8 moves
	const PossibleMoveCoordinatesLength = 2 //2d board
	genericPiece := pieces.Piece {
		IsActive: true,
	}

	genericPiece.Name = "King"
	piece := pieces.King {Piece: genericPiece}

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

		if move[0] > 1 || move[0] < -1 {
			t.Errorf("King should not move than one cell in one move")
		}

		if move[1] > 1 || move[1] < -1 {
			t.Errorf("King should not move than one cell in one move")
		}

	}
}
