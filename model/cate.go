package model

// PieceCate 棋子类型
type PieceCate struct {
	Code string
	Name string
}

var (
	PieceR = &PieceCate{Code: "R", Name: "车"}
	PieceN = &PieceCate{Code: "N", Name: "马"}
	PieceB = &PieceCate{Code: "B", Name: "象"}
	PieceA = &PieceCate{Code: "A", Name: "士"}
	PieceK = &PieceCate{Code: "K", Name: "将"}
	PieceC = &PieceCate{Code: "C", Name: "炮"}
	PieceP = &PieceCate{Code: "P", Name: "兵"}
)
