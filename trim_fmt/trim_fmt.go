package trim_fmt

type X struct {
	buf []byte
}

func (x *X) trimFmtGotoV0(off int) (int, bool) {
	n := len(x.buf)
	if n > off && x.buf[off] > ' ' {
		return off, false
	}
	_ = x.buf[n-1]
	var c byte
loop:
	if off >= n {
		return off, true
	}
	c = x.buf[off]
	if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
		return off, false
	}
	off++
	goto loop
}

func (x *X) trimFmtGotoV1(off int) (int, bool) {
	n := len(x.buf)
	if n > off && x.buf[off] > ' ' {
		return off, false
	}
	_ = x.buf[n-1]
	var c byte
loop:
	if off >= n {
		return off, true
	}
	c = x.buf[off]
	if c == ' ' {
		off++
		goto loop
	}
	if c == '\t' {
		off++
		goto loop
	}
	if c == '\n' {
		off++
		goto loop
	}
	if c == '\r' {
		off++
		goto loop
	}
	return off, false
}

func (x *X) trimFmtForV0(off int) (int, bool) {
	n := len(x.buf)
	if n > off && x.buf[off] > ' ' {
		return off, false
	}
	_ = x.buf[n-1]
	for ; off < n; off++ {
		c := x.buf[off]
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return off, false
		}
	}
	return off, true
}

func (x *X) trimFmtForV1(off int) (int, bool) {
	n := len(x.buf)
	if n > off && x.buf[off] > ' ' {
		return off, false
	}
	_ = x.buf[n-1]
	for ; off < n; off++ {
		c := x.buf[off]
		if c == ' ' {
			continue
		}
		if c == '\t' {
			continue
		}
		if c == '\n' {
			continue
		}
		if c == '\r' {
			continue
		}
		return off, false
	}
	return off - 1, true
}

func (x *X) trimFmtFJV0(off int) int {
	if len(x.buf) == 0 || x.buf[off] > 0x20 {
		return off
	}
	return x.trimFmtFJSlowV0(off)
}

func (x *X) trimFmtFJSlowV0(off int) int {
	s := x.buf
	if len(s) == 0 || s[0] != 0x20 && s[0] != 0x0A && s[0] != 0x09 && s[0] != 0x0D {
		return off
	}
	for i := off + 1; i < len(s); i++ {
		if s[i] != 0x20 && s[i] != 0x0A && s[i] != 0x09 && s[i] != 0x0D {
			return i
		}
	}
	return len(s) - 1
}

func (x *X) trimFmtFJV1(off int) int {
	s := x.buf[off:]
	if len(s) == 0 || s[off] > 0x20 {
		return off
	}
	return x.trimFmtFJSlowV1(off)
}

func (x *X) trimFmtFJSlowV1(off int) int {
	s := x.buf[off:]
	if len(s) == 0 || s[0] != 0x20 && s[0] != 0x0A && s[0] != 0x09 && s[0] != 0x0D {
		return off
	}
	for i := 1; i < len(s); i++ {
		if s[i] != 0x20 && s[i] != 0x0A && s[i] != 0x09 && s[i] != 0x0D {
			return off + i
		}
	}
	return len(s) - 1
}
