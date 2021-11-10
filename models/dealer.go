package models

import (
	"shuffle/utils"
)

// IDealer is an interface for entities that perform actions on Shoes.
// Actions include reording Cards, adding Cards to, and removing Cards from Shoes.
type IDealer interface {
	InitializeDeck()
	Shuffle()
	DealHand(size int)
	HandleDiscard(cards []Card)
}

// Dealer is an implementation of IDealer.
// It maintains a draw pile and discard pile.
// If the draw pile is exhausted during a deal, the Dealer will automatically reshuffle
// the discard pile and draw from it.
type Dealer struct {
	drawIdx int
	draw    Shoe
	discard Shoe
}

// Shuffle randomly shuffles the draw pile.
func (d *Dealer) Shuffle() {
	d.draw.shuffle()
	d.drawIdx = 0
}

// DealHand deals a number of Cards off the top of the draw pile.
// If the draw pile is exhausted during a deal, the Dealer will automatically reshuffle
// the discard pile and draw from it.
func (d *Dealer) DealHand(size int) Hand {
	// Calculate the size of the hand
	sz := utils.Min(size, len(d.discard)+(len(d.draw)-d.drawIdx))
	hand := make(Hand, sz)
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

// Reshuffle shuffles the discard pile and sets it as the draw pile.
func (d *Dealer) reshuffle() {
	d.draw, d.discard = d.discard, NewShoe(0)
	d.Shuffle()
}

// HandleDiscard adds the given cards to the discard pile.
func (d *Dealer) HandleDiscard(cards []Card) {
	d.discard = append(d.discard, cards...)
}

// InitializeDeck creates a new multi-deck Shoe and shuffles it.
func (d *Dealer) InitializeDeck(numDecks int) {
	d.draw = NewShoe(numDecks)
	d.discard = NewShoe(0)
	d.Shuffle()
}

// Debugging Helper Methods

// DrawPile returns the remaining cards in the Dealer's draw pile.
func (d Dealer) drawPile() Shoe {
	return d.draw[d.drawIdx:]
}

// RevealDeck prints out the remaining cards in the Dealer's draw pile.
func (d Dealer) revealDeck() {
	revealPile("Draw", d.drawPile())
}

// RevealDiscard print out the remaining cards in the Dealer's discard pile.
func (d Dealer) revealDiscard() {
	revealPile("Discard", d.discard)
}

// RevealDecks print out the remaining cards in the Dealer's draw and discard piles, respectively.
func (d Dealer) RevealDecks() {
	d.revealDeck()
	d.revealDiscard()
}
