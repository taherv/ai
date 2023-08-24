package config

type Config struct {
	CertificateFilenames []string
}

func LoadConfig() (*Config, error) {
	// Read the configuration file and populate the `Config` struct
	// Return the `Config` struct and any error encountered
}

