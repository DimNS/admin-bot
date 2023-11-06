package main

import (
	"os"
	"strconv"

	"geeksonator/internal/app/geeksonator"
)

func main() {
	var (
		tgBotToken       string
		tgTimeoutSeconds = 15
		debugMode        bool
	)

	tgBotToken = os.Getenv("GEEKSONATOR_TELEGRAM_BOT_TOKEN")

	tts := os.Getenv("GEEKSONATOR_TELEGRAM_TIMEOUT_SECONDS")
	if tts != "" {
		if v, err := strconv.ParseUint(tts, 10, 64); err == nil {
			tgTimeoutSeconds = int(v)
		}
	}

	dm := os.Getenv("GEEKSONATOR_DEBUG_MODE")
	if dm != "" {
		if v, err := strconv.ParseBool(dm); err == nil {
			debugMode = v
		}
	}

	geeksonator.Start(tgBotToken, tgTimeoutSeconds, debugMode)
}
