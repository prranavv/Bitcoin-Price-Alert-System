package main

import "fmt"

// sendMail is a function that sends a mail
func sendMail(alertID int) {
	fmt.Printf("Sent mail to the user regarding AlertNo-%d\n", alertID)
}
