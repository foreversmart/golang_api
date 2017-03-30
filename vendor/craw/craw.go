package craw

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	query "github.com/foreversmart/goquery"
	"strings"
)

func GetPkgsAndDesc(url string) (pkgs []string, descs []string) {
	doc, err := NewDocumentFromString(Craw("https://golang.org/pkg/"))
	if err != nil {
		log.Println(err)
		return
	}

	packagesArr := make([]string, 0, 5)
	doc.Find(".pkg-name").Each(func(i int, s *query.Selection) {
		if i != 0 {
			pkgName, _ := s.Html()
			pkgName = strings.TrimSpace(pkgName)
			pkgName = strings.Split(pkgName, `"`)[1]
			packagesArr = append(packagesArr, pkgName)
		}
	})

	packagesDescArr := make([]string, 0, 5)
	doc.Find(".pkg-synopsis").Each(func(i int, s *query.Selection) {
		if i != 0 {
			desc, _ := s.Html()
			desc = strings.TrimSpace(desc)
			packagesDescArr = append(packagesDescArr, desc)
		}
	})

	return packagesArr, packagesDescArr
}

func Craw(url string) string {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", url, nil)
	// ...
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Request page error at:", url)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func NewDocumentFromString(content string) (*query.Document, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(content)
	doc, err := query.NewDocumentFromReader(buf)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
