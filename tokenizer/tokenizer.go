package tokenizer

import "fmt"

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

type Tokenizer struct {
	data []byte
	i    int
}

func NewTokenizer(d []byte) Tokenizer {
	return Tokenizer{data: d}
}

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
