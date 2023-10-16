package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"strconv"
	"strings"
	"tochess/asset/sound"
	"tochess/model"
)

type Game struct {
	Players           [2]*model.Player
	Board             [10][9]*model.Piece
	SelectPoint       *model.Point
	IsGameOver        bool // 游戏是否结束
	Audio             *audio.Context
	Rounds            int // 当前回合数
	MoveNum           int // 行棋步数
	WithoutEatMoveNum int // 连续没有吃子的步数
	CurrentCamp       *model.PieceCamp
}

func (g *Game) toFen() string {
	var rows []string
	var fen strings.Builder
	for _, pieces := range g.Board {
		var emptySeries int
		for _, piece := range pieces {
			if piece == nil {
				emptySeries++
			} else {
				if emptySeries > 0 {
					fen.WriteString(strconv.Itoa(emptySeries))
					emptySeries = 0
				}
				fen.WriteString(piece.Code())
			}
		}
		if emptySeries > 0 {
			fen.WriteString(strconv.Itoa(emptySeries))
		}
		rows = append(rows, fen.String())
		fen.Reset()
	}
	fen.WriteString(strings.Join(rows, "/"))
	fen.WriteString(" ")
	fen.WriteString(strings.ToLower(g.CurrentCamp.Code)) // 行棋阵营
	fen.WriteString(" - - ")
	fen.WriteString(strconv.Itoa(g.WithoutEatMoveNum))
	fen.WriteString(" ")
	fen.WriteString(strconv.Itoa(g.Rounds)) // 回合数
	return fen.String()
}

func (g *Game) fromFen(fen string) (err error) {
	fens := strings.Split(fen, " ")
	if len(fens) != 6 {
		return errors.New("invalid fen, it should be divided into 6 parts of data")
	}

	// 阵营
	if strings.ToLower(fens[1]) == "b" {
		g.CurrentCamp = model.Black
	} else {
		g.CurrentCamp = model.Red
	}

	// 没有行棋的步数
	g.WithoutEatMoveNum, err = strconv.Atoi(fens[4])
	if err != nil {
		return errors.New("invalid fen, without eating number is a int")
	}

	// 回合数
	g.Rounds, err = strconv.Atoi(fens[5])
	if err != nil {
		return errors.New("invalid fen, rounds is a int")
	}

	// 棋盘
	rows := strings.Split(fens[0], "/")
	if len(rows) != 10 {
		return errors.New("invalid fen, chinese chess should be 10 rows")
	}
	for row, pieces := range rows {
		var col int
		for _, code := range pieces {
			code := string(code)
			num, err := strconv.Atoi(code)
			if err != nil {
				if piece := model.GetPieceByCode(code); piece != nil {
					g.Board[row][col] = piece
					col++
				} else {
					return fmt.Errorf("invalid fen, unknown row: %d, code: %s", row, code)
				}
			} else {
				col += num
			}
		}
	}
	return nil
}

func NewGame() *Game {
	g := &Game{
		Players:     [2]*model.Player{},
		Board:       [10][9]*model.Piece{},
		IsGameOver:  false,
		Audio:       audio.NewContext(48000),
		CurrentCamp: model.Red,
		Rounds:      1,
	}
	return g
}

func (g *Game) init() {

	// 加载声音
	sound.Init(g.Audio)

	// 黑色方
	g.Board[0][0] = model.BlackR
	g.Board[0][1] = model.BlackN
	g.Board[0][2] = model.BlackB
	g.Board[0][3] = model.BlackA
	g.Board[0][4] = model.BlackK
	g.Board[0][5] = model.BlackA
	g.Board[0][6] = model.BlackB
	g.Board[0][7] = model.BlackN
	g.Board[0][8] = model.BlackR
	g.Board[2][1] = model.BlackC
	g.Board[2][7] = model.BlackC
	g.Board[3][0] = model.BlackP
	g.Board[3][2] = model.BlackP
	g.Board[3][4] = model.BlackP
	g.Board[3][6] = model.BlackP
	g.Board[3][8] = model.BlackP

	// 红色
	g.Board[9][0] = model.RedR
	g.Board[9][1] = model.RedN
	g.Board[9][2] = model.RedB
	g.Board[9][3] = model.RedA
	g.Board[9][4] = model.RedK
	g.Board[9][5] = model.RedA
	g.Board[9][6] = model.RedB
	g.Board[9][7] = model.RedN
	g.Board[9][8] = model.RedR
	g.Board[7][1] = model.RedC
	g.Board[7][7] = model.RedC
	g.Board[6][0] = model.RedP
	g.Board[6][2] = model.RedP
	g.Board[6][4] = model.RedP
	g.Board[6][6] = model.RedP
	g.Board[6][8] = model.RedP
	fmt.Println(g.toFen())
}

func (g *Game) position(row, col int) model.Position {
	return model.Position{
		X: model.BroadPosition.X + float64(col*model.PieceLen),
		Y: model.BroadPosition.Y + float64(row*model.PieceLen),
	}
}

func (g *Game) getBoardPiece(point *model.Point) *model.Piece {
	return g.Board[point.Row][point.Col]
}

func (g *Game) setBoardPiece(point *model.Point, piece *model.Piece) {
	g.Board[point.Row][point.Col] = piece
}

func (g *Game) selected(x, y float64) *model.Point {
	if x < model.BroadPosition.X || x > model.BroadMaxPosition.X {
		return nil
	}
	if y < model.BroadPosition.Y || y > model.BroadMaxPosition.Y {
		return nil
	}
	return &model.Point{Row: int((y - model.BroadPosition.Y) / model.PieceLen), Col: int((x - model.BroadPosition.X) / model.PieceLen)}
}

func (g *Game) Update() error {
	// TODO 动画处理
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		target := g.selected(float64(x), float64(y))
		if g.SelectPoint != nil {
			// 落子
			piece := g.getBoardPiece(g.SelectPoint) // 原位置的棋子
			if target != nil && piece != nil {
				targetPiece := g.getBoardPiece(target) // 目标位置棋子

				// 如果选择的棋子和当前游标相同, 则取消选择
				if g.SelectPoint.Equal(target) {
					g.SelectPoint = nil
					return nil
				}

				// TODO 计算是否可以行子
				if targetPiece != nil && targetPiece.Camp == piece.Camp {
					// 不能吃自己的棋子
					sound.Play(sound.MusicInvalid)
					return nil
				}

				g.setBoardPiece(g.SelectPoint, nil) // 设置原位置为空
				g.setBoardPiece(target, piece)      // 设置新位置为选择的棋子

				if targetPiece != nil {
					// 吃子
					sound.Play(sound.MusicCapture)
					g.WithoutEatMoveNum = 0
				} else {
					// 非吃子
					g.WithoutEatMoveNum++
					sound.Play(sound.MusicMove)
				}

				g.MoveNum++ // 行棋步数+1

				// 交换行棋方
				switch g.CurrentCamp {
				case model.Red:
					g.CurrentCamp = model.Black
					if g.MoveNum > 1 {
						// 非首次行棋且是红色方行棋，回合数+1
						g.Rounds++
					}
				case model.Black:
					g.CurrentCamp = model.Red
				}
				fmt.Println(g.toFen())
			}
			g.SelectPoint = nil // 行棋后清空选择棋子
		} else {
			// 选子
			if target != nil && g.getBoardPiece(target) != nil {
				g.SelectPoint = target
				sound.Play(sound.MusicClick)
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	g.drawBroad(screen)
	g.drawPiece(screen)
}

func (g *Game) drawPiece(screen *ebiten.Image) {
	for row, pieces := range g.Board {
		for col, piece := range pieces {
			if piece == nil {
				continue
			}
			var asset []byte
			if g.SelectPoint != nil && row == g.SelectPoint.Row && col == g.SelectPoint.Col {
				asset = piece.ImageS
			} else {
				asset = piece.Image
			}
			image, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(asset))
			if err != nil {
				panic(err)
			}
			geo := ebiten.GeoM{}
			offset := g.position(row, col)
			geo.Translate(offset.X, offset.Y)
			screen.DrawImage(image, &ebiten.DrawImageOptions{GeoM: geo})
		}
	}
}

func (g *Game) drawBroad(screen *ebiten.Image) {
	geo := ebiten.GeoM{}
	geo.Translate(model.BroadPosition.X, model.BroadPosition.Y)
	image, _, err := ebitenutil.NewImageFromFile("asset/image/background.gif")
	if err != nil {
		panic(err)
	}
	screen.DrawImage(image, &ebiten.DrawImageOptions{GeoM: geo})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// 分辨率
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1200, 800)
	ebiten.SetWindowTitle("中国象棋")
	game := NewGame()
	game.init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
