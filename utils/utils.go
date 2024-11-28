package utils

import "log"

// FailOnError logs and exits on error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
