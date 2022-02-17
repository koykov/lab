package main

var (
	bAt    = []byte{'@'}
	bSpace = []byte{' '}
	bCol   = []byte{':'}
	bNl    = []byte{'\n'}
	bSep   = []byte{'|'}
	bTrim  = []byte(" \":!#%&'~?.*+-<=>¿")

	repl = [][]byte{
		[]byte("prep:"),
		[]byte("noun:"),
		[]byte("part:"),
		[]byte("phrase:"),
		[]byte("phrs:"),
		[]byte("phrz:"),
		[]byte("art:"),
		[]byte("pron:"),
		[]byte("art:"),
		[]byte("rep:"),
		[]byte("inf:"),
		[]byte("inf :"),
		[]byte("fem :"),
		[]byte("plu :"),
		[]byte("inan :"),
		[]byte("auxv:"),
		[]byte("sin :"),
		[]byte("plu :"),
		[]byte("inf :"),
		[]byte("masc :"),
		[]byte("neu :"),
		[]byte("pro:"),
		[]byte("name:"),
		[]byte("informal :"),
		[]byte("infor :"),
		[]byte("past :"),
		[]byte("pres. 2ps :"),
		[]byte("initial form :"),
		[]byte("past passive participle :"),
	}
)
