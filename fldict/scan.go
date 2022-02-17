package main

import (
	"bufio"
	"bytes"
	"os"

	"github.com/koykov/bytealg"
)

func scan(dst, dstEn *repo, filename string, reverse bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	scr := bufio.NewScanner(file)
	for i := 0; scr.Scan(); i++ {
		if i < 3 {
			continue
		}
		line := scr.Bytes()
		si := bytes.IndexByte(line, '|')
		if si == -1 || si == len(line)-1 {
			continue
		}
		l, r := bytealg.Trim(line[:si], bAt), bytealg.Trim(line[si+1:], bAt)
		l = clean(l)
		r = clean(r)
		if reverse {
			dst.add(r)
			dstEn.add(l)
		} else {
			dst.add(l)
			dstEn.add(r)
		}
	}

	return scr.Err()
}

func clean(p []byte) []byte {
	var bd int
	for i := 0; i < len(p); i++ {
		if p[i] == '(' {
			bd++
		}
		if p[i] == ')' {
			bd--
			if bd < 0 {
				bd = 0
			}
		}
		if bd > 0 || p[i] == ')' || p[i] == '@' {
			p[i] = ' '
		}
		if p[i] == ',' || p[i] == '/' || p[i] == ';' || p[i] == ':' {
			p[i] = '|'
		}
	}

	p = bytealg.Trim(p, bSpace)
	p = bytealg.Trim(p, bCol)
	for i := 0; i < len(repl); i++ {
		r := repl[i]
		if pos := bytes.Index(p, r); pos != -1 {
			for j := 0; j < len(r); j++ {
				p[pos+j] = ' '
			}
		}
	}
	return p
}
