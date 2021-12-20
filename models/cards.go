package models

import (
	"fmt"
	"math/rand"
	"shuffle/utils"
	"strings"

	"github.com/fatih/color"
)

// rank is a model class for playing card ranks.
type rank string

const (
	Ace   rank = "A"
	Two   rank = "2"
	Three rank = "3"
	Four  rank = "4"
	Five  rank = "5"
	Six   rank = "6"
	Seven rank = "7"
	Eight rank = "8"
	Nine  rank = "9"
	Ten   rank = "10"
	Jack  rank = "J"
	Queen rank = "Q"
	King  rank = "K"
)

var ranks = [13]rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// String converts the rank to its recognizable, symbolic representation
func (r rank) String() string {
	switch r {
	case Ace, Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King:
		return string(r)
	default:
		return "invalid rank"
	}
}

// suit is a model class for playing card suits.
type suit string

const (
	Clubs    suit = "Clubs"
	Diamonds suit = "Diamonds"
	Hearts   suit = "Hearts"
	Spades   suit = "Spades"
)

var suits = [4]suit{Clubs, Diamonds, Hearts, Spades}

// String converts the suit to its recognizable, symbolic representation.
func (s suit) String() string {
	switch s {
	case Clubs:
		return "\u2663" // ♣
	case Diamonds:
		return "\u2666" // ♦
	case Hearts:
		return "\u2665" // ♥
	case Spades:
		return "\u2660" // ♠
	default:
		return "invalid suit"
	}
}

// card is a model class for playing cards.
// cards are uniquely identified by their rank and suit.
type card struct {
	rank rank
	suit suit
}

// NewCard constructs a card with the given rank and suit.
func NewCard(r rank, s suit) card {
	return card{
		rank: r,
		suit: s,
	}
}

// String converts the card to its most compact representation, its rank and suit symbol.
// String is used primarily for command line applications, including debugging.
func (c card) String() string {
	r, s := c.rank.String(), c.suit.String()
	if strings.HasPrefix(r, "invalid") || strings.HasPrefix(s, "invalid") {
		return fmt.Sprintf("%s, %s", r, s)
	}
	return fmt.Sprintf("%s%s", r, s)
}

// colourString augments a Card's string representation by colouring it.
// Cards are coloured based on their suit:
// Diamonds and Hearts are red, Clubs and Spades are white.
// colourString is used primarily for command line applications, including debugging.
func (c card) colourString() string {
	switch c.suit {
	case Clubs, Spades:
		white := color.New(color.FgWhite).SprintfFunc()
		return white(c.String())
	case Diamonds, Hearts:
		red := color.New(color.FgRed).SprintFunc()
		return red(c.String())
	default:
		return fmt.Sprintln("no card found")
	}
}

// hand is a collection of cards that can be held by a Player.
type hand []card

// shoe is a collection of cards that can be managed by a Dealer.
// It contains an integer number of standard card decks.
type shoe []card

// NewShoe creates a Shoe composed of an integer number of standard card decks.
func NewShoe(numDecks int) shoe {
	sz := utils.Max(numDecks*len(suits)*len(ranks), 0)
	s := make(shoe, sz)
	i := 0
	for deck := 0; deck < numDecks; deck++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				s[i] = NewCard(rank, suit)
				i++
			}
		}
	}
	return s
}

// shuffle randomly shuffles a Shoe.
func (s shoe) shuffle() {
	for i := 0; i < len(s); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			s[r], s[i] = s[i], s[r]
		}
	}
}

// String converts a Card slice to a compact, coloured representation.
func String(cards []card, label string) string {
	str := make([]string, len(cards))
	for i, card := range cards {
		str[i] = card.colourString()
	}
	if label != "" {
		label += ": "
	}
	return fmt.Sprintf("%s[%s]", label, strings.Join(str, ", "))
}

// String converts a Hand to a compact, coloured representation.
func (h hand) String() string {
	return String(h, "Hand")
}

// String converts a Shoe to a compact, coloured representation.
func (s shoe) String() string {
	return String(s, "Shoe")
}

// Debugging Helper Methods

// revealPile prints out all the remaining Cards in a Shoe.
func revealPile(name string, s shoe) {
	fmt.Printf("%s pile: %d cards\n", name, len(s))
	fmt.Println(s)
}
