package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	bytes, err := ReadFrom(strings.NewReader("from string"), 11)
	if err != nil {
		log.Fatal("read error")
	}
	fmt.Println(bytes)
}
