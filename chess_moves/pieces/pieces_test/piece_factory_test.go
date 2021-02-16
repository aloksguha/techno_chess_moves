package pieces_test

import (
	"fmt"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
	"reflect"
	"testing"
)

func TestPieceFactory(t *testing.T) {
	testPieces := []string {"KING", "QUEEN", "ROOK", "BISHOP", "KNIGHT", "PAWN"}

	for _, piece := range testPieces {
		fmt.Printf("Testing for Piece : %s\n", piece)

		genericPiece := pieces.Piece {
			IsActive: true,
		}

		switch {

		case piece == "KING":
			{
				genericPiece.Name = "King"
				initializedPiece := pieces.King{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece){
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}
			}

		case piece == "QUEEN":
			{
				genericPiece.Name = "QUEEN"
				initializedPiece := pieces.Queen{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece){
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}
			}

		case piece == "BISHOP":
			{
				genericPiece.Name = "BISHOP"
				initializedPiece := pieces.Bishop{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece) {
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}
			}
		case piece == "ROOK":
			{
				genericPiece.Name = "ROOK"
				initializedPiece := pieces.Rook{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece){
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}
			}
		case piece == "KNIGHT":
			{
				genericPiece.Name = "KNIGHT"
				initializedPiece := pieces.Knight{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece){
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}

			}
		case piece == "PAWN":
			{
				genericPiece.Name = "Pawn"
				initializedPiece := pieces.Pawn{Piece: genericPiece}
				generatedPiece := pieces.GetSpecificPiece(piece)
				fmt.Println(reflect.TypeOf(initializedPiece))

				if reflect.TypeOf(initializedPiece) != reflect.TypeOf(generatedPiece){
					t.Errorf("Type of piece gennerated from factory is different, it should be %s, got %s", reflect.TypeOf(initializedPiece), reflect.TypeOf(generatedPiece))
				}

			}
		}

	}


	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {

				err = fmt.Errorf("Panic Thrown : %v", r)
				fmt.Println(err)

			}
		}

	}()
	errorPronePiece := "ROBLOX"
	fmt.Printf("Testing for Piece : %s\n", errorPronePiece)
	pieces.GetSpecificPiece(errorPronePiece)
	panic("It should not get printed, otherwise panic test didnt pass")



}
