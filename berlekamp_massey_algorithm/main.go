package main

import "fmt"

func BerlekampMassey(s []byte) []byte {
	n := len(s)
	if n == 0 {
		return nil
	}

	c := make([]byte, n)
	b := make([]byte, n)
	c[0], b[0] = 1, 1

	var L, m int
	var d, mask byte

	for nIter := 0; nIter < n; nIter++ {

		d = s[nIter]
		for i := 1; i <= L; i++ {
			d ^= c[i] & s[nIter-i]
		}

		if d == 1 {

			temp := make([]byte, n)
			copy(temp, c)

			mask = 1
			shift := nIter - m
			for i := 0; i < n; i++ {
				if i+shift < n {
					c[i+shift] ^= b[i] & mask
				}
			}

			if L <= nIter/2 {
				L = nIter + 1 - L
				m = nIter
				copy(b, temp)
			}
		}
	}

	res := make([]byte, L+1)
	copy(res, c[:L+1])
	return res
}

func main() {
	sequence := []byte{1, 0, 1, 1, 0, 1, 0, 0, 1, 1}
	lfsr := BerlekampMassey(sequence)
	fmt.Printf("LFSR: %v\n", lfsr)
}
