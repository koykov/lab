package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

type languagesTuple struct {
	Name    string   `json:"name"`
	Native  string   `json:"native"`
	Iso6391 string   `json:"iso639_1"`
	Iso6393 string   `json:"iso639_3"`
	Scripts []string `json:"scripts"`
}

var (
	sidx = map[string]struct{}{
		"Adlam":                  {},
		"Ahom":                   {},
		"Anatolian Hieroglyphs":  {},
		"Arabic":                 {},
		"Armenian":               {},
		"Avestan":                {},
		"Balinese":               {},
		"Bamum":                  {},
		"Bassa Vah":              {},
		"Batak":                  {},
		"Bengali":                {},
		"Bhaiksuki":              {},
		"Bopomofo":               {},
		"Brahmi":                 {},
		"Braille":                {},
		"Buginese":               {},
		"Buhid":                  {},
		"Canadian Aboriginal":    {},
		"Carian":                 {},
		"Caucasian Albanian":     {},
		"Chakma":                 {},
		"Cham":                   {},
		"Cherokee":               {},
		"Chorasmian":             {},
		"Common":                 {},
		"Coptic":                 {},
		"Cuneiform":              {},
		"Cypriot":                {},
		"Cyrillic":               {},
		"Deseret":                {},
		"Devanagari":             {},
		"Dives Akuru":            {},
		"Dogra":                  {},
		"Duployan":               {},
		"Egyptian Hieroglyphs":   {},
		"Elbasan":                {},
		"Elymaic":                {},
		"Ethiopic":               {},
		"Georgian":               {},
		"Glagolitic":             {},
		"Gothic":                 {},
		"Grantha":                {},
		"Greek":                  {},
		"Gujarati":               {},
		"Gunjala Gondi":          {},
		"Gurmukhi":               {},
		"Han":                    {},
		"Hangul":                 {},
		"Hanifi Rohingya":        {},
		"Hanunoo":                {},
		"Hatran":                 {},
		"Hebrew":                 {},
		"Hiragana":               {},
		"Imperial Aramaic":       {},
		"Inherited":              {},
		"Inscriptional Pahlavi":  {},
		"Inscriptional Parthian": {},
		"Javanese":               {},
		"Kaithi":                 {},
		"Kannada":                {},
		"Katakana":               {},
		"Kayah Li":               {},
		"Kharoshthi":             {},
		"Khitan Small Script":    {},
		"Khmer":                  {},
		"Khojki":                 {},
		"Khudawadi":              {},
		"Lao":                    {},
		"Latin":                  {},
		"Lepcha":                 {},
		"Limbu":                  {},
		"Linear A":               {},
		"Linear B":               {},
		"Lisu":                   {},
		"Lycian":                 {},
		"Lydian":                 {},
		"Mahajani":               {},
		"Makasar":                {},
		"Malayalam":              {},
		"Mandaic":                {},
		"Manichaean":             {},
		"Marchen":                {},
		"Masaram Gondi":          {},
		"Medefaidrin":            {},
		"Meetei Mayek":           {},
		"Mende Kikakui":          {},
		"Meroitic Cursive":       {},
		"Meroitic Hieroglyphs":   {},
		"Miao":                   {},
		"Modi":                   {},
		"Mongolian":              {},
		"Mro":                    {},
		"Multani":                {},
		"Myanmar":                {},
		"Nabataean":              {},
		"Nandinagari":            {},
		"New Tai Lue":            {},
		"Newa":                   {},
		"Nko":                    {},
		"Nushu":                  {},
		"Nyiakeng Puachue Hmong": {},
		"Ogham":                  {},
		"Ol Chiki":               {},
		"Old Hungarian":          {},
		"Old Italic":             {},
		"Old North Arabian":      {},
		"Old Permic":             {},
		"Old Persian":            {},
		"Old Sogdian":            {},
		"Old South Arabian":      {},
		"Old Turkic":             {},
		"Oriya":                  {},
		"Osage":                  {},
		"Osmanya":                {},
		"Pahawh Hmong":           {},
		"Palmyrene":              {},
		"Pau Cin Hau":            {},
		"Phags Pa":               {},
		"Phoenician":             {},
		"Psalter Pahlavi":        {},
		"Rejang":                 {},
		"Runic":                  {},
		"Samaritan":              {},
		"Saurashtra":             {},
		"Sharada":                {},
		"Shavian":                {},
		"Siddham":                {},
		"SignWriting":            {},
		"Sinhala":                {},
		"Sogdian":                {},
		"Sora Sompeng":           {},
		"Soyombo":                {},
		"Sundanese":              {},
		"Syloti Nagri":           {},
		"Syriac":                 {},
		"Tagalog":                {},
		"Tagbanwa":               {},
		"Tai Le":                 {},
		"Tai Tham":               {},
		"Tai Viet":               {},
		"Takri":                  {},
		"Tamil":                  {},
		"Tangut":                 {},
		"Telugu":                 {},
		"Thaana":                 {},
		"Thai":                   {},
		"Tibetan":                {},
		"Tifinagh":               {},
		"Tirhuta":                {},
		"Ugaritic":               {},
		"Vai":                    {},
		"Wancho":                 {},
		"Warang Citi":            {},
		"Yezidi":                 {},
		"Yi":                     {},
		"Zanabazar Square":       {},
	}
)

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
	if err = os.WriteFile("target.json", raw, 0644); err != nil {
		log.Fatalln(err)
	}
}
