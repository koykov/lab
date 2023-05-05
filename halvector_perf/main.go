package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/koykov/bytealg"
	"github.com/koykov/halvector"
)

type Tuple struct {
	HAL           string
	LangNoLimit   string
	LangWithLimit string
	Speed         int64
	Speed50       int64
	Speed10       int64
	Speed5        int64
	Speed3        int64
}

func main() {
	fr, err := os.Open("dataset.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scr := bufio.NewScanner(fr)
	scr.Split(bufio.ScanLines)

	var (
		tuples []Tuple
		c      int
		n0, n1 int64
	)
	vec := halvector.NewVector()
	for scr.Scan() {
		vec.Reset()
		p := scr.Bytes()
		n0 = time.Now().UnixNano()
		err = vec.Parse(p)
		n1 = time.Now().UnixNano()
		if err != nil {
			log.Println("error:", err)
			continue
		}
		t := Tuple{
			HAL:         string(p),
			LangNoLimit: bytealg.Copy(vec.Root().DotString("0.code")),
			Speed:       n1 - n0,
		}

		vec.Reset()
		n0 = time.Now().UnixNano()
		_ = vec.SetLimit(50).Parse(p)
		n1 = time.Now().UnixNano()
		t.LangWithLimit = bytealg.Copy[string](vec.Root().DotString("0.code"))
		t.Speed50 = n1 - n0

		vec.Reset()
		n0 = time.Now().UnixNano()
		_ = vec.SetLimit(10).Parse(p)
		n1 = time.Now().UnixNano()
		t.LangWithLimit = bytealg.Copy[string](vec.Root().DotString("0.code"))
		t.Speed10 = n1 - n0

		vec.Reset()
		n0 = time.Now().UnixNano()
		_ = vec.SetLimit(5).Parse(p)
		n1 = time.Now().UnixNano()
		t.LangWithLimit = bytealg.Copy[string](vec.Root().DotString("0.code"))
		t.Speed5 = n1 - n0

		vec.Reset()
		n0 = time.Now().UnixNano()
		_ = vec.SetLimit(3).Parse(p)
		n1 = time.Now().UnixNano()
		t.LangWithLimit = bytealg.Copy[string](vec.Root().DotString("0.code"))
		t.Speed3 = n1 - n0

		tuples = append(tuples, t)

		c++
	}
	log.Println("parsed:", c)

	fw, err := os.Create("records.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(fw)
	defer w.Flush()
	row := []string{"hal", "lang_no_limit", "lang_with_limit", "speed", "speed_limit50", "speed_limit10", "speed_limit5", "speed_limit3"}
	_ = w.Write(row)
	for _, t := range tuples {
		row = []string{t.HAL, t.LangNoLimit, t.LangWithLimit, strconv.Itoa(int(t.Speed)), strconv.Itoa(int(t.Speed50)), strconv.Itoa(int(t.Speed10)), strconv.Itoa(int(t.Speed5)), strconv.Itoa(int(t.Speed3))}
		_ = w.Write(row)
	}

	_ = fr.Close()
	_ = fw.Close()
}
