package ansicode

// These ANSI code constants can be used to manually constuct
// ANSI coded strings.
const (
	// Colour
	Black   = "\x1B[30m"
	Red     = "\x1B[31m"
	Green   = "\x1B[32m"
	Yellow  = "\x1B[33m"
	Blue    = "\x1B[34m"
	Magenta = "\x1B[35m"
	Cyan    = "\x1B[36m"
	White   = "\x1B[37m"

	// Format
	Bold      = "\x1B[1m"
	Underline = "\x1B[4m"
	Inverse   = "\x1B[7m"

	// Misc
	Reset = "\x1B[0m"
)