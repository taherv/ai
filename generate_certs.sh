#!/bin/bash

# Generate a CA certificate and private key
openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_crt.pem -days 365 -nodes -subj "/CN=My CA"

# Generate a server certificate and private key, and sign the server certificate with the CA
openssl req -newkey rsa:4096 -keyout server_key.pem -out server_csr.pem -days 365 -nodes -subj "/CN=localhost"
openssl x509 -req -in server_csr.pem -CA ca_crt.pem -CAkey ca_key.pem -CAcreateserial -out server_crt.pem -days 365

# Generate a client certificate and private key, and sign the client certificate with the CA
openssl req -newkey rsa:4096 -keyout client_key.pem -out client_csr.pem -days 365 -nodes -subj "/CN=Client"
openssl x509 -req -in client_csr.pem -CA ca_crt.pem -CAkey ca_key.pem -CAcreateserial -out client_crt.pem -days 365

