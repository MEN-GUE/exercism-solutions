package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer);
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine);
}

func CleanupMessage(oldMsg string) string {
	// Remove all stars
	cleanedMsg := strings.ReplaceAll(oldMsg, "*", "")
	// Trim leading and trailing whitespace
	cleanedMsg = strings.TrimSpace(cleanedMsg)
	// Return the cleaned message
	return cleanedMsg
}
