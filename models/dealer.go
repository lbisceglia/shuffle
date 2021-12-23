package models

import (
	"shuffle/utils"
)

// Dealer is an interface for entities that perform actions on Shoes.
// Actions include reording Cards, adding Cards to, and removing Cards from Shoes.
type Dealer interface {
	ReplaceShoe(numDecks int)
	Shuffle()
	DealHand(size int)
	HandleDiscard(cards []card)
}

// dealer is an implementation of IDealer.
// It maintains a draw pile and discard pile.
// If the draw pile is exhausted during a deal, the dealer will automatically reshuffle
// the discard pile and draw from it.
type dealer struct {
	drawIdx int
	draw    shoe
	discard shoe
}

// NewDealer constructs a new dealer with the given number of decks, shuffled by default.
func NewDealer(numDecks int, shuffle ...bool) *dealer {
	d := new(dealer)
	d.replaceShoe(numDecks)
	if len(shuffle) == 0 || shuffle[0] {
		d.Shuffle()
	}
	return d
}

// Shuffle randomly shuffles the draw pile.
func (d *dealer) Shuffle() {
	d.draw.shuffle()
	d.drawIdx = 0
}

// DealHand deals a number of Cards off the top of the draw pile.
// If the draw pile is exhausted during a deal, the Dealer will automatically reshuffle
// the discard pile and draw from it.
func (d *dealer) DealHand(size int) hand {
	// Calculate the size of the hand
	sz := utils.Min(size, len(d.discard)+(len(d.draw)-d.drawIdx))
	hand := make(hand, sz)
	if len(d.draw) == 0 {
		d.reshuffle()
	}
	if len(d.draw) > 0 {
		end := utils.Min(len(d.draw), d.drawIdx+sz)
		copied := copy(hand, d.draw[d.drawIdx:end])
		d.drawIdx += copied
		if copied < sz {
			d.reshuffle()
			copy(hand[copied:], d.DealHand(sz-copied))
			d.drawIdx = sz - copied
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
func (d *dealer) HandleDiscard(cards []card) {
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

// Debugging Helper Methods

// DrawPile returns the remaining cards in the Dealer's draw pile.
func (d dealer) drawPile() shoe {
	return d.draw[d.drawIdx:]
}

// RevealDeck prints out the remaining cards in the Dealer's draw pile.
func (d dealer) revealDeck() {
	revealPile("Draw", d.drawPile())
}

// RevealDiscard print out the remaining cards in the Dealer's discard pile.
func (d dealer) revealDiscard() {
	revealPile("Discard", d.discard)
}

// revealDecks print out the remaining cards in the Dealer's draw and discard piles, respectively.
func (d dealer) revealDecks() {
	d.revealDeck()
	d.revealDiscard()
}
