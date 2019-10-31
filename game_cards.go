package rung

func (g *game) ShuffleDeck(n int) error {
	return g.deck.Shuffle(n)
}
func (g *game) DistributeCards() error {

	cards := len(g.deck.CardsInDeck())
	playerTurn := 0
	for i := cards - 1; i >= 0; i-- {
		card, _ := g.deck.DrawCard(i)
		g.players[playerTurn].ReceiveCard(card)
		playerTurn++
		if playerTurn == 4 {
			playerTurn = 0
		}

	}

	return nil
}
