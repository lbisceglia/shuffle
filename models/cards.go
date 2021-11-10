package models

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Rank is a model class for playing card ranks.
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

// Suit is a model class for playing card suits.
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
		return "no suit found"
	}
}

// Card is a model class for playing cards.
// Cards are uniquely identified by their rank and suit.
type Card struct {
	rank rank
	suit suit
}

// String converts the card to its most compact representation, its rank and suit symbol.
// String is used primarily for command line applications, including debugging.
func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.rank, c.suit)
}

// ColourString augments a Card's string representation by colouring it.
// Cards are coloured based on their suit:
// Diamonds and Hearts are red, Clubs and Spades are white.
// ColourString is used primarily for command line applications, including debugging.
func (c Card) ColourString() string {
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
	sz := numDecks * len(suits) * len(ranks)
	s := make(Shoe, sz)
	i := 0
	for deck := 0; deck < numDecks; deck++ {
		for _, suit := range suits {
			for _, rank := range ranks {
				card := Card{
					rank: rank,
					suit: suit,
				}
				s[i] = card
				i++
			}
		}
	}
	return s
}

// Shuffle randomly shuffles a Shoe.
func (s Shoe) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(s); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			s[r], s[i] = s[i], s[r]
		}
	}
}

// String converts a Card slice to a compact, coloured representation.
func String(cards []Card, label string) string {
	str := make([]string, len(cards))
	for i, card := range cards {
		str[i] = card.ColourString()
	}
	if label != "" {
		label += ": "
	}
	return fmt.Sprintf("%s[%s]", label, strings.Join(str, ", "))
}

// String converts a Hand to a compact, coloured representation.
func (h Hand) String() string {
	return String(h, "Hand")
}

// String converts a Shoe to a compact, coloured representation.
func (s Shoe) String() string {
	return String(s, "Shoe")
}

// NewDeck creates a Shoe with exactly 1 standard deck.
func NewDeck() (s Shoe) {
	return NewShoe(1)
}

// Debugging Helper Methods

// RevealPile prints out all the remaining Cards in a Shoe.
func revealPile(name string, s Shoe) {
	fmt.Printf("%s pile: %d cards\n", name, len(s))
	fmt.Println(s)
}
