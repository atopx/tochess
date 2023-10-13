package main

import (
	"fmt"
	"testing"
)

func TestGame_fromFen(t *testing.T) {
	// rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR r - - 0 1
	fen := "rnbakabnr/9/1c5c1/p1p1p1p1p/9/9/P1P1P1P1P/1C5C1/9/RNBAKABNR r - - 0 1"
	g := NewGame()
	fmt.Println(g.Board)
	err := g.fromFen(fen)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g.Board)
}
