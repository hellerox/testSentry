package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cast"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "http://6c15089b20fd4cffa0f3b100e9d817a8@192.168.2.10:9000/2",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	for i := 1; i <= 1000000000; i++ {
		time.Sleep(100 * time.Millisecond)
		sentry.CaptureException(fmt.Errorf(cast.ToString(i)))

	}
}
