package rung

import "github.com/minhajuddinkhan/pattay"

func makePlayers() []pattay.Player {
	playerNames := []string{pattay.EastPlayer, pattay.WestPlayer, pattay.NorthPlayer, pattay.SouthPlayer}
	var players []pattay.Player
	for i := 0; i < 4; i++ {
		players = append(players, pattay.NewPlayer(playerNames[i]))
	}

	return players
}

func makeRing(players []pattay.Player) pattay.Ring {
	var rp []pattay.RingPlayer
	for i := 0; i < len(players); i++ {
		rp = append(rp, players[i].(pattay.RingPlayer))
	}

	r, _ := pattay.NewRing(rp...)
	return r
}

func (g *game) Players() []pattay.Player {
	return g.players
}
