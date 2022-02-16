package main

import (
	"bufio"
	"bytes"
	"os"

	"github.com/koykov/bytealg"
)

var (
	bAt    = []byte{'@'}
	bSpace = []byte{' '}
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
			dst.add(r)
		}
	}

	return scr.Err()
}

func clean(p []byte) []byte {
	var pos int
loop:
	bp := bytes.IndexByte(p, '(')
	bp1 := bytes.IndexByte(p, ')')
	if bp != -1 && bp1 != -1 && bp < bp1 {
		if bp1 == len(p)-1 {
			p = p[:bp]
			goto loop
		}
		copy(p, p[bp1:])
		p = p[:bp+(bp1-bp)]
		goto loop
	}
	if pos = bytes.IndexByte(p, ','); pos == -1 {
		if pos = bytes.IndexByte(p, '/'); pos == -1 {
			pos = bytes.IndexByte(p, ';')
		}
	}
	if pos != -1 {
		p[pos] = '|'
		if pos < len(p)-1 && p[pos+1] == ' ' {
			copy(p[pos+1:], p[pos+2:])
			p = p[:len(p)-1]
		}
		goto loop
	}
	return bytealg.Trim(p, bSpace)
}
