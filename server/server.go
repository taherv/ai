
import (
	"crypto/tls" // Package for TLS (Transport Layer Security) cryptographic protocol
	"crypto/x509" // Package for X.509 certificate parsing
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"config" // Import the `config` package
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	// Load the configuration using `LoadConfig` from the `config` package
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Load the server's certificate and private key from the configuration
	cert, err := tls.LoadX509KeyPair(cfg.CertificateFilenames[0], cfg.CertificateFilenames[1])
	if err != nil {
		log.Fatal(err) // Error handling when loading the server's certificate and private key
	}

	// Load the CA certificate from the configuration
	caCert, err := ioutil.ReadFile(cfg.CertificateFilenames[2])
	if err != nil {
		log.Fatal(err) // Error handling when loading the CA certificate
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a tls.Config object with the server's certificate, the CA certificate pool, and set ClientAuth to tls.RequireAndVerifyClientCert
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Listen to port 8080 and wait
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}
	log.Fatal(server.ListenAndServeTLS("", ""))
}