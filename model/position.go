package model

type Position struct {
	X, Y float64
}

var (
	BroadPosition    = Position{X: 300, Y: 100}
	BroadMaxPosition = Position{X: BroadPosition.X + PieceLen*9, Y: BroadPosition.Y + PieceLen*10}
)
