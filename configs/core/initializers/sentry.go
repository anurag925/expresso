package initializers

import (
	"os"

	"github.com/getsentry/sentry-go"
)

func InitSentryConnection() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SentryKey"),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return err
	}
	return nil
}