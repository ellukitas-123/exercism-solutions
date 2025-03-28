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
	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	valDealer := ParseCard(dealerCard)

	sum := val1 + val2
	switch {
		case val1 == 11 && val2 == 11:
			return "P"
		case sum == 21:
			if valDealer >= 10 {
				return "S"
			}
			return "W"
		case sum >= 17 && sum <= 20:
			return "S"
		case sum >= 12 && sum <= 16:
			if valDealer >= 7 {
				return "H"
			}
			return "S"
		default:
			return "H"
	}
}
