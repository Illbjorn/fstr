package fstr

import "bytes"

/*
Map allows input of a format string incorporating a string interpolation
syntax where the interpolation signaller is a key enclosed in curly braces
(example: "Hello, {name}."). Values to interpolate to the input string may then
be provided via map arg `vs`.

Example:

	m := map[string]string{
		"name1": "John",
		"name2": "Bill",
	}

	v := Map("Hello, {name1} and {name2}!", m)

	fmt.Println(v) // Hello, John and Bill!
*/
func Map(v string, vs map[string]string) string {
	// If the map has no keys, just return the input.
	if len(vs) == 0 {
		return v
	}

	// Convert input string to runes.
	runes := []rune(v)

	// Initialize the buffer to write the resulting string to.
	out := bytes.NewBuffer(nil)

	// Track input string position.
	// We always look ahead so we start from -1.
	pos := -1

	// Loop until we hit the end of the line.
outer:
	for pos < len(runes)-1 {
		// Lookahead one rune.
		r := runes[pos+1]

		// If we find a backslash and the next character is an open brace - skip
		// string interpolation for this token (escaped).
		if r == '\\' {
			if pos+2 < len(runes) {
				if runes[pos+2] == '{' {
					pos++                     // current : '\'
					pos++                     // current : '{'
					out.WriteRune(runes[pos]) // write   : '{'
					continue
				}
			}
		}

		// If we found an opening brace and it's not escaped, consume it as an
		// interpolation token.
		if r == '{' {
			pos++ // current: {

			// Find the closing brace position.
			start := pos + 1
			for runes[pos+1] != '}' {
				pos++
			}

			// Take the start + end+1 positions as the interpolation token slice range.
			token := string(runes[start : pos+1])

			pos++ // }

			// Look for a hit in our interpolation map.
			if v, ok := vs[token]; ok {
				// If we find one, write its value and continue on to avoid the raw
				// token getting written back out.
				out.WriteString(v)
				continue outer
			}

			// If we find no matching interpolation token, simply write the token out.
			out.WriteString(token)
		} else {
			// Just a normal rune, write it to the buffer.
			pos++
			out.WriteRune(runes[pos])
		}
	}

	return out.String()
}
