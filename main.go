package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"log"
	"tochess/asset/sound"
	"tochess/model"
)

type State uint

const (
	Pending State = iota
	Running
)

const (
	AuditSampleRate = 48000
)

var (
	MusicClick *mp3.Stream
	//
)

type Game struct {
	Players     [2]*model.Player
	Board       [10][9]*model.Piece
	SelectPoint *model.Point
	IsGameOver  bool // 游戏是否结束
	state       State
	Audio       *audio.Context
}

func NewGame() *Game {
	g := &Game{
		Players:    [2]*model.Player{},
		Board:      [10][9]*model.Piece{},
		IsGameOver: false,
		state:      0,
		Audio:      audio.NewContext(AuditSampleRate),
	}
	g.init()
	return g
}

func (g *Game) init() {
	var err error
	MusicClick, err = mp3.DecodeWithoutResampling(bytes.NewReader(sound.Click))
	if err != nil {
		panic(err)
	}

	// 黑色方
	g.Board[0][0] = &model.Piece{Camp: model.Black, Cate: model.PieceR}
	g.Board[0][1] = &model.Piece{Camp: model.Black, Cate: model.PieceN}
	g.Board[0][2] = &model.Piece{Camp: model.Black, Cate: model.PieceB}
	g.Board[0][3] = &model.Piece{Camp: model.Black, Cate: model.PieceA}
	g.Board[0][4] = &model.Piece{Camp: model.Black, Cate: model.PieceK}
	g.Board[0][5] = &model.Piece{Camp: model.Black, Cate: model.PieceA}
	g.Board[0][6] = &model.Piece{Camp: model.Black, Cate: model.PieceB}
	g.Board[0][7] = &model.Piece{Camp: model.Black, Cate: model.PieceN}
	g.Board[0][8] = &model.Piece{Camp: model.Black, Cate: model.PieceR}
	g.Board[2][1] = &model.Piece{Camp: model.Black, Cate: model.PieceC}
	g.Board[2][7] = &model.Piece{Camp: model.Black, Cate: model.PieceC}
	g.Board[3][0] = &model.Piece{Camp: model.Black, Cate: model.PieceP}
	g.Board[3][2] = &model.Piece{Camp: model.Black, Cate: model.PieceP}
	g.Board[3][4] = &model.Piece{Camp: model.Black, Cate: model.PieceP}
	g.Board[3][6] = &model.Piece{Camp: model.Black, Cate: model.PieceP}
	g.Board[3][8] = &model.Piece{Camp: model.Black, Cate: model.PieceP}

	// 红色方
	g.Board[9][0] = &model.Piece{Camp: model.Red, Cate: model.PieceR}
	g.Board[9][1] = &model.Piece{Camp: model.Red, Cate: model.PieceN}
	g.Board[9][2] = &model.Piece{Camp: model.Red, Cate: model.PieceB}
	g.Board[9][3] = &model.Piece{Camp: model.Red, Cate: model.PieceA}
	g.Board[9][4] = &model.Piece{Camp: model.Red, Cate: model.PieceK}
	g.Board[9][5] = &model.Piece{Camp: model.Red, Cate: model.PieceA}
	g.Board[9][6] = &model.Piece{Camp: model.Red, Cate: model.PieceB}
	g.Board[9][7] = &model.Piece{Camp: model.Red, Cate: model.PieceN}
	g.Board[9][8] = &model.Piece{Camp: model.Red, Cate: model.PieceR}
	g.Board[7][1] = &model.Piece{Camp: model.Red, Cate: model.PieceC}
	g.Board[7][7] = &model.Piece{Camp: model.Red, Cate: model.PieceC}
	g.Board[6][0] = &model.Piece{Camp: model.Red, Cate: model.PieceP}
	g.Board[6][2] = &model.Piece{Camp: model.Red, Cate: model.PieceP}
	g.Board[6][4] = &model.Piece{Camp: model.Red, Cate: model.PieceP}
	g.Board[6][6] = &model.Piece{Camp: model.Red, Cate: model.PieceP}
	g.Board[6][8] = &model.Piece{Camp: model.Red, Cate: model.PieceP}

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
			var asset string
			if g.SelectPoint != nil && row == g.SelectPoint.Row && col == g.SelectPoint.Col {
				asset = piece.SelectedImagePath()
			} else {
				asset = piece.ImagePath()
			}
			image, _, err := ebitenutil.NewImageFromFile(asset)
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
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
