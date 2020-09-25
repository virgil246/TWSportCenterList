package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type SportCenter struct {
	ZHname string
	Uri    string
}

func main() {
	res, err := http.Get("https://sc.cyc.org.tw/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []SportCenter
	doc.Find("body > footer > div > div > div.footer_right.col-md-8 > div.col-sm-9.col-md-6").Each(func(i int, s *goquery.Selection) {
		s.Find("p > a").Each(func(i int, s *goquery.Selection) {
			if i > 1 {

				var sportcenter SportCenter
				sportcenter.ZHname = s.Text()
				uri, _ := s.Attr("href")
				sportcenter.Uri = uri
				if sportcenter.ZHname!="南區青少年活動中心"{
					data = append(data, sportcenter)}
			}

		})

	})
	out, _ := json.MarshalIndent(data, "", " ")
	fileWrite(out)
}
func fileWrite(data []byte) {
	// fmt.Println(string(data))

	filename := "AllSportCenter.json"
	ioutil.WriteFile(filename, data, os.ModePerm)
}
