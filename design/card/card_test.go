package card

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	c1 := NewCard(Hearts, Queen)
	c2 := NewCard(Spades, Jack)

	if assert.True(t, c1.Beats(c2)) {
		fmt.Println(c1.String(), "beats", c2.String())
	}

	if assert.False(t, c2.Beats(c1)) {
		fmt.Println(c2.String(), "does not beat", c1.String())
	}
}

func TestJoker(t *testing.T) {
	j1 := NewJoker(Red)
	j2 := NewJoker(Black)
	c1 := NewCard(Hearts, Ace)

	if assert.True(t, j1.Beats(c1)) {
		fmt.Println(j1.String(), "beats", c1.String())
	}

	if assert.False(t, c1.Beats(j1)) {
		fmt.Println(c1.String(), "does not beat", j1.String())
	}

	if assert.False(t, j1.Beats(j2)) {
		fmt.Println(j1.String(), "does not beat", j2.String())
	}
}
