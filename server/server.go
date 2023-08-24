package main

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Read the certificate and private key files
	cert, _ := ioutil.ReadFile("cert.pem")
	key, _ := ioutil.ReadFile("private.key")

	// Create a certificate object
	certificate, _ := tls.X509KeyPair(cert, key)

	// Create a TLS config object
	tlsConfig := &tls.Config{
		Certificates:             []tls.Certificate{certificate},
		ClientAuth:               tls.RequireAndVerifyClientCert,
		PreferServerCipherSuites: true,
	}

	// Listen to port 8080 and wait
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "private.key", nil))
}
