#!/bin/bash

# Generate a CA certificate and private key
openssl req -x509 -newkey rsa:4096 -keyout <ca_key_filename> -out <ca_crt_filename> -days 365 -nodes -subj "/CN=My CA"

# Generate a server certificate and private key, and sign the server certificate with the CA
openssl req -newkey rsa:4096 -keyout <server_key_filename> -out <server_csr_filename> -days 365 -nodes -subj "/CN=localhost"
openssl x509 -req -in <server_csr_filename> -CA <ca_crt_filename> -CAkey <ca_key_filename> -CAcreateserial -out <server_crt_filename> -days 365

# Generate a client certificate and private key, and sign the client certificate with the CA
openssl req -newkey rsa:4096 -keyout <client_key_filename> -out <client_csr_filename> -days 365 -nodes -subj "/CN=Client"
openssl x509 -req -in <client_csr_filename> -CA <ca_crt_filename> -CAkey <ca_key_filename> -CAcreateserial -out <client_crt_filename> -days 365

