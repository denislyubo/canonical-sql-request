package parser

import (
	"bytes"
	"fmt"
	tknz "github.com/denislyubo/canonical-sql-request/tokenizer"
	"strings"
)

var (
	tokens = map[string]struct{}{
		"select": {},
		"*":      {},
		"from":   {},
		"in":     {},
	}
)

type Parser struct {
	sb *strings.Builder
}

func NewParser() Parser {
	return Parser{sb: &strings.Builder{}}
}

func (p *Parser) Parse(tokenizer tknz.Tokenizer) error {
outer:
	for {
		token, err := tokenizer.NextToken()
		if err == tknz.ErrEnd {
			break
		} else if err != nil {
			return err
		}

		t := bytes.ToLower(token)
		st := string(t)

		if st == "in" {
			if t, _ := tokenizer.NextToken(); string(t) == "(" {
				for {
					token, err := tokenizer.NextToken()
					if err != nil {
						return fmt.Errorf("error parsing closing in brzcket statment: %w", err)
					}

					if string(token) == ")" {
						p.sb.WriteString("in (...)")
						continue outer
					}
				}
			}
		}
		p.sb.Write(t)
		p.sb.WriteRune(' ')
	}

	return nil
}

func (p Parser) String() string {
	s := p.sb.String()

	if s[len(s)-1] == ' ' {
		return s[0 : len(s)-1]
	}
	return s
}

func (p *Parser) Reset() {
	p.sb = &strings.Builder{}
}
