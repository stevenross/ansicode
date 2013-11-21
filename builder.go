package ansicode

// Builder is a string builder that has chainable methods that can be
// used to apply ANSI codes to spans of text. Use the Span function to
// initilise a new Builder.
type Builder struct {
	spans    []*span
	currSpan *span
}

// Span returns a newly created Builder with the given text as the
// first span.
func Span(s string) *Builder {
	b := &Builder{
		make([]*span, 0, 4),
		nil,
	}
	return b.Span(s)
}

// Span is also provided as a chainable method to append a new span to
// the string Builder with the given string. Appended spans do not
// inherit ANSI codes set on the previous span, for this use the Copy
// method.
func (b *Builder) Span(s string) *Builder {
	b.currSpan = &span{
		s,
		"",
		make([]string, 0, 4),
	}
	b.spans = append(b.spans, b.currSpan)
	return b
}

// Code appends the given ANSI code to the current span.
func (b *Builder) Code(code int) *Builder {
	b.currSpan.addCode(code)
	return b
}

// Copy inherits all ANSI codes set on the previous span to the
// current span.
func (b *Builder) Copy() *Builder {
	if l := len(b.spans); l > 1 {
		b.currSpan.colourCode, b.currSpan.codes = b.spans[l-2].colourCode, b.spans[l-2].codes
	}
	return b
}

// String returns the full string from all spans with ANSI codes
// applied.
func (b *Builder) String() (ansiString string) {
	for _, s := range b.spans {
		ansiString += "\x1B["
		last := len(s.codes)-1
		for i, code := range s.codes {
			ansiString += code
			if (i < last) || (i == last && s.colourCode != "") {
				ansiString += ";"
			}
		}
		ansiString += s.colourCode + "m" + s.text + Reset
	}
	return
}

// PlainString returns the full string from all spans without ANSI
// codes.
func (b* Builder) PlainString() (plainString string) {
	for _, s := range b.spans {
		plainString += s.text
	}
	return
}