package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/gif"
	"log"
	"tochess/model"
)

type Game struct {
	Players    [2]*model.Player
	Board      [9][10]*model.Piece
	IsGameOver bool // 游戏是否结束
}

func (g *Game) Update() error {
	// TODO 动画处理

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println("选子")
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		fmt.Println("落子")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// TODO 初始资源加载绘制
	bg, _, err := ebitenutil.NewImageFromFile("asset/image/background.gif")
	if err != nil {
		panic(err)
	}
	screen.DrawImage(bg, nil)

	ba, _, err := ebitenutil.NewImageFromFile("asset/image/chess/BA.GIF")
	if err != nil {
		panic(err)
	}
	geo := ebiten.GeoM{}
	screen.DrawImage(ba, &ebiten.DrawImageOptions{
		GeoM:       geo,
		ColorScale: ebiten.ColorScale{},
		Blend:      ebiten.Blend{},
		Filter:     0,
	})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// 分辨率
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(521, 577)
	ebiten.SetWindowTitle("中国象棋")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
