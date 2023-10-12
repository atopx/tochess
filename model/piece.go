package model

import (
	"fmt"
	"strings"
)

const PieceLen = 57

type Piece struct {
	Camp  *PieceCamp
	Cate  *PieceCate
	Point *Point
}

func (p *Piece) Code() string {
	// 小写表示黑方，大写表示红方
	if p.Camp == Black {
		return strings.ToLower(p.Cate.Code)
	}
	return p.Cate.Code
}

func (p *Piece) String() string {
	return p.Camp.Prefix + p.Cate.Name
}

func (p *Piece) ImagePath() string {
	return fmt.Sprintf("asset/image/chess/%s%s.GIF", p.Camp.Code, p.Cate.Code)
}

func (p *Piece) SelectedImagePath() string {
	return fmt.Sprintf("asset/image/chess/%s%sS.GIF", p.Camp.Code, p.Cate.Code)
}
