package utils

import (
	"log"
)

func HandlePanic() {
	// detect if panic occurs or not
	if a := recover(); a != nil {
		log.Printf("panic occured %+v", a)
	}
}
