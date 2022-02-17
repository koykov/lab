package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/koykov/hash/fnv"
)

var (
	fsrc, fdst string
	src        []string
	rl, rr     *repo
)

func init() {
	flag.StringVar(&fsrc, "dataset", "", "path to dataset directory")
	flag.StringVar(&fdst, "destination", "", "path to destination directory")
	flag.Parse()

	if len(fsrc) == 0 {
		log.Fatalln("param --dataset is required")
	}
	stat, err := os.Stat(fsrc)
	if os.IsNotExist(err) {
		log.Fatalf("dataset '%s' doesn't exists\n", fsrc)
	}
	if !stat.IsDir() {
		log.Fatalf("dataset '%s' must be directory\n", fsrc)
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

	rl, rr = newRepo(fnv.BHasher{}), newRepo(fnv.BHasher{})
	rr.lng = "English"
}

func main() {
	for i := 0; i < len(src); i++ {
		fn := src[i]
		dir := path.Dir(fn)
		base := path.Base(fn)
		if ext := path.Ext(base); len(ext) > 0 {
			base = base[:len(base)-len(ext)]
		}
		sp := strings.Index(base, "_")
		if sp == -1 {
			continue
		}
		l, r := base[:sp], base[sp+1:]
		rfn := dir + "/" + r + "_" + l + ".txt"
		rl.reset()
		switch {
		case r == "English":
			rl.lng = l
		case l == "English":
			fallthrough
		default:
			continue
		}
		log.Printf("processing '%s' ...\n", fn)
		if err := scan(rl, rr, fn, false); err != nil {
			log.Printf("error: %s\n", err.Error())
			continue
		}
		_, err := os.Stat(rfn)
		if os.IsNotExist(err) {
			continue
		}
		log.Printf("processing '%s' ...\n", rfn)
		if err := scan(rl, rr, rfn, true); err != nil {
			log.Printf("error: %s\n", err.Error())
		}

		if err := rl.flush(fdst + "/" + l + ".txt"); err != nil {
			log.Printf("error: %s\n", err.Error())
			continue
		}
	}

	if err := rr.flush(fdst + "/Engligh.txt"); err != nil {
		log.Printf("error: %s\n", err.Error())
	}

	log.Println("done")
}
