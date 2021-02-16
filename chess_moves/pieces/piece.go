package pieces


type TypePiece interface {
	GetPossibleMoves() [][]int
}

type Piece struct {
	Name string
	IsActive bool
}

func (p *Piece) GetName() string {
	return p.Name
}


func (p *Piece) GetActiveStatus() bool {
	return p.IsActive
}

func (p *Piece) SetActiveStatus(isActive bool) {
	p.IsActive = isActive
}



func (p *Piece) GetPossibleMoves() [][]int{
	panic("This is abstract method, must provide implementation !!")
}

