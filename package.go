package main

import (
	"craw"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"match"
	"os"
	"strings"
	"workflow"
)

func main() {
	var input string
	var member string
	var isMember bool
	flag.StringVar(&input, "input", "", "query")
	flag.Parse()
	input = strings.Replace(input, "  ", " ", -1)
	input = strings.TrimSpace(input)
	inputs := strings.Split(input, " ")
	if len(inputs) > 1 {
		isMember = true
		member = inputs[1]
		input = inputs[0]
	}

	// read cache
	data, err := readCache()
	if err != nil {
		log.Println(err)
		return
	}

	// recover cache
	var pkgs, descs []string
	if data == nil {
		pkgs, descs = craw.GetPkgsAndDesc("https://golang.org/pkg/")
		data = &Data{
			Pkgs:  pkgs,
			Descs: descs,
		}

		data.Flush()
	} else {
		pkgs = data.Pkgs
		descs = data.Descs
	}

	res := &workflow.Content{
		Items: make([]*workflow.Item, 0, 5),
	}

	matchPkgs := match.MatchArr(input, pkgs)

	// deal member
	if isMember {
		links, members, descs := craw.CrawPkgMembers("https://golang.org/pkg/" + matchPkgs[0])

		// make reverse map
		reverseMap := make(map[string]int)
		for i, value := range members {
			reverseMap[value] = i
		}

		matchMembers := match.MatchArr(member, members)

		for _, value := range matchMembers {
			res.Items = append(res.Items, &workflow.Item{
				Title:    value,
				Valid:    "yes",
				SubTitle: descs[reverseMap[value]],
				Icon:     "icon.png",
				Arg:      matchPkgs[0] + "/" + links[reverseMap[value]],
			})
		}

		fmt.Println(res.Str())
		return
	}

	pkgMap := make(map[string]string)
	for i, pkg := range pkgs {
		pkgMap[pkg] = descs[i]
	}

	for _, pkg := range matchPkgs {
		res.Items = append(res.Items, &workflow.Item{
			Title:    pkg,
			Valid:    "yes",
			SubTitle: pkgMap[pkg],
			Icon:     "icon.png",
			Arg:      pkg,
		})
	}

	fmt.Println(res.Str())
}

func readCache() (data *Data, err error) {
	path, err := cachePath()
	if err != nil {
		return data, err
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}

	if len(content) == 0 {
		return
	}

	err = json.Unmarshal(content, &data)
	return
}

type Data struct {
	Pkgs  []string `json:"pkgs"`
	Descs []string `json:"descs"`
}

func (data *Data) Flush() error {
	path, err := cachePath()
	if err != nil {
		return err
	}

	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, content, 0666)
}

func cachePath() (path string, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return pwd + "/.cache", nil
}
