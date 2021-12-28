package cards

import (
	"shuffle/utils"
)

// A Dealer perform actions on Shoes.
// Actions include reording Cards, adding Cards to, and removing Cards from Shoes.
type Dealer interface {
	ReplaceShoe(numDecks int)
	Shuffle()
	DealHand(size int)
	HandleDiscard(cards []Card)
}

// dealer is an implementation of Dealer.
// It maintains a draw pile and discard pile.
// If the draw pile is exhausted during a deal, the dealer will automatically reshuffle
// the discard pile and draw from it.
type dealer struct {
	drawIdx int
	rand    Randomizer
	draw    Shoe
	discard Shoe
}

// NewDealer constructs a new dealer with the given number of decks, shuffled by default.
func NewDealer(numDecks int, rng Randomizer, shuffle ...bool) *dealer {
	d := new(dealer)
	d.rand = rng
	d.replaceShoe(numDecks)
	if len(shuffle) == 0 || shuffle[0] {
		d.Shuffle()
	}
	return d
}

// Shuffle randomly shuffles the draw pile.
func (d *dealer) Shuffle() {
	d.rand.Shuffle(d.draw)
	d.drawIdx = 0
}

// DealHand deals a number of Cards off the top of the draw pile.
// If the draw pile is exhausted during a deal, the Dealer will automatically reshuffle
// the discard pile and draw from it.
func (d *dealer) DealHand(size int) hand {
	// Calculate the size of the hand
	sz := utils.Min(size, d.drawSize()+d.discardSize())
	hand := make(hand, sz)
	if sz > 0 {
		end := utils.Min(len(d.draw), d.drawIdx+sz)
		copied := copy(hand, d.draw[d.drawIdx:end])
		d.drawIdx += copied
		if d.drawEmpty() {
			d.reshuffle()
		}
		if copied < sz {
			copy(hand[copied:], d.DealHand(sz-copied))
		}
		if d.drawEmpty() {
			d.reshuffle()
		}
	}
	return hand
}

// reshuffle shuffles the discard pile and sets it as the draw pile.
func (d *dealer) reshuffle() {
	d.draw, d.discard = d.discard, NewShoe(0)
	d.Shuffle()
}

// HandleDiscard adds the given cards to the discard pile.
func (d *dealer) HandleDiscard(cards []Card) {
	d.discard = append(d.discard, cards...)
}

// ReplaceShoe creates a new multi-deck Shoe and shuffles it.
func (d *dealer) ReplaceShoe(numDecks int) {
	d.replaceShoe(numDecks)
	d.Shuffle()
}

// replaceShoe creates a new multi-deck Shoe.
func (d *dealer) replaceShoe(numDecks int) {
	d.draw = NewShoe(numDecks)
	d.discard = NewShoe(0)
	d.drawIdx = 0
}

// drawSize returns the size of the draw pile.
func (d dealer) drawSize() int {
	return len(d.draw) - d.drawIdx
}

// drawSize returns the size of the draw pile.
func (d dealer) discardSize() int {
	return len(d.discard)
}

// drawEmpty returns true if the draw pile is empty, false otherwise.
func (d dealer) drawEmpty() bool {
	return len(d.draw) == d.drawIdx
}
