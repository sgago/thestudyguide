// Package card represents the standard 52 playing cards and 2 jokers.
//
// See: https://algo.monster/problems/oop_playing_cards for details.
package card

import "fmt"

// Represents a playing card's color.
type Color string

const (
	// Represents the Red playing card color.
	Red Color = "Red"

	// Represents the Black playing card color.
	Black Color = "Black"

	// A special None color for non-Joker Other-typed cards.
	None Color = "None"
)

// String returns the friendly name of this color: "Red", "Black", or "None".
func (c Color) String() string {
	return string(c)
}

// Valid color returns true if the color is Red or Black.
// None is considered Invalid.
func ValidColor(c Color) bool {
	return c == Red || c == Black
}

// Represent's a cards Suit: Hearts, Spades, Clubs, Diamonds, and Other.
type Suit int

const (
	// Represents the Hearts suit.
	Hearts Suit = 0

	// Represents the Spades suit.
	Spades Suit = 1

	// Represents the Clubs suit.
	Clubs Suit = 2

	// Represents the Diamonds suit.
	Diamonds Suit = 3

	// Number of suits: hearts, spades, clubs, and diamonds. Does not count Other suit.
	NumSuits int = 4

	// Special Suit for representing special cards like Jokers.
	Other Suit = 5

	// Number of all suits: other, hearts, spades, clubs, and diamonds.
	NumAllSuits int = 5
)

// Int returns this Suit as an int value.
func (s Suit) Int() int {
	return int(s)
}

// Name returns this Suit's friendly name.
func (s Suit) Name() string {
	return SuitName(s)
}

// Color returns this Suit's Color.
func (s Suit) Color() Color {
	return SuitColor(s)
}

// ValidSuit returns true if s is Hearts, Spades, Clubs, or Diamonds; otherwise, false.
// Other is considered an invalid Suit.
func ValidSuit(s Suit) bool {
	return s.Int() >= Hearts.Int() && s.Int() <= Diamonds.Int()
}

// SuitName returns the friendly name of a Suit.
func SuitName(s Suit) string {
	return SuitNames()[s]
}

// SuitColor returns the Color of the Suit.
func SuitColor(s Suit) Color {
	return SuitColors()[s]
}

// Suits returns a slice of all Suits excluding Other.
func Suits() []Suit {
	return []Suit{
		Hearts,
		Spades,
		Clubs,
		Diamonds,
	}
}

// AllSuits returns a slice of all Suits including Other.
func AllSuits() []Suit {
	return append(Suits(), Other)
}

// SuitNames returns a map Suits to their friendly names.
func SuitNames() map[Suit]string {
	return map[Suit]string{
		Hearts:   "Hearts",
		Spades:   "Spades",
		Clubs:    "Clubs",
		Diamonds: "Diamonds",
		Other:    "Other",
	}
}

// SuitColors returns a map of Suits to their associated Color.
func SuitColors() map[Suit]Color {
	return map[Suit]Color{
		Hearts:   Red,
		Spades:   Black,
		Clubs:    Black,
		Diamonds: Red,
		Other:    None,
	}
}

// Represents a playing card value.
//   - Ace, 2-10, Jack, Queen, and King have positive integer values.
//   - Other cards like Joker have negative integer values.
type Value int

const (
	Ace   Value = 1
	Two   Value = 2
	Three Value = 3
	Four  Value = 4
	Five  Value = 5
	Six   Value = 6
	Seven Value = 7
	Eight Value = 8
	Nine  Value = 9
	Ten   Value = 10
	Jack  Value = 11
	Queen Value = 12
	King  Value = 13

	// A count of the number of playing cards excluding Jokers.
	NumValues int = 13

	RedJoker   Value = -1
	BlackJoker Value = -2

	// A count of the number of Jokers.
	NumJokers int = 2

	// A count of all playing cards including special Other cards like Jokers.
	NumAllValues int = NumValues + NumJokers
)

func (v Value) Int() int {
	return int(v)
}

func (v Value) Joker() bool {
	return v < Ace
}

func (v Value) PlayingCard() bool {
	return v >= Ace
}

func (v Value) Name() string {
	return ValueNames()[v]
}

func ValidValue(v Value) bool {
	return v.Int() >= Ace.Int() && v.Int() <= King.Int()
}

func (v Value) Beats(other Value) bool {
	// Jokers beat everything except other Jokers
	if v.Joker() {
		return !other.Joker()
	} else if other.Joker() {
		return false
	}

	return v.Int() > other.Int()
}

func Values() []Value {
	return []Value{
		Ace,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
	}
}

func AllValues() []Value {
	return append(
		Values(),
		RedJoker,
		BlackJoker,
	)
}

func ValueNames() map[Value]string {
	return map[Value]string{
		Ace:        "Ace",
		Two:        "Two",
		Three:      "Three",
		Four:       "Four",
		Five:       "Five",
		Six:        "Six",
		Seven:      "Seven",
		Eight:      "Eight",
		Nine:       "Nine",
		Ten:        "Ten",
		Jack:       "Jack",
		Queen:      "Queen",
		King:       "King",
		RedJoker:   "Red Joker",
		BlackJoker: "Black Joker",
	}
}

// Represents a Card with a Suit, Value, and Color.
//
// Cards are considered immutable types.
type Card struct {
	// This Card's Suit.
	suit Suit

	// This Card's Value.
	value Value

	// This Card's Color.
	color Color
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Color() Color {
	return c.color
}

func (c Card) Value() Value {
	return c.value
}

/*
	TODO:
	We should represent all 54 cards with an immutable struct type and hide the ctors.
	From sheer laziness and time saving, I'll keep the card and card ctor func
	exposed as they are, but normally I'd prefer creating
	these 54 distinct pre-canned cards following an immutable enum struct pattern
	and omit these ctor functions entirely.

	Reasoning:
	There are exactly 54 cards per the problem statement;
	there's no Bishop card, no eleven card, or anything like that;
	Allowing non-standard cards like a "bishop" card or something would be surprising here.
	We want to follow the law of least surprises.
	Put another way, if you need Uno or Magic the Gather cards,
	then they should be written in a separate package
	that concerns itself with the rules of those cards and rules.
*/

func NewCard(s Suit, v Value) Card {
	if !ValidSuit(s) {
		panic(fmt.Sprintf("the suit %d is invalid", s.Int()))
	}

	if !ValidValue(v) {
		panic(fmt.Sprintf("the value %d is invalid", v.Int()))
	}

	if v.Joker() {
		panic("use the NewJoker(c Color) ctor for new'ing up Joker cards")
	}

	return Card{
		suit:  s,
		value: v,
		color: s.Color(),
	}
}

func NewJoker(c Color) Card {
	if !ValidColor(c) {
		panic(fmt.Sprintf("the color %s is invalid", string(c)))
	}

	value := RedJoker
	if c == Black {
		value = BlackJoker
	}

	return Card{
		suit:  Other,
		value: value,
		color: c,
	}
}

func (c Card) String() string {
	if c.Joker() {
		return c.value.Name()
	}

	return fmt.Sprintf("%s (%d) of %s", c.value.Name(), c.value.Int(), c.suit.Name())
}

func (c Card) Beats(other Card) bool {
	return c.value.Beats(other.value)
}

func (c Card) Joker() bool {
	return c.value.Joker()
}

type Game struct {
	count int
	deck  map[Suit]map[Value]Card
}

func NewGame() *Game {
	cards := make(map[Suit]map[Value]Card, NumSuits+1)

	for i := 0; i < NumSuits; i++ {
		cards[Suit(i)] = make(map[Value]Card, NumValues)
	}

	cards[Other] = make(map[Value]Card, NumJokers)

	return &Game{
		deck: cards,
	}
}

func (g *Game) Count() int {
	return g.count
}

func (g *Game) Add(c Card) int {
	// Only allow new, distinct cards into our deck.
	if _, ok := g.deck[c.suit][c.value]; !ok {
		g.deck[c.suit][Value(c.suit)] = c
		g.count++
	}

	return g.count
}
