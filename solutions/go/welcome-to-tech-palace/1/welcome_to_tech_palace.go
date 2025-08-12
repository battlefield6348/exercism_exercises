package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg ,stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var content string
    lines := strings.Split(oldMsg, "\n")
	if len(lines) >= 2 {
		content = strings.TrimSpace(lines[1])
		content = strings.Trim(content, "* ")
	}
	return content
}
