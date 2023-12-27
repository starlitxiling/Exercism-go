package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
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
// func FirstTurn(card1, card2, dealerCard string) string {
// 	// panic("Please implement the FirstTurn function")
// 	if card1 == "ace" && card2 == "ace" {
// 		return "P"
// 	}
// 	if ParseCard(card1)+ParseCard(card2) == 21 {
// 		if dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" {
// 			return "W"
// 		} else {
// 			return "S"
// 		}
// 	} else if ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20 {
// 		return "S"
// 	} else if ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 {
// 		if ParseCard(dealerCard) >= 7 {
// 			return "H"
// 		} else {
// 			return "S"
// 		}
// 	} else {
// 		return "H"
// 	}
// }

func FirstTurn(card1, card2, dealerCard string) string {
	playerTotal := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerTotal == 21:
		if dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" {
			return "W"
		}
		return "S"
	case playerTotal >= 17 && playerTotal <= 20:
		return "S"
	case playerTotal >= 12 && playerTotal <= 16:
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
