package craw

import (
	"log"

	query "github.com/foreversmart/goquery"
	"strings"
)

func CrawPkgMembers(url string) (links, members, descs []string) {
	doc, err := NewDocumentFromString(Craw(url))
	if err != nil {
		log.Println(err)
		return
	}

	members = make([]string, 0, 5)
	links = make([]string, 0, 5)
	descs = make([]string, 0, 5)
	doc.Find("#manual-nav").Find("dl").Find("dd").Each(func(i int, s *query.Selection) {
		// fetch member
		member := s.Text()
		member = strings.TrimSpace(member)
		members = append(members, member)

		// fetch member link
		link := ""
		s.Find("a").Each(func(i int, se *query.Selection) {
			link, _ = se.Attr("href")
		})

		links = append(links, link)
		link = strings.Replace(link, ".", `\.`, -1)

		// fetch member description
		desc := ""
		next := doc.Find(link).Next()
		if next.Is("p") {
			desc = next.Text()
		} else if next.Next().Is("p") {
			desc = next.Next().Text()
		}

		desc = strings.TrimSpace(desc)
		descs = append(descs, desc)

	})

	return
}

func MemberUrl(member string) string {
	return ""
}
