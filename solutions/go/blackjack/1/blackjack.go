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
    c1p := ParseCard(card1)
    c2p := ParseCard(card2)
    dcp := ParseCard(dealerCard)
    ourCardSum := c1p + c2p
    
	switch {
        case c1p == 11 && c2p == 11:
        	return "P"
        case ourCardSum == 21 && dcp != 11 && dcp!= 10:
        	return "W"
        case ourCardSum == 21 && (dcp == 11 || dcp== 10):
        	return "S"
        case ourCardSum >= 17 && ourCardSum <= 20:
        	return "S"
        case ourCardSum >= 12 && ourCardSum <= 16 && dcp < 7:
        	return "S"
        case ourCardSum >= 12 && ourCardSum <= 16 && dcp >= 7:
        	return "H"
        case ourCardSum <= 11:
        	return "H"
        default:
        	return "S"
    }
}
