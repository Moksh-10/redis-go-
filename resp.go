package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	input := "$4\r\nECHO\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// fmt.Println(strings.NewReader(input), reader)

	b, _ := reader.ReadByte()

	fmt.Println(b)

	if b != '$' {
		fmt.Println("Expected $")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))
}