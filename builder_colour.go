package ansicode

// Black sets the ANSI code for black coloured text to the current span.
func (b *Builder) Black() *Builder {
	b.currSpan.colourCode = "30"
	return b
}

// Red sets the ANSI code for red coloured text to the current span.
func (b *Builder) Red() *Builder {
	b.currSpan.colourCode = "31"
	return b
}

// Green sets the ANSI code for green coloured text to the current span.
func (b *Builder) Green() *Builder {
	b.currSpan.colourCode = "32"
	return b
}

// Yellow sets the ANSI code for yellow coloured text to the current span.
func (b *Builder) Yellow() *Builder {
	b.currSpan.colourCode = "33"
	return b
}

// Blue sets the ANSI code for blue coloured text to the current span.
func (b *Builder) Blue() *Builder {
	b.currSpan.colourCode = "34"
	return b
}

// Magenta sets the ANSI code for magenta coloured text to the current span.
func (b *Builder) Magenta() *Builder {
	b.currSpan.colourCode = "35"
	return b
}

// Cyan sets the ANSI code for cyan coloured text to the current span.
func (b *Builder) Cyan() *Builder {
	b.currSpan.colourCode = "36"
	return b
}

// White sets the ANSI code for white coloured text to the current span.
func (b *Builder) White() *Builder {
	b.currSpan.colourCode = "37"
	return b
}