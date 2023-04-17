package handler

import "log"

func Error(Path string, err error) {
	if err != nil {
		log.Fatal("An error had occurred on ", Path, " showing the stack: ", err.Error())
	}
}
