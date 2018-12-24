package common

import (
	"log"
	"fmt"
)

func CheckError(err error) error {
	if err != nil {
		log.Fatal (err)
	}
	return err
}

func CheckErrorf(err error, format string, a ...interface{}) error {
	if err != nil {
		log.Fatal(err)
		log.Fatal(fmt.Sprintf(format,a ...))
	}
	return err
}

func CheckErrorln(err error, message string) error {
	if err != nil {
		log.Fatal(err)
		log.Fatal(message)
	}
	return err
}
