#!/bin/bash

# Generate a CA certificate and private key
openssl req -x509 -newkey rsa:4096 -keyout ca.key -out ca.crt -days 365 -nodes -subj "/CN=My CA"

# Generate a server certificate and private key, and sign the server certificate with the CA
openssl req -newkey rsa:4096 -keyout server.key -out server.csr -days 365 -nodes -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365

# Generate a client certificate and private key, and sign the client certificate with the CA
openssl req -newkey rsa:4096 -keyout client.key -out client.csr -days 365 -nodes -subj "/CN=Client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365

