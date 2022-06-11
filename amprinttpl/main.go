package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	ps   = string(os.PathSeparator)
	now  = time.Now().Format("2006-01-02")
	rtpl []byte
	fdb  = flag.String("db", fmt.Sprintf("local%sdb.csv", ps), "path to local database")
	ftpl = flag.String("tpl", fmt.Sprintf("local%stpl.html", ps), "path to template file")
	fout = flag.String("out", fmt.Sprintf(".%sout", ps), "path to output directory")

	bSeria     = []byte("{SERIA}")
	bNumber    = []byte("{NUMBER}")
	bDateDay   = []byte("{DATE_DAY}")
	bDateMonth = []byte("{DATE_MONTH}")
	bDateYear  = []byte("{DATE_YEAR}")
	bCarModel  = []byte("{CAR_MODEL}")
	bCarNumber = []byte("{CAR_NUMBER}")
	bDriver    = []byte("{DRIVER_NAME}")
)

func init() {
	if !fileExists(*fdb) {
		log.Fatalf("local database not exists %s\n", *fdb)
	}
	if !fileExists(*ftpl) {
		log.Fatalf("template file not exists %s\n", *ftpl)
	}
	var err error
	if rtpl, err = os.ReadFile(*ftpl); err != nil {
		log.Fatalf("failed to read template file: %s", err.Error())
	}
	if err = dirProbe(*fout); err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open(*fdb)
	if err != nil {
		log.Fatalf("coudn't open local database '%s': %s", *fdb, err)
	}
	defer func() { _ = f.Close() }()

	csvr := csv.NewReader(f)
	csvr.Comma = ';'
	rows, err := csvr.ReadAll()
	if err != nil {
		log.Fatalf("couldn't parse database file '%s': %s", *fdb, err)
	}
	if len(rows) <= 1 {
		log.Fatal("local database is empty")
	}
	rows = rows[1:]
	var buf []byte
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		dn := strings.ReplaceAll(row[4], " ", "_")
		out := fmt.Sprintf("%s%s%s__%s.html", *fout, ps, now, dn)
		log.Printf("processing '%s' to '%s", row[4], out)
		buf = append(buf[:0], rtpl...)
		buf = bytes.ReplaceAll(buf, bSeria, []byte(row[0]))
		buf = bytes.ReplaceAll(buf, bNumber, []byte(row[1]))
		buf = bytes.ReplaceAll(buf, bDateDay, []byte(time.Now().Format("2")))
		buf = bytes.ReplaceAll(buf, bDateMonth, []byte(time.Now().Format("1")))
		buf = bytes.ReplaceAll(buf, bDateYear, []byte(time.Now().Format("2006")))
		buf = bytes.ReplaceAll(buf, bCarModel, []byte(row[2]))
		buf = bytes.ReplaceAll(buf, bCarNumber, []byte(row[3]))
		buf = bytes.ReplaceAll(buf, bDriver, []byte(row[4]))
		if err = os.WriteFile(out, buf, 0777); err != nil {
			log.Fatalf("failed to write file '%s: %s", out, err.Error())
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func dirProbe(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
