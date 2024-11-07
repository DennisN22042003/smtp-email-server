@echo off
echo Generating self-signed SSL certificate...
openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

echo Certificate and key have been generated as 'server.crt' and 'server.key'