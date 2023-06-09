package handler

import (
	"log"
)

func Error(Path string, err error) {
	if err != nil {
		log.Fatal("An error had occurred on ", Path, " showing the stack: ", err.Error())
	}
}

func MinorError(Path string, err error) {
	if err != nil {
		log.Printf("An error had occurred on %s showing the stack: %s", Path, err.Error())
	}
}