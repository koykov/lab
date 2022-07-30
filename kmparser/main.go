package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type config struct {
	Source  string            `json:"source"`
	Mapping map[string]string `json:"mapping"`
}

type Tuple struct {
	UA     string `json:"user_agent"`
	Client struct {
		Type          string `json:"type,omitempty"`
		Name          string `json:"name,omitempty"`
		Version       string `json:"version,omitempty"`
		Engine        string `json:"engine,omitempty"`
		EngineVersion string `json:"engine_version,omitempty"`
		Family        string `json:"family,omitempty"`
	} `json:"client,omitempty"`
	Device struct {
		Type      string `json:"type,omitempty"`
		Brand     string `json:"brand,omitempty"`
		Model     string `json:"model,omitempty"`
		OS        string `json:"os,omitempty"`
		OSVersion string `json:"os_version,omitempty"`
	} `json:"device,omitempty"`
}

var conf config

func init() {
	contents, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(contents, &conf); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	resp, err := http.Get(conf.Source)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}

	q, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var buf bytes.Buffer

	_ = buf.WriteByte('[')
	q.Find("tr.bg-warning").Each(func(i int, tr *goquery.Selection) {
		var t Tuple
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			var (
				key,
				value string
				ok bool
			)
			switch j {
			case 2:
				t.UA = td.Text()
			case 3:
				td.Find("div").Each(func(k int, div *goquery.Selection) {
					if div.Length() == 0 {
						return
					}
					div.Find("div").Each(func(l int, div1 *goquery.Selection) {
						switch l {
						case 0:
							ko := div1.Find("strong").Text()
							if key, ok = conf.Mapping[ko]; !ok {
								panic(ko)
							}
						case 1:
							value = div1.Find("span").Text()
						}
					})
					switch key {
					case "device_type":
						t.Device.Type = value
					case "model":
						t.Device.Model = value
					case "vendor":
						t.Device.Brand = value
					case "name":
						t.Client.Name = value
					case "version":
						t.Client.Version = value
					case "os_version":
						t.Device.OSVersion = value
					case "type":
						t.Client.Type = value
					case "os":
						t.Device.OS = value
					}
				})
			}
		})
		b, _ := json.Marshal(&t)
		_, _ = buf.Write(b)
		_ = buf.WriteByte(',')
	})
	bb := buf.Bytes()
	bb[len(bb)-1] = ']'

	if err = os.WriteFile("out/km.json", bb, 0644); err != nil {
		log.Fatalln(err)
	}
}
