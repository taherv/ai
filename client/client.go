
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Load the client's certificate and private key
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	// Load the CA certificate
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a tls.Config object with the client's certificate, the CA certificate pool, and set InsecureSkipVerify to false
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}

	// Request /hello over port 8080 via the GET method
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	r, err := client.Get("https://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}
