package parser_test

import (
	prsr "github.com/denislyubo/canonical-sql-request/parser"
	tknz "github.com/denislyubo/canonical-sql-request/tokenizer"
	"testing"
)

type testData struct {
	name, given, expected string
}

func TestParser1(t *testing.T) {
	parser := prsr.NewParser()

	tests := []testData{
		{name: "select", given: "SELECT * FROM Account", expected: "select * from account"},
		{name: "select in", given: "SELECT * FROM Account WHERE AccountName in (\"first\", \"second\")",
			expected: "select * from account where accountname in (...)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parser.Parse(tknz.NewTokenizer([]byte(tt.given)))
			defer parser.Reset()
			if err != nil {
				t.Errorf("returned error %v", err)
			}
			if parser.String() != tt.expected {
				t.Errorf("got: %s, expected: %s", parser, tt.expected)
			}
		})
	}
}
