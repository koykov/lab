package xtou

var tab [256]byte

func init() {
	for c := '0'; c <= '9'; c++ {
		tab[c] = uint8(c - '0')
	}
	for c := 'a'; c <= 'f'; c++ {
		tab[c] = uint8(c - 'a' + 10)
	}
	for c := 'A'; c <= 'F'; c++ {
		tab[c] = uint8(c - 'A' + 10)
	}
}

func xtouTable(s []byte) uint64 {
	_, _ = tab[255], s[3]
	return uint64(tab[s[0]])<<12 | uint64(tab[s[1]])<<8 | uint64(tab[s[2]])<<4 | uint64(tab[s[3]])
}
