package helpers

import (
	"log"
	"time"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func UtilsNowUtc() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
}
