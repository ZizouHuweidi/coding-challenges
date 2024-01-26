package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		log.Fatal("Invalid type, expecting bulk strings only")
	}

	size, _ := reader.ReadByte()
	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// consume \r\n
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)

	reader.Read(name)

	log.Println(string(name))
}
