package blackjack

const BLACKJACK = 21
const WIN = "W"
const STAND = "S"
const HIT = "H"
const SPLIT = "P"

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
	default :
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == BLACKJACK
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack && dealerScore == 11 {
		return SPLIT
	} else if isBlackjack && dealerScore < 10 {
		return WIN
	}
	return STAND
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore == 20 && dealerScore == 11 {
		return STAND
	} else if handScore >= 17 {
		return STAND
	} else if handScore <= 11 {
		return HIT
	} else if dealerScore < 7 {
		return STAND
	}

	return HIT
}
