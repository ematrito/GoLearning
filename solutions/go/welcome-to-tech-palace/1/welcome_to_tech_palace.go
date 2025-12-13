package techpalace

import (
	"strings"
	"unicode"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcomeMessage := "Welcome to the Tech Palace, "
    capitalizeCustomer := strings.ToUpper(customer)
    return welcomeMessage + capitalizeCustomer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    
    lineOfStars := strings.Repeat("*", numStarsPerLine)
    return lineOfStars + "\n" + welcomeMsg + "\n" + lineOfStars 
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimFunc(oldMsg, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '%'
        })
}

