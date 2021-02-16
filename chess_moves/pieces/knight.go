package pieces

type Knight struct {Piece}

func (p Knight) GetPossibleMoves() [][]int {
	return [][] int{
		{-2, -1},
		{-2, 1},
		{-1, -2},
		{-1, 2},

		{1, -2},
		{2, -1},
		{2, 1},
		{1, 2},
	}
}

