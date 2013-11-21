
// Package ansicode provides a string builder with a chainable API to
// apply ANSI code styles to spans of text.
package ansicode

import "strconv"

// span represents a string of text that can be sylised with ANSI codes.
type span struct {
	text, colourCode string
	codes            []string
}

// addCode adds the ANSI code given to the span if it hasn't already
// been added.
func (s *span) addCode(code int) {
	s.addCodeString(strconv.Itoa(code))
}

// addCodeString is the same as addCode but accepts a string.
func (s *span) addCodeString(code string) {
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