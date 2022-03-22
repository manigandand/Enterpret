package ingester

// Delims contains the info of left and right delimiters
type Delims struct {
	Left  string
	Right string
}

// Enclose method is used to enclose a given string
// into the delimiters
func (d Delims) Enclose(a string) string {
	return d.Left + a + d.Right
}

// default gohtml delimiters
// DefaultDelimiters is used as a configuration variable to store the default
// delimiters pair (which is "{{", "}}")
var DefaultDelimiters = Delims{
	Left:  "{{",
	Right: "}}",
}
