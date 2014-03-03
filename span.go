// Package ansicode provides a string builder with a chainable API to
// apply ANSI code styles to spans of text.
package ansicode

import "strconv"

// span represents a string of text that can be stylised with ANSI codes.
type span struct {
	text       string
	colourCode string
	codes      []string
}

// code adds the ANSI code given to the span if it hasn't already
// been added.
func (s *span) code(code int) {
	s.codeString(strconv.Itoa(code))
}

// codeString is the same as code but accepts a string.
func (s *span) codeString(code string) {
	if len(code) == 2 && code[:1] == "3" && code[1:] != "8" && code[1:] != "9" {
		s.colourCode = code
		return
	}
	for _, existingCode := range s.codes {
		if code == existingCode {
			return
		}
	}
	s.codes = append(s.codes, code)
}
