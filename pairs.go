package fstr

/*
Pairs allows input of a format string incorporating a string interpolation
syntax where the interpolation signaller is a key enclosed in braces
(example: "Hello, {name}."). Values to interpolate to the input string may then
be provided to variadic arg `vs` in groups of two where for each two inputs the
first represents a key and the second represents the correlative value

Example:

	v := Pairs(
		"Hello, {name1} and {name2}!",
		"name1", "John",
		"name2", "Bill",
	)
	fmt.Println(v) // Hello, John and Bill!
*/
func Pairs(v string, vs ...string) string {
	// If the map has no keys, just return the input
	if len(vs) == 0 {
		return v
	}

	// Initialize the buffer to write the resulting string to
	out := buffer{}

	// Track input string position
	// We always look ahead so we start from -1
	pos := -1

	// Loop until we hit the end of the line
outer:
	for pos < len(v)-1 {
		// Lookahead one byte
		r := v[pos+1]

		switch r {
		// If we find a backslash and the next character is an open brace - skip
		// string interpolation for this token (escaped)
		case '\\':
			if pos+2 < len(v) {
				if v[pos+2] == '{' {
					pos++                       // current : '\'
					pos++                       // current : '{'
					out = out.writeByte(v[pos]) // write   : '{'
					continue
				}

				// Write out the escape character normally
				pos++                     // current : '\'
				out = out.writeByte('\\') // write : '\'
			}

		// If we found an opening brace and it's not escaped, consume it as an
		// interpolation token
		case '{':
			pos++ // current: {

			// Find the closing brace position
			start := pos + 1
			for v[pos+1] != '}' {
				pos++
			}

			// Take the start + end+1 positions as the interpolation token slice range
			token := string(v[start : pos+1])

			// Look for a hit in our interpolation pairs. The pairs are laid out in
			// twos as `"token", "value"`, so we step by 2
			for i := 0; i < len(vs); i += 2 {
				// Make sure we're not indexing outside the bounds of the input array
				if i+1 >= len(vs) {
					break
				}

				if vs[i] == token {
					// If we find one, write its corresponding (pos+1) value and continue
					// on to avoid the raw token getting written back out
					out = out.writeString(vs[i+1])
					pos++ // current: }
					continue outer
				}
			}

			// If we find no matching interpolation token, simply write the token out
			for i := start; i < pos+1; i++ {
				out = out.writeByte(v[i])
			}
			pos++ // current: }

		default:
			// Just a normal byte, write it to the buffer
			pos++
			out = out.writeByte(v[pos])
		}
	}

	return string(out)
}
