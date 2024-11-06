package main

type Email struct {
	Recipient string
	Content   string
}

var emailQueue []Email

func enqueueEmail(recipient string, content string) {
	emailQueue = append(emailQueue, Email{Recipient: recipient, Content: content})
}

/* func processQueue() {
	for _, email := range emailQueue {
		// Process email delivery (e.g., send via SMTP)
	}
} */
