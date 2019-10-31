package rung

import (
	"github.com/minhajuddinkhan/pattay"
)

//Game a game of court piece
type Game interface {
	//Players returns the players
	Players() []pattay.Player
	//DistributeCards distrubutes card among players of the game
	DistributeCards() error

	//PlayHand begins play the hand
	PlayHand(turn int, trump *string, lastHead pattay.Player) (Hand, error)

	//ShuffleDeck shuffes the deck n times
	ShuffleDeck(n int) error

	//HandsOnGround returns the hands on ground that are not won yet.
	HandsOnGround() []Hand

	//HandsWonBy returns the number of hands won by a player
	HandsWonBy(player pattay.Player) int
}

type game struct {
	players       []pattay.Player
	deck          pattay.Deck
	handsOnGround []Hand
	handsWon      map[string]int
	ring          pattay.Ring
}

//NewGame NewGame
func NewGame() Game {

	deck := pattay.NewDeck()
	players := makePlayers()
	r := makeRing(players)

	return &game{
		ring:     r,
		players:  players,
		deck:     deck,
		handsWon: make(map[string]int, 4),
	}
}

func (g *game) PlayHand(turn int, trump *string, lastHead pattay.Player) (Hand, error) {

	hand := NewHand(trump)
	cardsDelt := 0

	if isFirstHand(turn) {
		g.playFirstHand(hand)
		cardsDelt++
	}

	if cardsDelt == 0 {
		g.ring.SetCurrentPlayer(lastHead)
	}

	for i := 0; i < 4-cardsDelt; i++ {
		if err := g.playRing(cardsDelt, hand); err != nil {
			return nil, err
		}
	}

	g.handsOnGround = append(g.handsOnGround, hand)
	head, err := hand.Head()
	if err != nil {
		return nil, err
	}
	g.ring.SetCurrentPlayer(head)

	if !canWinHand(turn) {
		return hand, nil
	}

	if head.Name() == lastHead.Name() {
		g.handsWon[lastHead.Name()] = len(g.handsOnGround)
		g.handsOnGround = nil
	}
	return hand, nil

}
