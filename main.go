package main

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load() // The Original .env
}

func main() {
	os.Exit(run())
}

func run() int { _ = "STUB: not implemented"; return 0 }

// Check whether driftCTL is run under Snyk CLI

// Enable colorization when driftctl is launched under snyk cli (piped)

// Handle panic and log them to sentry if error reporting is enabled

func flushSentry() { _ = "STUB: not implemented"; return }
