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
}

func main() {
	var rev bool
	for i := 0; i < len(src); i++ {
		fn := src[i]
		base := path.Base(fn)
		if ext := path.Ext(base); len(ext) > 0 {
			base = base[:len(base)-len(ext)]
		}
		sp := strings.Index(base, "_")
		if sp == -1 {
			continue
		}
		l, r := base[:sp], base[sp+1:]
		rl.reset()
		switch {
		case l == "English":
			rl.lng = r
			rev = true
		case r == "English":
			rl.lng = l
			rev = false
		default:
			continue
		}
		log.Printf("processing '%s' ...\n", fn)
		log.Printf("processing '%s' ...\n", rfn)
		// todo process fn and rfn
	}
}
