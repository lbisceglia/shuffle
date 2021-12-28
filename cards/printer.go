package cards

import (
	"fmt"
)

// =============
// *** CARDS ***
// =============

// revealPile prints out all the remaining Cards in a Shoe.
func revealPile(name string, s Shoe) {
	fmt.Printf("%s pile: %d cards\n", name, len(s))
	fmt.Println(s)
}

// longString displays a Rank in long form.
func (r Rank) longString() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "invalid rank"
	}
}

// longString displays a Suit in long form.
func (s Suit) longString() string {
	return string(s)
}

// produceCode creates the initialization code for a collection of cards.
func produceCode(cards []Card) {
	fmt.Printf("{\n")
	for _, card := range cards {
		fmt.Printf("\tNewCard(%v, %v),\n", card.rank.longString(), card.suit.longString())
	}
	fmt.Printf("}\n")
}

// ==============
// *** DEALER ***
// ==============

// drawPile returns the remaining cards in the Dealer's draw pile.
func (d *dealer) drawPile() Shoe {
	return d.draw[d.drawIdx:]
}

// revealDeck prints out the remaining cards in the Dealer's draw pile.
func (d *dealer) revealDeck() {
	revealPile("Draw", d.drawPile())
}

// revealDiscard print out the remaining cards in the Dealer's discard pile.
func (d *dealer) revealDiscard() {
	revealPile("Discard", d.discard)
}

// revealDecks print out the remaining cards in the Dealer's draw and discard piles, respectively.
func (d *dealer) revealDecks() {
	d.revealDeck()
	d.revealDiscard()
}

// ==============
// *** PLAYER ***
// ==============

// revealHand prints out the Player's Hand.
func (p NNPlayer) revealHand() {
	fmt.Printf("%v\t(p%v): \t%v\n", p.Name, p.id, p.hand.String())
}
