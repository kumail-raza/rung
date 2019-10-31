package rung

import "github.com/minhajuddinkhan/pattay"

const (
	//FirstHandForClub FirstHandForClub
	FirstHandForClub = 0
	//SecondLastHand SecondLastHand
	SecondLastHand = 11
)

func (g *game) HandsOnGround() []Hand {
	return g.handsOnGround
}

func (g *game) HandsWonBy(player pattay.Player) int {
	return g.handsWon[player.Name()]
}

func isFirstHand(turn int) bool {
	return turn == FirstHandForClub
}
func isSecondLastHand(turn int) bool {
	return turn == SecondLastHand
}
func canWinHand(turn int) bool {
	if isFirstHand(turn) {
		return false
	}
	if isSecondLastHand(turn) {
		return false
	}
	return true
}
