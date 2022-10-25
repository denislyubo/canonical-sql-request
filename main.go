package main

import (
	"bufio"
	t "github.com/denislyubo/canonical-sql-request/tokenizer"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Printf("Failed to read: %v", scanner.Err())
		return
	}
	data := scanner.Bytes()

	tokenizer := t.NewTokenizer(data)

	tokenizer.NextToken()

}
