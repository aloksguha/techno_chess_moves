package pieces

type King struct { Piece }

func (p King) GetPossibleMoves() [][]int {
	return [][] int{
		{-1, 0},
		{-1, -1},
		{-1, 1},

		{0,-1},
		{0,1},

		{1, 0},
		{1, -1},
		{1, 1},
	}
}

