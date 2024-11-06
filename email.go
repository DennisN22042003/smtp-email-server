package main

import (
	"log"
	"os"
	"strings"
)

func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func isSpam(content string) bool {
	spamKeywords := []string{"free", "win", "money", "urgent"}
	for _, keyword := range spamKeywords {
		if strings.Contains(strings.ToLower(content), keyword) {
			return true
		}
	}
	return false
}

func saveEmail(content string) {
	file, err := os.OpenFile("emails.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	if _, err := file.WriteString(content + "\n"); err != nil {
		log.Println("Error writing to file:", err)
	}
}
