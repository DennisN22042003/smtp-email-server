#!/bin/bash
# Generate a self-signed SSL certificate

echo "Generating self-signed SSL certificate..."
openssl req -x509 -newKey rsa:2048 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

echo "Certificate and Key have been generated as 'server.crt' and 'server.key'."