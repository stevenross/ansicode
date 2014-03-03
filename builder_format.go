package ansicode

// Bold sets the ANSI code for bold text to the current span.
func (b *Builder) Bold() *Builder {
	b.currSpan.codeString("1")
	return b
}

// Underline sets the ANSI code for underlined text to the current span.
func (b *Builder) Underline() *Builder {
	b.currSpan.codeString("4")
	return b
}

// Inverse sets the ANSI code for inversed text to the current span.
func (b *Builder) Inverse() *Builder {
	b.currSpan.codeString("7")
	return b
}