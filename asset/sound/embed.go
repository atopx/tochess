package sound

import (
	_ "embed"
)

var (
	//go:embed bg_music.mp3
	BgMusic []byte

	//go:embed capture.mp3
	Capture []byte

	//go:embed check.mp3
	Check []byte

	//go:embed click.mp3
	Click []byte

	//go:embed draw.mp3
	Draw []byte

	//go:embed invalid.mp3
	Invalid []byte

	//go:embed lose.mp3
	Lose []byte

	//go:embed win.mp3
	Win []byte
)
