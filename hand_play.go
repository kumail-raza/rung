package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

func (h *hand) AddCard(pl pattay.Player, cardAtHandIndex int) error {

	c, err := pl.DrawCard(cardAtHandIndex)
	if err != nil {
		return err
	}

	if err = h.validateCard(pl, c); err != nil {
		return err
	}
	defer func() {
		h.hasPlayed = append(h.hasPlayed, pl)
		h.cards = append(h.cards, c)
	}()

	if h.IsEmpty() {
		h.house = c.House()
		h.head = pl
		return nil
	}

	if !h.isTrumpDeclared() {

		if h.isSameHouse(c) {
			h.setHeadForBiggestCard(h.cards, c, h.house, pl)
			return nil
		}

		if pl.HasHouse(h.house) {
			return fmt.Errorf("player has cards of the same house")
		}

		h.declareTrump(c, pl)
		return nil

	}

	//if trump is running at hand
	if h.trump == h.house {
		if !h.isSameHouse(c) {
			return nil
		}
		h.setHeadForBiggestCard(h.cards, c, h.trump, pl)
		return nil
	}
	if c.House() == h.trump {
		trumpCards := h.trumpCardsAtHand()
		if len(trumpCards) == 0 {
			h.head = pl
			return nil
		}
		h.setHeadForBiggestCard(h.trumpCardsAtHand(), c, h.trump, pl)
	}

	return nil

}
