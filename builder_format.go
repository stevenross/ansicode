package ansicode

// Bold sets the ANSI code for bold text to the current span.
func (b *Builder) Bold() *Builder {
	b.currSpan.addCodeString("1")
	return b
}

// Underline sets the ANSI code for underlined text to the current span.
func (b *Builder) Underline() *Builder {
	b.currSpan.addCodeString("4")
	return b
}

// Inverse sets the ANSI code for inversed text to the current span.
func (b *Builder) Inverse() *Builder {
	b.currSpan.addCodeString("7")
	return b
}