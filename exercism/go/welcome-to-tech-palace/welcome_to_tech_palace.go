package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsPerLine := strings.Repeat("*", numStarsPerLine)
	welcomeMessage := starsPerLine + "\n" + welcomeMsg + "\n" + starsPerLine
	return welcomeMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	newMsg = strings.TrimSpace(newMsg)
	return newMsg
}
