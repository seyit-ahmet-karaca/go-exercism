package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starts := strings.Repeat("*", numStarsPerLine)
	return starts + "\n" + welcomeMsg + "\n" + starts
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oneLineString := strings.ReplaceAll(oldMsg, "\n", "")
	withOutStars := strings.ReplaceAll(oneLineString, "*","")
	return strings.Trim(withOutStars, " ")
}
