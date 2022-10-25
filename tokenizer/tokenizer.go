// Package tokenizer responsible for tokenizing input byte slice.
package tokenizer

import (
	"fmt"
	"sync"
)

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
	pool *sync.Pool
}

// NewTokenizer constructs Tokenizer object
func NewTokenizer(d []byte) *Tokenizer {
	var p = sync.Pool{
		New: func() any {
			return make([]byte, 1024)
		},
	}
	return &Tokenizer{data: d, pool: &p}
}

// NextToken goes by input slice and returns tokens.
// Delimiters are stored in tokens map. Default delimiter is gap.
// Should be run in a loop.
// ErrEnd returned means end of processing.
func (t *Tokenizer) NextToken() ([]byte, error) {
	buffer := t.pool.Get().([]byte)[:0]
	defer t.pool.Put(buffer)

	started, exit := false, false
	for {
		if t.i >= len(t.data) {
			if started {
				return append([]byte(nil), buffer...), nil
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
			if _, ok := tokens[string(t.data[t.i])]; ok && len(buffer) == 0 {
				buffer = append(buffer, t.data[t.i])
				t.i++
				exit = true
			} else if ok && len(buffer) > 0 || t.data[t.i] == ' ' {
				exit = true
			} else {
				buffer = append(buffer, t.data[t.i])
				t.i++
			}
		}

		if exit {
			break
		}
	}

	return append([]byte(nil), buffer...), nil
}
