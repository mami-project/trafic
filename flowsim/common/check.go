package common

import (
	"fmt"
	"log"
)

func WarnErrorf(err error, format string, v ...interface{}) error {
	if err != nil {
		log.Printf("Warning: %v %s", err, fmt.Sprintf(format, v...))
	}
	return err
}

func WarnError(err error) error {
	if err != nil {
		log.Printf("Warning: %v\n", err)
	}
	return err
}

func FatalError(err error) error {
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return err
}

func FatalErrorf(err error, format string, a ...interface{}) error {
	if err != nil {
		log.Fatalf("Error: %v %s", err, fmt.Sprintf(format, a...))
	}
	return err
}

func FatalErrorln(err error, message string) error {
	if err != nil {
		log.Fatalf("Error: %v %s\n", err, message)
	}
	return err
}
