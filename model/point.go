package model

type Point struct{ Row, Col int }

func (p *Point) Equal(t *Point) bool {
	return p.Row == t.Row && p.Col == t.Col
}
