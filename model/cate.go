package model

// PieceCate 棋子类型
type PieceCate struct {
	Code string
	Name string
}

var (
	PiaceR = &PieceCate{Code: "R", Name: "车"}
	PiaceN = &PieceCate{Code: "N", Name: "马"}
	PiaceB = &PieceCate{Code: "B", Name: "象"}
	PiaceA = &PieceCate{Code: "A", Name: "士"}
	PiaceK = &PieceCate{Code: "K", Name: "将"}
	PiaceC = &PieceCate{Code: "C", Name: "炮"}
	PiaceP = &PieceCate{Code: "P", Name: "兵"}
)
