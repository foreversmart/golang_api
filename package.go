package main

import (
	"craw"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"qbox.us/encoding/json"
	"strings"
	"workflow"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "", "query")
	flag.Parse()

	pkgs, descs := craw.GetPkgsAndDesc("https://golang.org/pkg/")

	res := &workflow.Content{
		Items: make([]*workflow.Item, 0, 5),
	}

	for i, pkg := range pkgs {
		if strings.HasPrefix(pkg, input) {
			res.Items = append(res.Items, &workflow.Item{
				Title:    pkg,
				Valid:    "yes",
				SubTitle: descs[i],
				Icon:     "icon.png",
				Arg:      pkg,
			})
		}
	}

	fmt.Println(res.Str())
}


