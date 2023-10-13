package sound

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"sync"
)

type Sound []byte

var (
	//go:embed bg_music.mp3
	backgroudSound Sound

	//go:embed capture.mp3
	captureSound Sound

	//go:embed check.mp3
	checkSound Sound

	//go:embed click.mp3
	clickSound Sound

	//go:embed move.mp3
	moveSound Sound

	//go:embed draw.mp3
	drawSound Sound

	//go:embed invalid.mp3
	invalidSound Sound

	//go:embed lose.mp3
	loseSound Sound

	//go:embed win.mp3
	winSound Sound
)

func (s Sound) Player(ctx *audio.Context) *audio.Player {
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(s))
	if err != nil {
		panic(err)
	}
	player, err := ctx.NewPlayer(stream)
	if err != nil {
		panic(err)
	}
	return player
}

var (
	MusicBackground *audio.Player
	MusicCapture    *audio.Player // 吃子
	MusicClick      *audio.Player // 点击
	MusicCheck      *audio.Player
	MusicMove       *audio.Player // 移动
	MusicDraw       *audio.Player
	MusicInvalid    *audio.Player
	MusicLose       *audio.Player
	MusicWin        *audio.Player
)

var once sync.Once

func Init(ctx *audio.Context) {
	once.Do(func() {
		MusicBackground = backgroudSound.Player(ctx)
		MusicCapture = captureSound.Player(ctx)
		MusicClick = clickSound.Player(ctx)
		MusicCheck = checkSound.Player(ctx)
		MusicMove = moveSound.Player(ctx)
		MusicDraw = drawSound.Player(ctx)
		MusicInvalid = invalidSound.Player(ctx)
		MusicLose = loseSound.Player(ctx)
		MusicWin = winSound.Player(ctx)
	})
}

func Play(player *audio.Player) {
	_ = player.Rewind()
	player.Play()
}
