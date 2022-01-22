package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"sort"
)

type languagesTuple struct {
	Name    string   `json:"name"`
	Native  string   `json:"native"`
	Iso6391 string   `json:"iso639_1"`
	Iso6393 string   `json:"iso639_3"`
	Weight  uint     `json:"weight"`
	Scripts []string `json:"scripts"`
}

type scriptsTuple struct {
	Name      string   `json:"name"`
	Weight    uint     `json:"weight"`
	Languages []string `json:"languages"`
}

func main() {
	var (
		tupl []languagesTuple
		idx  = make(map[string]*languagesTuple)
		raw  []byte
		err  error
	)
	if raw, err = os.ReadFile("origin.json"); err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(raw, &tupl); err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(tupl); i++ {
		t := &tupl[i]
		idx[t.Name] = t
	}

	var (
		f   *os.File
		rdr *csv.Reader
		rec [][]string
	)
	if f, err = os.Open("script.csv"); err != nil {
		log.Fatalln(err)
	}
	rdr = csv.NewReader(f)
	if rec, err = rdr.ReadAll(); err != nil {
		log.Fatalln(err)
	}
	var script string
	for i := 0; i < len(rec); i++ {
		r := rec[i]
		script1, lang := r[0], r[3]
		if len(script1) > 0 {
			script = script1
		}
		if len(lang) == 0 {
			continue
		}
		if t, ok := idx[lang]; ok {
			t.Scripts = append(t.Scripts, script)
			continue
		}
	}
	if raw, err = json.MarshalIndent(tupl, "", "\t"); err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile("languages.json", raw, 0644); err != nil {
		log.Fatalln(err)
	}

	var (
		scrs []scriptsTuple
		scri = make(map[string]int)
	)
	for i := 0; i < len(tupl); i++ {
		for j := 0; j < len(tupl[i].Scripts); j++ {
			scr := tupl[i].Scripts[j]
			if idx, ok := scri[scr]; ok {
				scrs[idx].Weight++
				ok = false
				for k := 0; k < len(scrs[idx].Languages); k++ {
					ok = ok || scrs[idx].Languages[k] == tupl[i].Name
				}
				if !ok {
					scrs[idx].Languages = append(scrs[idx].Languages, tupl[i].Name)
				}
			} else {
				scrs = append(scrs, scriptsTuple{
					Name:      scr,
					Weight:    1,
					Languages: []string{tupl[i].Name},
				})
				scri[scr] = len(scrs) - 1
			}
		}
	}

	sort.Slice(scrs, func(i, j int) bool {
		return scrs[i].Weight > scrs[j].Weight
	})

	if raw, err = json.MarshalIndent(scrs, "", "\t"); err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile("scripts.json", raw, 0644); err != nil {
		log.Fatalln(err)
	}

	return
}
