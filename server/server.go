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
 cert, err := ioutil.ReadFile(certFile)
 if err != nil {
 	log.Fatal(err)
 }
 key, err := ioutil.ReadFile(privateKeyFile)
 if err != nil {
 	log.Fatal(err)
 }

	// Create a certificate object
 certificate, err := tls.X509KeyPair(cert, key)
 if err != nil {
 	log.Fatal(err)
 }

	// Create a TLS config object
	tlsConfig := &tls.Config{
		Certificates:             []tls.Certificate{certificate},
		ClientAuth:               tls.RequireAndVerifyClientCert,
		PreferServerCipherSuites: true,
	}

	// Listen to port 8080 and wait
 log.Fatal(http.ListenAndServeTLS(":8080", certFile, privateKeyFile, nil))
}
