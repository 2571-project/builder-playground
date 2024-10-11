package main

import (
	"fmt"
	"os"

	mevboostrelay "github.com/ferranbt/builder-playground/mev-boost-relay"
)

func main() {
	// Load configuration from environment variables
	config, err := loadConfigFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Initialize MevBoostRelay
	relay, err := mevboostrelay.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create MevBoostRelay: %v\n", err)
		os.Exit(1)
	}

	// Start the relay
	if err := relay.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "MevBoostRelay stopped with error: %v\n", err)
		os.Exit(1)
	}
}

// loadConfigFromEnv initializes the Config struct from environment variables
func loadConfigFromEnv() (*mevboostrelay.Config, error) {
	apiListenAddr := getEnv("API_LISTEN_ADDR", "127.0.0.1")
	apiListenPort := getEnvAsUint64("API_LISTEN_PORT", 5555)
	apiSecretKey := getEnv("API_SECRET_KEY", "5eae315483f028b5cdd5d1090ff0c7618b18737ea9bf3c35047189db22835c48")
	beaconClientAddr := getEnv("BEACON_CLIENT_ADDR", "http://localhost:4500")
	logOutput := os.Stdout // You can modify this to load from env if needed

	return &mevboostrelay.Config{
		ApiListenAddr:    apiListenAddr,
		ApiListenPort:    apiListenPort,
		ApiSecretKey:     apiSecretKey,
		BeaconClientAddr: beaconClientAddr,
		LogOutput:        logOutput,
	}, nil
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvAsUint64 retrieves an environment variable as uint64 or returns a default value
func getEnvAsUint64(name string, defaultVal uint64) uint64 {
	if valStr, exists := os.LookupEnv(name); exists {
		var val uint64
		_, err := fmt.Sscan(valStr, &val)
		if err == nil {
			return val
		}
	}
	return defaultVal
}
