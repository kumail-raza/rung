package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

func (h *hand) validateCard(pl pattay.Player, c pattay.Card) error {

	if h.IsComplete() {
		return fmt.Errorf("hand is complete")
	}

	if h.HasAlreadyPlayed(pl) {
		return fmt.Errorf("player %s has already played", pl.Name())
	}

	for _, card := range h.cards {
		if pattay.IsSameCard(card, c) {
			return fmt.Errorf("one hand cannot have two same cards")
		}
	}
	return nil

}

func (h *hand) HasAlreadyPlayed(pl pattay.Player) bool {

	for _, player := range h.hasPlayed {
		if player.Name() == pl.Name() {
			return true
		}
	}
	return false
}
