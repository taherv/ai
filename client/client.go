

import (
	"crypto/tls" // Package for TLS (Transport Layer Security) cryptographic protocol
	"crypto/x509" // Package for X.509 certificate parsing
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"config" // Import the `config` package
)

func main() {
	// Load the configuration using `LoadConfig` from the `config` package
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Load the client's certificate and private key from the configuration
	cert, err := tls.LoadX509KeyPair(cfg.CertificateFilenames[0], cfg.CertificateFilenames[1])
	if err != nil {
		log.Fatal(err) // Error handling when loading the client's certificate and private key
	}

	// Load the CA certificate from the configuration
	caCert, err := ioutil.ReadFile(cfg.CertificateFilenames[2])
	if err != nil {
		log.Fatal(err) // Error handling when loading the CA certificate
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