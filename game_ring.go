package rung

import "github.com/minhajuddinkhan/pattay"

func (g *game) playRing(cardsDelt int, hand Hand) error {

	rp, err := g.ring.Next()
	player := (rp).(pattay.Player)
	if err != nil {
		return err
	}
	cardAt := player.CardOnTable()
	err = hand.AddCard(player, cardAt)
	if err != nil {
		return err
	}

	return nil
}

func (g *game) playFirstHand(hand Hand) {
	for i, p := range g.players {
		clubTwo := pattay.NewCard(pattay.Club, pattay.Two)

		if has, cardAt := p.HasCard(clubTwo); has {
			hand.AddCard(p, cardAt)
			g.players = append(g.players[:i], g.players[i+1:]...)

			g.ring.SetCurrentPlayer(p)
			break
		}
	}

}
