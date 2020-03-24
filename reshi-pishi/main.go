package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	exclude = map[string]bool{
		"/site/about": true,
	}
	dlReg = map[string]bool{}

	dlDir, dlAnsDir, chapter string
	idx, sidx                int
)

func main() {
	usr, _ := user.Current()

	resp, err := http.Get("https://reshi-pishi.ru/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#sidebar-menu a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		if _, ok := exclude[href]; ok {
			return
		}
		title := strings.Trim(selection.Text(), " \n")
		if href == "#" {
			idx++
			sidx = 0
			chapter = title
			fmt.Println("chap", chapter)
			return
		}
		sidx++
		fmt.Println(" * page", title)

		dlDir = fmt.Sprintf("%s/%s/%s", usr.HomeDir, "Documents", "reshi-pishi.ru")
		dlAnsDir = dlDir + "/answers"
		if _, err := os.Stat(dlAnsDir); os.IsNotExist(err) {
			_ = os.MkdirAll(dlAnsDir, 0755)
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://reshi-pishi.ru"+href, nil)
		req.Header.Set("Cookie", "subscriber=1; Domain=reshi-pishi.ru; Path=/")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".vari a").Each(func(j int, sel *goquery.Selection) {
			href, _ := sel.Attr("href")
			if strings.Contains(href, ".pdf") && !strings.Contains(href, "javascript") {
				if _, ok := dlReg[href]; ok {
					return
				}
				var destDir string
				if strings.Contains(href, "(ОТВЕТЫ)") {
					destDir = dlAnsDir
				} else {
					destDir = dlDir
				}
				dest := fmt.Sprintf("%s/%02d-%s - %02d-%s", destDir, idx, chapter, sidx, strings.Replace(href, "/sheets/", "", 1))
				href = "https://reshi-pishi.ru" + href
				fmt.Println(" * * file", href, " -> ", dest)
				dlReg[href] = true

				fh, err := os.Create(dest)
				if err != nil {
					log.Fatal(err)
				}
				defer fh.Close()

				resp, err := http.Get(href)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()

				_, err = io.Copy(fh, resp.Body)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	})
}
