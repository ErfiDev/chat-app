package tools

import "os"

func LoadFile(n string) *os.File {
	file, err := os.OpenFile(n, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return file
}
