package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		os.Exit(1)
	}

	input := make([]string, length)

	for i := 0; i < length; i++ {
		scanner.Scan()
		input[i] = scanner.Text()
	}

	for j := length - 1; j >= 0; j-- {
		fmt.Print(input[j] + " ")
	}
}
