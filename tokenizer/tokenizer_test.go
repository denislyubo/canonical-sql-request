package tokenizer_test

import (
	tknz "github.com/denislyubo/canonical-sql-request/tokenizer"
	"testing"
)

type TestData struct {
	Name     string
	Expected string
}

func TestTokenizer1(t *testing.T) {
	input := []byte("SELECT * FROM Account")
	tokenizer := tknz.NewTokenizer(input)

	tests := []TestData{
		{Name: "select", Expected: "SELECT"},
		{Name: "asterisk", Expected: "*"},
		{Name: "from", Expected: "FROM"},
		{Name: "account", Expected: "Account"},
		{Name: "end of string", Expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tokenizer.NextToken()
			if err != nil {
				if tt.Name == "end of string" && err != tknz.ErrEnd {
					t.Errorf("returned error %v", err)
				}
			}
			if got != nil && string(got) != tt.Expected {
				t.Errorf("got: %s, expected: %s", got, tt.Expected)
			}
		})
	}
}

func TestTokenizer2(t *testing.T) {
	input := []byte("SELECT * FROM Account WHERE AccountName in (\"first\", \"second\")")
	tokenizer := tknz.NewTokenizer(input)

	tests := []TestData{
		{Name: "select", Expected: "SELECT"},
		{Name: "asterisk", Expected: "*"},
		{Name: "from", Expected: "FROM"},
		{Name: "account", Expected: "Account"},
		{Name: "where", Expected: "WHERE"},
		{Name: "AccountName", Expected: "AccountName"},
		{Name: "in", Expected: "in"},
		{Name: "(", Expected: "("},
		{Name: "\"", Expected: "\""},
		{Name: "first", Expected: "first"},
		{Name: "\"", Expected: "\""},
		{Name: "comma", Expected: ","},
		{Name: "\"", Expected: "\""},
		{Name: "second", Expected: "second"},
		{Name: "\"", Expected: "\""},
		{Name: ")", Expected: ")"},
		{Name: "end of string", Expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tokenizer.NextToken()
			if err != nil {
				if tt.Name == "end of string" && err != tknz.ErrEnd {
					t.Errorf("returned error %v", err)
				}
			}
			if got != nil && string(got) != tt.Expected {
				t.Errorf("got: %s, expected: %s", got, tt.Expected)
			}
		})
	}
}
