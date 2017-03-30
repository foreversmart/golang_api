package workflow

import (
	"encoding/xml"
	"log"
	"strings"
)

type Content struct {
	Items []*Item `xml:"item"`
}

func (t *Content) Str() string {
	outbytes, err := xml.Marshal(t)
	if err != nil {
		log.Println(err)
		return ""
	}

	output := string(outbytes)
	output = strings.TrimPrefix(output, "<Content>")
	output = strings.TrimSuffix(output, "</Content>")
	return `<?xml version="1.0" encoding="utf-8"?><items>` + output + `</items>`
}

type Item struct {
	Title    string `xml:"title"`
	SubTitle string `xml:"subtitle"`
	Icon     string `xml:"icon"`
	Valid    string `xml:"valid,attr"`
	Arg      string `xml:"arg,attr"`
}
