package cards

import (
	"fmt"
	"shuffle/utils"
	"strings"

	"github.com/fatih/color"
)

// Rank is a model class for playing card ranks.
type Rank string

const (
	Ace   Rank = "A"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
)

var Ranks = [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// String converts the rank to its recognizable, symbolic representation
func (r Rank) String() string {
	switch r {
	case Ace, Two, Three, Four, Five, Six,
		Seven, Eight, Nine, Ten, Jack, Queen, King:
		return string(r)
	default:
		return "invalid rank"
	}
}

// Suit is a model class for playing card suits.
type Suit string

const (
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
)

var Suits = [4]Suit{Clubs, Diamonds, Hearts, Spades}

// String converts the suit to its recognizable, symbolic representation.
func (s Suit) String() string {
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

// Card is a model class for playing cards.
// cards are uniquely identified by their rank and suit.
type Card struct {
	rank Rank
	suit Suit
}

// NewCard constructs a card with the given rank and suit.
func NewCard(r Rank, s Suit) Card {
	return Card{
		rank: r,
		suit: s,
	}
}

// String converts the card to its most compact representation, its rank and suit symbol.
// String is used primarily for command line applications, including debugging.
func (c Card) String() string {
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
func (c Card) colourString() string {
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

// Hand is a collection of cards that can be held by a Player.
type Hand []Card

// Shoe is a collection of cards that can be managed by a Dealer.
// It contains an integer number of standard card decks.
type Shoe []Card

// NewShoe creates a Shoe composed of an integer number of standard card decks.
func NewShoe(numDecks int) Shoe {
	sz := utils.Max(numDecks*len(Suits)*len(Ranks), 0)
	s := make(Shoe, sz)
	i := 0
	for deck := 0; deck < numDecks; deck++ {
		for _, suit := range Suits {
			for _, rank := range Ranks {
				s[i] = NewCard(rank, suit)
				i++
			}
		}
	}
	return s
}

// String converts a Card slice to a compact, coloured representation.
func String(cards []Card) string {
	str := make([]string, len(cards))
	for i, card := range cards {
		str[i] = card.colourString()
	}
	return fmt.Sprintf("[%s]", strings.Join(str, ", "))
}

// String converts a Hand to a compact, coloured representation.
func (h Hand) String() string {
	return String(h)
}

// String converts a Shoe to a compact, coloured representation.
func (s Shoe) String() string {
	return String(s)
}
