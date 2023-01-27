package myexport

import (
	"log"
	"os"
)

func MyWriteToFileJson(path string, to_write []byte) {
	var err error = os.WriteFile(path, to_write, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
