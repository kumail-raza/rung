package rung

import (
	"github.com/minhajuddinkhan/pattay"
)

//Hand Round of a card
type Hand interface {
	//Cards returns the list of cards on a hand
	Cards() []pattay.Card

	//AddCard adds a card at the current hand
	AddCard(playedBy pattay.Player, cardAtHandIndex int) error

	//HasCard checks if hand has card
	HasCard(c pattay.Card) (hasCard bool, atIndex int)

	//IsComplete returns if a hand is complete or not
	IsComplete() bool

	//Head returns the player who has thrown the biggest card
	Head() (pattay.Player, error)

	//House returns the House/Color of the hand being played
	House() (string, error)

	//Trump returns the trump house of the hand being played
	Trump() (string, error)
}
type hand struct {
	cards     []pattay.Card
	hasPlayed []pattay.Player
	house     string
	head      pattay.Player
	trump     string
}

//NewHand creates a new hand
func NewHand(trump *string) Hand {
	if trump != nil {
		return &hand{trump: *trump}
	}
	return &hand{}
}
