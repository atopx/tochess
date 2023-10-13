package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
	"tochess/asset/sound"
	"tochess/model"
)

const (
	AuditSampleRate = 48000
)

var (
	MusicBg      *mp3.Stream
	MusicCapture *mp3.Stream
	MusicClick   *mp3.Stream
	MusicCheck   *mp3.Stream
	MusicDraw    *mp3.Stream
	MusicInvalid *mp3.Stream
	MusicLose    *mp3.Stream
	MusicWin     *mp3.Stream
)

type Game struct {
	Players       [2]*model.Player
	Board         [10][9]*model.Piece
	SelectPoint   *model.Point
	IsGameOver    bool // 游戏是否结束
	Audio         *audio.Context
	rounds        int // 当前回合数
	withoutEatNum int // 连续没有吃子的步数
	CurrentCamp   *model.PieceCamp
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
		if emptySeries == 9 {
			fen.WriteString("9")
		} else if emptySeries > 0 {
			fen.WriteString(strconv.Itoa(emptySeries))
		}
		rows = append(rows, fen.String())
		fen.Reset()
	}
	fen.WriteString(strings.Join(rows, "/"))
	fen.WriteString(" ")
	fen.WriteString(strings.ToLower(g.CurrentCamp.Code)) // 行棋阵营
	fen.WriteString(" - - ")
	fen.WriteString(strconv.Itoa(g.withoutEatNum))
	fen.WriteString(" ")
	fen.WriteString(strconv.Itoa(g.rounds)) // 回合数
	return fen.String()
}

func (g *Game) fromFen(fen string) (err error) {
	fens := strings.Split(fen, " ")
	if len(fens) != 6 {
		return errors.New("invalid fen, it should be divided into 6 parts of data")
	}
	// 回合数
	g.rounds, err = strconv.Atoi(fens[5])
	if err != nil {
		return errors.New("invalid fen, rounds is a int")
	}
	// 没有行棋的步数
	g.withoutEatNum, err = strconv.Atoi(fens[4])
	if err != nil {
		return errors.New("invalid fen, without eating number is a int")
	}
	// 阵营
	if strings.ToLower(fens[1]) == "b" {
		g.CurrentCamp = model.Black
	} else {
		g.CurrentCamp = model.Red
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
				if piece, ok := model.PieceMap[code]; ok {
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
		Audio:       audio.NewContext(AuditSampleRate),
		CurrentCamp: model.Red,
		rounds:      1,
	}
	return g
}

func (g *Game) init() {
	var err error
	MusicClick, err = mp3.DecodeWithoutResampling(bytes.NewReader(sound.Click))
	if err != nil {
		panic(err)
	}

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
	os.Exit(0)
}

func (g *Game) position(row, col int) model.Position {
	return model.Position{
		X: model.BroadPosition.X + float64(col*model.PieceLen),
		Y: model.BroadPosition.Y + float64(row*model.PieceLen),
	}
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
		if g.SelectPoint != nil {
			// 落子
			x, y := ebiten.CursorPosition()
			target := g.selected(float64(x), float64(y))
			if target != nil {
				// TODO 计算是否可以行子

				// 如果不可以, 提示音效, 跳过行棋

				// 行棋
				piece := g.Board[g.SelectPoint.Row][g.SelectPoint.Col]
				g.Board[g.SelectPoint.Row][g.SelectPoint.Col] = nil
				g.Board[target.Row][target.Col] = piece
				music, err := g.Audio.NewPlayer(MusicClick)
				if err != nil {
					panic(err)
				}
				music.Play()
			}
			//MusicClick.
			g.SelectPoint = nil
		} else {
			// 选子
			x, y := ebiten.CursorPosition()
			g.SelectPoint = g.selected(float64(x), float64(y))

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
