package pieces

type Pawn struct {Piece}

func (p Pawn) GetPossibleMoves() [][]int {
	return [][] int{
		//if playing from black side
		{-1, 0},
		{-1, -1},
		{-1, 1},

		//if playing from black side
		{1, 0},
		{1, -1},
		{1, 1},
	}
}

