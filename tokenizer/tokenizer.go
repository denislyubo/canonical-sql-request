// Package tokenizer responsible for tokenizing input byte slice.
package tokenizer

import "fmt"

// ErrEnd means end of input byte slice.
var (
	ErrEnd = fmt.Errorf("end of input")
)

var (
	tokens = map[string]struct{}{
		"(":  {},
		")":  {},
		"\"": {},
		",":  {},
	}
)

// Tokenizer is main object for storing data.
type Tokenizer struct {
	data []byte
	i    int
}

// NewTokenizer constructs Tokenizer object
func NewTokenizer(d []byte) Tokenizer {
	return Tokenizer{data: d}
}

// NextToken goes by input slice and returns tokens.
// Delimiters are stored in tokens map. Default delimiter is gap.
// Should be run in a loop.
// ErrEnd returned means end of processing.
func (t *Tokenizer) NextToken() ([]byte, error) {
	var b []byte
	started, exit := false, false
	for {
		if t.i >= len(t.data) {
			if started {
				return b, nil
			}
			return nil, ErrEnd
		}

		if !started {
			if t.data[t.i] == ' ' {
				t.i++
				continue
			} else {
				started = true
			}
		}

		if started {
			if _, ok := tokens[string(t.data[t.i])]; ok && len(b) == 0 {
				b = append(b, t.data[t.i])
				t.i++
				exit = true
			} else if ok && len(b) > 0 || t.data[t.i] == ' ' {
				exit = true
			} else {
				b = append(b, t.data[t.i])
				t.i++
			}
		}

		if exit {
			break
		}
	}

	return b, nil
}
