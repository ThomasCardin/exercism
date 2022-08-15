package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerHands := ParseCard(card1) + ParseCard(card2)
	dealerHands := ParseCard(dealerCard)

	decision := ""
	switch {
	case playerHands == 22:
		decision = "P"
	case playerHands == 21 && dealerHands != 11 && dealerHands != 10:
		decision = "W"
	case playerHands == 21 && (dealerHands == 11 || dealerHands == 10):
		decision = "S"
	case playerHands >= 17 && playerHands <= 20:
		decision = "S"
	case playerHands >= 12 && playerHands <= 16 && dealerHands < 7:
		decision = "S"
	case playerHands >= 12 && playerHands <= 16 && dealerHands >= 7:
		decision = "H"
	case playerHands <= 11:
		decision = "H"
	}

	return decision
}
