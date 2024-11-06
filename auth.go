package main

import (
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

var authenticated bool
var validUsername = "dennis@thecurtainshop.com"
var validPassword = "securepassword"

func handleAUTH(conn net.Conn, command string) {
	parts := strings.Split(command, " ")
	if len(parts) < 2 || parts[1] != "PLAIN" {
		fmt.Fprintf(conn, "504 Unrecognized authentication type\r\n")
		return
	}
	authData := make([]byte, 1024)
	n, _ := conn.Read(authData)
	decoded, _ := base64.StdEncoding.DecodeString(string(authData[:n]))
	creds := strings.Split(string(decoded), "\000")
	if len(creds) == 3 && creds[1] == validUsername && creds[2] == validPassword {
		fmt.Fprintf(conn, "235 Authentication Successful\r\n")
		authenticated = true
	} else {
		fmt.Fprintf(conn, "535 Authentication Failed!\r\n")
	}
}
