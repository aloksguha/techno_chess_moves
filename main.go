package main

import game "github.com/aloksguha/techno_chess_moves/chess_moves/game"

func main()  {
	g := game.NewGame()
	g.Board.PrintBoard()
}
