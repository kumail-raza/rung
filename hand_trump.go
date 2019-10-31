package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

func (h *hand) isTrumpDeclared() bool {
	return h.trump != ""
}

func (h *hand) declareTrump(c pattay.Card, pl pattay.Player) {
	h.trump = c.House()
	h.head = pl

}

func (h *hand) trumpCardsAtHand() []pattay.Card {

	var cards []pattay.Card
	for _, c := range h.cards {
		if c.House() == h.trump {
			cards = append(cards, c)
		}
	}
	return cards
}

func (h *hand) Trump() (string, error) {
	if h.trump == "" {
		return "", fmt.Errorf("trump not declared yet")
	}
	return h.trump, nil
}
