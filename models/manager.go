package models

import (
	"fmt"
	"shuffle/utils"
	"strconv"

	"github.com/pkg/errors"
)

// // FIXME: enforce good design by exporting constructors instead of types where appropriate.

// NNPlayerStats is a model class that stores metadata about NNPlayers relevant to game play.
// This class is used and maintained by the NNGameManager.
type NNPlayerStats struct {
	player *NNPlayer
	lives  int
}

// NNWildCards specifies all relevant wild cards for a game of 99.
type NNWildCards struct {
	Reverse    rank
	NinetyNine rank
	MinusTen   rank
	Zero       rank
}

// NNGameSettings holds the relevant settings for a game of 99.
type NNGameSettings struct {
	CardsPerPlayer int
	LivesPerPlayer int
	MaxCount       int
	WildCards      NNWildCards
	// TODO: add AutoPickup as a setting (eventually needed for AI players)
}

// NNDefaultSettings are the "house rules" for a game of 99.
var NNDefaultSettings = &NNGameSettings{
	CardsPerPlayer: 3,
	LivesPerPlayer: 3,
	MaxCount:       99,
	WildCards: NNWildCards{
		Reverse:    Four,
		NinetyNine: Nine,
		MinusTen:   Ten,
		Zero:       King,
	},
}

// IGameManager is an interface for entities that specify and enforce the rules of a specific game.
// It contains the logic for starting and ending games, maintains the order of play,
// validates player moves, and shifts cards between players and dealers.
type IGameManager interface {
	NewGame()
	StartGame(humans []*NNPlayer, robots int, settings *NNGameSettings)
	Deal()
	Play(p *Player, h hand)
	EndGame()
}

// NNGameManager is an implementation of IGameManager that plays 99.
// It maintains the overall count and the rules for scoring cards that are played.
// It keeps track of the order of play and ensures that players only play valid cards, during their turn.
type NNGameManager struct {
	settings   *NNGameSettings
	players    map[int]NNPlayerStats
	dealer     *dealer
	playing    bool
	round      int
	count      int
	currPlayer int
	direction  int
	// TODO: Add a concurrency-safe field to keep track of which player can draw from the deck. Needed for AutoPickup.
}

// NewGame begins a game of 99 with a number of players. The game follows the default house rules.
func (mgr *NNGameManager) NewGame(players []*NNPlayer) {
	mgr.newGame(players, NNDefaultSettings)
}

// TODO: refactor as a server that can connect to multiple clients.

// NewGame begins a game of 99 with a number of players and custom rules.
// It currently only supports the command line version of 99.
func (mgr *NNGameManager) newGame(players []*NNPlayer, settings *NNGameSettings) {
	fmt.Println("Welcome to 99!")
	mgr.StartGame(players, 0, settings)

	for mgr.playing {
		mgr.revealTable() // TODO: temporary, remove after debugging
		player := mgr.CurrPlayer()
		fmt.Printf("%v, it's your turn. Select a card from 1-%v\n", player.Name, len(player.hand))
		var card int
		for {
			fmt.Printf("Count: %v\n", mgr.count)
			_, err := fmt.Scanf("%d", &card)
			if err != nil || card <= 0 || card > len(player.hand) {
				fmt.Println("Please make a valid selection.")
			} else {
				// TODO: re-route control flow in a more elegant way
				break
			}
		}
		err := player.playCardAt(card - 1)
		if err != nil {
			handlePlayError(err)
		}
	}
	// TODO: add multiple rounds and take advantage of players' lives
	// TODO: announce results of each round once concluded
}

// StartGame initializes and begins a new game of 99.
// It may create games that have both human and AI players.
func (mgr *NNGameManager) StartGame(humans []*NNPlayer, robots int, settings *NNGameSettings) {
	mgr.setSettings(settings)
	mgr.players = make(map[int]NNPlayerStats)
	for i, human := range humans {
		human.id = i
		human.mgr = mgr
		mgr.players[i] = NNPlayerStats{
			player: human,
			lives:  mgr.settings.LivesPerPlayer,
		}
	}
	// TODO: support AI (robot) players
	mgr.dealer = new(dealer)
	mgr.Deal()
	mgr.round = 0
	mgr.currPlayer = 0
	mgr.direction = 1
	mgr.playing = true
}

// SetSettings installs custom rules if provided, else defaults to house rules.
func (mgr *NNGameManager) setSettings(settings *NNGameSettings) {
	if settings != nil {
		mgr.settings = settings
	} else {
		mgr.settings = NNDefaultSettings
	}
}

// Deal initializes the Dealer with an appropriately-sized shoe and deals cards to each player.
// It also deals one card face up to begin the game.
func (mgr *NNGameManager) Deal() {
	set := mgr.settings
	decks := MinDecks(set.CardsPerPlayer, set.WildCards, len(mgr.players))
	mgr.dealer.InitializeDeck(decks)
	for _, stat := range mgr.players {
		hand := mgr.dealer.DealHand(set.CardsPerPlayer)
		stat.player.ReplaceHand(hand)
	}
	h := mgr.dealer.DealHand(1)
	initCount, _ := mgr.ScoreCard(h[0])
	mgr.count = initCount
	mgr.dealer.HandleDiscard(h)
}

// MinDecks calculates the minimum number of decks to use in the Shoe, to avoid endless games of 99.
// It uses cards per player, designated wild cards, and the number of players to recommend a shoe size.
func MinDecks(cardsEach int, wilds NNWildCards, numPlayers int) int {
	// TODO: implement
	return 1
}

// Play validates a player's move, scores their card, and advances play to the next player.
func (mgr *NNGameManager) Play(p *NNPlayer, h hand) (err error) {
	if mgr.CurrPlayer().id != p.id {
		return errors.New("playing out of turn")
	} else if c, ok := mgr.getCardFromPlayer(p, h[0]); !ok {
		return errors.New("cheating")
	} else {
		if toAdd, err := mgr.ScoreCard(c); err != nil {
			return errors.Wrap(err, "invalid card")
		} else {
			mgr.count += toAdd
			mgr.dealer.HandleDiscard(h)
		}
		if mgr.count > mgr.settings.MaxCount {
			mgr.DeclareLoser(p)
			// TODO: re-route control flow in a more elegant way
			return
		}
		mgr.reverseIfNeeded(c)
		hand := mgr.dealer.DealHand(1)
		p.AcceptCards(hand)
		mgr.AdvanceCurrPlayer()
	}
	return err
}

// GetCardFromPlayer removes a Card from a player's hand, if it exists.
// It returns true and the Card if it exists, else false and an empty Card.
func (mgr *NNGameManager) getCardFromPlayer(p *NNPlayer, c card) (card, bool) {
	if v, ok := mgr.players[p.id]; ok {
		for i, card := range v.player.hand {
			if card == c {
				v.player.hand = append(v.player.hand[:i], v.player.hand[i+1:]...)
				return c, true
			}
		}
	}
	return card{}, false
}

// DeclareLoser announces the loser of the round and ends the game.
func (mgr *NNGameManager) DeclareLoser(p *NNPlayer) {
	// TODO: remove the *NNPlayer arg and just use currPlayer
	fmt.Printf("%d points busts %d! %v loses!\n", mgr.count, mgr.settings.MaxCount, p.Name)
	// TODO: decrement lives instead of just ending the game
	mgr.EndGame()
}

// ScoreCard determines the effect of the card on the count.
func (mgr NNGameManager) ScoreCard(c card) (toAdd int, err error) {
	return mgr.ScoreRank(c.rank)
}

// ScoreRank determines the effect of the rank on the count.
// Assumes count-altering wild cards (Zero, NinetyNine, MinusTen), are all different ranks,
// as one Rank card cannot have multiple competing effects on the Count.
// Assumes Reverse has no impact on the count, unless it also happens to be a wild card.
func (mgr *NNGameManager) ScoreRank(r rank) (toAdd int, err error) {
	set := mgr.settings
	wilds := set.WildCards
	switch r {
	case wilds.NinetyNine:
		toAdd = set.MaxCount - mgr.count
	case wilds.MinusTen:
		toAdd = -10
	case wilds.Zero, wilds.Reverse:
		// no change to the count
	case Ace:
		toAdd = 1
	case Jack, Queen, King:
		toAdd = 10
	default:
		toAdd, err = strconv.Atoi(string(r))
	}
	return toAdd, err
}

// ReverseIfNeeded reverses the direction of play if a Reverse card is played.
func (mgr *NNGameManager) reverseIfNeeded(c card) {
	if c.rank == mgr.settings.WildCards.Reverse {
		mgr.direction *= -1
	}
}

// EndGame ends the game.
func (mgr *NNGameManager) EndGame() {
	// TODO: end game only if there's one winner standing, declare winner
	mgr.playing = false
	fmt.Println("Thanks for playing!")
	mgr.revealTable()
}

// AdvanceCurrPlayer advances play to the next player in the circle, according to the direction of play.
func (mgr *NNGameManager) AdvanceCurrPlayer() {
	if curr, err := utils.Mod(mgr.currPlayer+mgr.direction, len(mgr.players)); err != nil {
		mgr.currPlayer = curr
	} else {
		mgr.currPlayer = 0
	}
}

// CurrPlayer returns the player who is currently taking their turn.
func (mgr *NNGameManager) CurrPlayer() *NNPlayer {
	return mgr.players[mgr.currPlayer].player
}

// HandlePlayError handles invalid plays attempted by Players during gameplay.
func handlePlayError(e error) {
	// TODO: implement robust error handling
	fmt.Println(errors.Cause(e))
}

// Debugging Helper Methods

// RevealTable shows all cards in the game, including the draw pile, discard pile,
// and every player's hand.
func (mgr NNGameManager) revealTable() {
	fmt.Printf("Count: %v\n", mgr.count)
	for _, stat := range mgr.players {
		stat.player.revealHand()
	}
	mgr.dealer.revealDecks()
}
