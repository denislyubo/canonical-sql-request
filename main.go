package main

import (
	"bufio"
	"fmt"
	p "github.com/denislyubo/canonical-sql-request/parser"
	"github.com/denislyubo/canonical-sql-request/tokenizer"
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
	parser := p.NewParser()
	parser.Parse(tokenizer.NewTokenizer(data))
	fmt.Println(parser)
	parser.Reset()
}
