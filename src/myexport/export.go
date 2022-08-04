package myexport

import (
	"io/ioutil"
)

func MyWriteToFileJson(path string, to_write []byte) {
	ioutil.WriteFile(path, to_write, 0644)
}
