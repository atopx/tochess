package model

import (
	"strings"
	"tochess/asset/image/chess"
)

const PieceLen = 57

type Piece struct {
	Camp   *PieceCamp
	Cate   *PieceCate
	Point  *Point
	Image  []byte
	ImageS []byte
	ImageM []byte
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

var (
	RedR = &Piece{Camp: Red, Cate: PieceR, Image: chess.RedR, ImageS: chess.RedRS, ImageM: nil}
	RedN = &Piece{Camp: Red, Cate: PieceN, Image: chess.RedN, ImageS: chess.RedNS, ImageM: nil}
	RedB = &Piece{Camp: Red, Cate: PieceB, Image: chess.RedB, ImageS: chess.RedBS, ImageM: nil}
	RedA = &Piece{Camp: Red, Cate: PieceA, Image: chess.RedA, ImageS: chess.RedAS, ImageM: nil}
	RedK = &Piece{Camp: Red, Cate: PieceK, Image: chess.RedK, ImageS: chess.RedKS, ImageM: chess.RedKM}
	RedC = &Piece{Camp: Red, Cate: PieceC, Image: chess.RedC, ImageS: chess.RedCS, ImageM: nil}
	RedP = &Piece{Camp: Red, Cate: PieceP, Image: chess.RedP, ImageS: chess.RedPS, ImageM: nil}

	BlackR = &Piece{Camp: Black, Cate: PieceR, Image: chess.BlackR, ImageS: chess.BlackRS, ImageM: nil}
	BlackN = &Piece{Camp: Black, Cate: PieceN, Image: chess.BlackN, ImageS: chess.BlackNS, ImageM: nil}
	BlackB = &Piece{Camp: Black, Cate: PieceB, Image: chess.BlackB, ImageS: chess.BlackBS, ImageM: nil}
	BlackA = &Piece{Camp: Black, Cate: PieceA, Image: chess.BlackA, ImageS: chess.BlackAS, ImageM: nil}
	BlackK = &Piece{Camp: Black, Cate: PieceK, Image: chess.BlackK, ImageS: chess.BlackKS, ImageM: chess.BlackKM}
	BlackC = &Piece{Camp: Black, Cate: PieceC, Image: chess.BlackC, ImageS: chess.BlackCS, ImageM: nil}
	BlackP = &Piece{Camp: Black, Cate: PieceP, Image: chess.BlackP, ImageS: chess.BlackPS, ImageM: nil}

	PieceMap = map[string]*Piece{
		RedR.Code():   RedR,
		RedN.Code():   RedN,
		RedB.Code():   RedB,
		RedA.Code():   RedA,
		RedK.Code():   RedK,
		RedC.Code():   RedC,
		RedP.Code():   RedP,
		BlackR.Code(): BlackR,
		BlackN.Code(): BlackN,
		BlackB.Code(): BlackB,
		BlackA.Code(): BlackA,
		BlackK.Code(): BlackK,
		BlackC.Code(): BlackC,
		BlackP.Code(): BlackP,
	}
)

func (p *Piece) Copy() *Piece {
	return &Piece{
		Camp:   p.Camp,
		Cate:   p.Cate,
		Point:  p.Point,
		Image:  p.Image,
		ImageS: p.ImageS,
		ImageM: p.ImageM,
	}
}

func GetPieceByCode(code string) {}
