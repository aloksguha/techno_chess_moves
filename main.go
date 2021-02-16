package main

import (
	"fmt"
	"github.com/aloksguha/techno_chess_moves/chess_moves/game"
	"github.com/aloksguha/techno_chess_moves/chess_moves/pieces"
)

func main()  {
	g := game.NewGame()
	g.Board.PrintBoard()

	pieceChoice := ""
	fmt.Print("Enter type of piece (King, Queen, Rook, Bishop, Knight, Pawn) : ")
	if n, err := fmt.Scanf("%s", &pieceChoice); err != nil || n != 1 {
		pieceChoice = "Knight"
	}
	fmt.Printf("Your choice : %s\n", pieceChoice)
	piece := pieces.GetSpecificPiece(pieceChoice)


	desiredPlace := ""
	fmt.Print("Enter type of piece on board (ex : A1, B5) : ")
	if n, err := fmt.Scanf("%s", &desiredPlace); err != nil || n != 1 {
		pieceChoice = "Knight"
	}
	fmt.Printf("Your choice : %s\n", desiredPlace)
	possibleMoves := g.Board.Place(desiredPlace, piece)
	fmt.Printf("Possible move from "+desiredPlace+" : ")
	for _, possibleMove := range possibleMoves {
		fmt.Printf(" %s ", possibleMove)
	}


	g.Board.PrintBoard()
}
