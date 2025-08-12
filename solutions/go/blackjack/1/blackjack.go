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
        case "ten", "jack", "queen", "king":
        	return 10
        default:
        	return 0
    }
    return 0
}

const Stand = "S"
const Hit = "H"
const Split = "P"
const Win = "W"

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    total := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    
    action := ""
	switch {
	case card1 == "ace" && card2 == "ace":
		action = Split
	case total == 21 && dealerValue < 10:
        action = Win
    case total == 21 && dealerValue >= 10:
        action = Stand
    case total <= 20 && total >= 17:
        action = Stand
    case total <= 16 && total >= 12:
        if dealerValue >= 7 {
            action = Hit
            break
        }
        action = Stand
    case total <= 11:
        action = Hit
	}
    return action
}
