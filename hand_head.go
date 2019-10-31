package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

func (h *hand) Head() (pattay.Player, error) {
	if h.IsEmpty() {
		return nil, fmt.Errorf("no head because no card has been played yet")
	}

	return h.head, nil
}

func (h *hand) setHeadForBiggestCard(cards []pattay.Card, c pattay.Card, house string, pl pattay.Player) {

	biggestCardAtHand := pattay.GetBiggestCard(cards, house)
	if c.Number() > biggestCardAtHand.Number() {
		h.head = pl
	}

}
