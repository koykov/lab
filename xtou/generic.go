package xtou

func xtouGeneric(s []byte) (n uint64) {
	for i := 0; i < 4; i++ {
		var v byte
		c := s[i]
		switch {
		case '0' <= c && c <= '9':
			v = c - '0'
		case 'a' <= c && c <= 'f':
			v = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			v = c - 'A' + 10
		default:
			return 0
		}
		n = n<<4 | uint64(v)
	}
	return
}
