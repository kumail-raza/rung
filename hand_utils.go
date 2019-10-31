package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

func (h *hand) IsComplete() bool {
	return len(h.cards) == 4
}

func (h *hand) IsEmpty() bool {
	return len(h.cards) == 0
}

func (h *hand) isSameHouse(c pattay.Card) bool {
	return c.House() == h.house
}

func (h *hand) HasCard(c pattay.Card) (bool, int) {
	_, at, err := pattay.FindCardInCards(c, h.cards)
	if err != nil {
		return false, -1
	}
	return true, at

}
func (h *hand) House() (string, error) {
	if h.IsEmpty() {
		return "", fmt.Errorf("no house because no card has been played yet")
	}
	return h.house, nil
}

func (h *hand) Cards() []pattay.Card {
	return h.cards
}
