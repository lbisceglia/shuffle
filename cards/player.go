package cards

import (
	"github.com/pkg/errors"
)

// Player is an interface for entities that perform actions on Hands.
// Player actions include playing a subset of a Hand or replacing an entire Hand.
type Player interface {
	Play(h Hand)
	// TODO: Implement pick-up, handling concurrent requests on the deck from players.
	// PickUp(c Card)
	AcceptCards(h Hand)
	ReplaceHand(h Hand)
}

// NNPlayer is an implementation of Player designed to play the 99 card game.
// It plays a single card per turn.
type NNPlayer struct {
	id   int // TODO: make this a UUID
	Name string
	mgr  *NNGameManager
	hand Hand
}

// NewNNPlayer creates a 99 player with the given name.
func NewNNPlayer(name string) *NNPlayer {
	p := new(NNPlayer)
	p.Name = name
	return p
}

// selectCardAt selects the card at index i.
// It returns an error if i is invalid.
func (p *NNPlayer) selectCardAt(i int) (c Card, err error) {
	if 0 <= i && i < len(p.hand) {
		return p.hand[i], nil
	} else {
		return Card{}, errors.Wrap(err, "invalid card selected")
	}
}

// playCardAt plays the card at index i.
// An error is returned if i is invalid or the card played is invalid.
// It is the default card selection mechanism for the command line version of the game.
func (p *NNPlayer) playCardAt(i int) error {
	h := make(Hand, 0)
	c, err := p.selectCardAt(i)
	if err != nil {
		return err
	}
	h = append(h, c)
	return p.Play(h)
}

// Play submits a Player's Hand to the GameManager.
// The GameManager returns an error if the play is invalid, which is forwarded along.
func (p *NNPlayer) Play(h Hand) error {
	return p.mgr.Play(p, h)
}

// AcceptCards adds more Cards to a Player's existing Hand.
func (p *NNPlayer) AcceptCards(h Hand) {
	p.hand = append(p.hand, h...)
}

// ReplaceHand replaces a Player's existing Hand with an entirely new Hand.
func (p *NNPlayer) ReplaceHand(h Hand) {
	p.hand = h
}
