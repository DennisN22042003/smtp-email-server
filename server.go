package main

import (
	"crypto/tls"
	"log"
	"net"
)

func startServer() {

	// Load TLS certificate and key
	cert, err := tls.LoadX509KeyPair(config.Server.CertFile, config.Server.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the TLS config
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	// Start the plaintext (unencrypted) listener
	go func() {
		listener, err := net.Listen("tcp", ":"+config.Server.PlainPort)
		if err != nil {
			log.Fatalf("Error starting plaintext listener: %v", err)
		}
		defer listener.Close()
		log.Printf("SMTP  server listening on plaintext port %s\n", config.Server.PlainPort)
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			go handlePlainConnection(conn, tlsConfig)
		}
	}()

	// Start the TLS listener
	tlsListener, err := tls.Listen("tcp", ":"+config.Server.TLSPort, tlsConfig)
	if err != nil {
		log.Fatalf("Error starting TLS listener: %v", err)
	}
	defer tlsListener.Close()
	log.Printf("SMTP server listening on TLS port %s\n", config.Server.TLSPort)
	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}
