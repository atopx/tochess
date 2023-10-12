package model

// PieceCamp 阵营
type PieceCamp struct {
	Code   string
	Name   string
	Prefix string
	Owner  *Player
}

var (
	Red   = &PieceCamp{Code: "R", Name: "红方", Prefix: "红"}
	Black = &PieceCamp{Code: "B", Name: "黑方", Prefix: "黑"}
)

func (c *PieceCamp) SetPlayer(player *Player) {
	c.Owner = player
}
