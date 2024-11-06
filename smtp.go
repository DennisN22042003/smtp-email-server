package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
)

/* var authenticated bool */

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "220 Welcome to The Curtain Shop Email Server\r\n")

	reader := bufio.NewReader(conn)

	// Read the commands from the client
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading:", err)
			return
		}
		log.Println("Receive command:", command)

		switch {
		case strings.HasPrefix(command, "EHLO"):
			handleEHLO(conn)
		case strings.HasPrefix(command, "MAIL"):
			fmt.Fprintf(conn, "250 OK\r\n")
		case strings.HasPrefix(command, "RCPT"):
			handleRCPT(conn, command)
		case strings.HasPrefix(command, "DATA"):
			handleDATA(conn, reader)
		case strings.HasPrefix(command, "QUIT"):
			fmt.Fprintf(conn, "221 Bye\r\n")
			return
		case strings.HasPrefix(command, "AUTH"):
			handleAUTH(conn, command)
		default:
			fmt.Fprintf(conn, "500 Unknown command\r\n")
		}
		/* if !authenticated {
			fmt.Fprintf(conn, "530 Authentication Required!\r\n")
			return
		} */
	}
}

func handlePlainConnection(conn net.Conn, tlsConfig *tls.Config) {
	defer conn.Close()
	fmt.Fprintf(conn, "220 Welcome to The Curtain Shop Email Server\r\n")

	reader := bufio.NewReader(conn)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading:", err)
			return
		}
		log.Println("Received command:", command)

		switch {
		case strings.HasPrefix(command, "EHLO"):
			fmt.Fprintf(conn, "250-Hello, you can STARTTLS\r\n")
		case strings.HasPrefix(command, "STARTTLS"):
			//start TLS handshake
			fmt.Fprintf(conn, "220 Ready to start TLS\r\n")
			tlsConn := tls.Server(conn, tlsConfig)
			if err := tlsConn.Handshake(); err != nil {
				log.Printf("TLS handshake failed: %v\n", err)
				return
			}
			handleConnection(tlsConn)
			return
		default:
			fmt.Fprintf(conn, "500 Unknown command\r\n")
		}
	}
}

func handleEHLO(conn net.Conn) {
	fmt.Fprintf(conn, "250 Hello, glad to meet you\r\n")
}

func handleRCPT(conn net.Conn, command string) {
	rcptTo := strings.TrimSpace(command[5:])
	if !validateEmail(rcptTo) {
		fmt.Fprintf(conn, "550 Invalid recipient address\r\n")
	} else {
		fmt.Fprintf(conn, "250 OK\r\n")
	}
}

func handleDATA(conn net.Conn, reader *bufio.Reader) {
	fmt.Fprintf(conn, "354 End Data with <CR><LF>.<CR><LF>\r\n")
	var emailContent strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading data:", err)
			return
		}
		if line == ".\r\n" {
			break
		}
		emailContent.WriteString(line)
	}
	if isSpam(emailContent.String()) {
		fmt.Fprintf(conn, "550 Spam detected, message rejected\r\n")
		return
	}
}
