package staking_service

import (
	"os"
)

var (
	PRIVATE_KEY              string
	ADDRESS                  string
	STAKING_SERVICE_BASE_URL string
)

func setupTest() {
	PRIVATE_KEY = getEnv("PRIVATE_KEY", "")
	ADDRESS = getEnv("ADDRESS", "")
	STAKING_SERVICE_BASE_URL = getEnv("STAKING_SERVICE_BASE_URL", "http://127.0.0.1:7070")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
