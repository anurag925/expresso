package utils

import (
	"context"
	"time"
)

func TimedContext(contextTimeOut ...int) (context.Context, context.CancelFunc) {
	// contextTimeOut := os.Getenv("CONTEXT_TIME_OUT")
	var timeout int
	if len(contextTimeOut) == 0 {
		timeout = 5
	} else {
		timeout = contextTimeOut[0]
	}
	return context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
}
