package fstr

type buffer []byte

func (buf buffer) writeByte(b byte) buffer {
	buf = append(buf, b)

	return buf
}

func (buf buffer) writeRune(r rune) buffer {
	buf = append(buf, byte(r))

	return buf
}

func (buf buffer) writeString(s string) buffer {
	for _, r := range s {
		buf = append(buf, byte(r))
	}

	return buf
}
