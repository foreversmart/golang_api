package main

import (
	"craw"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"workflow"
)

func main() {
	var input string
	flag.StringVar(&input, "input", "", "query")
	flag.Parse()

	data, err := readCache()
	if err != nil {
		log.Println(err)
		return
	}

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
