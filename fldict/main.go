package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	fsrc, fdst string
	src        []string
)

func init() {
	flag.StringVar(&fsrc, "source", "", "path to source directory")
	flag.StringVar(&fdst, "destination", "", "path to destination directory")
	flag.Parse()

	if len(fsrc) == 0 {
		log.Fatalln("param --source is required")
	}
	stat, err := os.Stat(fsrc)
	if os.IsNotExist(err) {
		log.Fatalf("source '%s' doesn't exists\n", fsrc)
	}
	if !stat.IsDir() {
		log.Fatalf("source '%s' must be directory\n", fsrc)
	}

	if src, err = filepath.Glob(fsrc + "/*.txt"); err != nil || len(src) == 0 {
		log.Fatalf("cannot read files list in '%s'\n", fsrc)
	}

	if stat, err = os.Stat(fdst); os.IsNotExist(err) {
		if err = os.MkdirAll(fdst, 0755); err != nil {
			log.Fatalf("cannot create destination '%s'\n", fdst)
		}
	}
	if stat == nil {
		stat, _ = os.Stat(fdst)
	}
	if !stat.IsDir() {
		log.Fatalf("destination '%s' must be directory\n", fdst)
	}
}

func main() {
	pfx := fsrc + "/"
	for i := 0; i < len(src); i++ {
		fn := src[i]
		base := path.Base(fn)
		rbase := "Reversed" + base
		rfn := pfx + rbase
		if strings.HasPrefix(base, "Reversed") {
			continue
		}
		log.Printf("processing '%s' ...\n", fn)
		log.Printf("processing '%s' ...\n", rfn)
		// todo process fn and rfn
	}
}
