package pieces

import "strings"

//If it doesnt find matching Piece, it will panic.

func GetSpecificPiece(piecetype string) TypePiece {
	piecetype = strings.ToUpper(piecetype)
	genericPiece := Piece{
		IsActive: true,
	}

	switch {

	case piecetype == "KING":
		{
			genericPiece.Name = "King"
			piece := King{Piece: genericPiece}
			return piece
		}

	case piecetype == "QUEEN":
		{
			genericPiece.Name = "Queen"
			piece := Queen{Piece: genericPiece}
			return piece
		}

	case piecetype == "BISHOP":
		{
			genericPiece.Name = "Bishop"
			piece := Bishop{Piece: genericPiece}
			return piece
		}
	case piecetype == "ROOK":
		{
			genericPiece.Name = "Rook"
			piece := Rook{Piece: genericPiece}
			return piece
		}
	case piecetype == "KNIGHT":
		{
			genericPiece.Name = "Knight"
			piece := Knight{Piece: genericPiece}
			return piece
		}
	case piecetype == "PAWN":
		{
			genericPiece.Name = "Pawn"
			piece := Pawn{Piece: genericPiece}
			return piece
		}
	default:
		{
			panic("Invalid piece choice, please check input again !!")
		}
	}


}
