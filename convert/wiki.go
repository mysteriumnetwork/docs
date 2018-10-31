package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type Pages struct {
	Pages []yaml.MapSlice `yaml:"pages"`
}

var (
	input  = flag.String("input", "mkdocs.yml", "input file name")
	output = flag.String("output", "docs/_Sidebar.md", "output file name")
)

func init() {
	flag.Parse()
}

func main() {
	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}

	var p Pages
	if err := yaml.Unmarshal(data, &p); err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	f, err := os.Create(*output)
	if err != nil {
		log.Fatalf("cannot create file: %v", err)
	}

	for _, item := range p.Pages {
		for _, v := range item {
			link := reflect.ValueOf(v.Value)

			switch link.Kind() {
			case reflect.String:
				fmt.Fprintf(f, "**[%s](%s)**\n", v.Key, trimLink(v.Value))

			case reflect.Slice:
				fmt.Fprintf(f, "\n<details><summary><a href=#><b>%s</b></a></summary>\n\n", v.Key)
				for i := 0; i < link.Len(); i++ {
					for _, v := range link.Index(i).Interface().(yaml.MapSlice) {
						fmt.Fprintf(f, "* [%s](%s)\n", v.Key, trimLink(v.Value))
					}
				}
				fmt.Fprintf(f, "</details>\n")
			}
		}
	}
}

func trimLink(path interface{}) string {
	s := fmt.Sprint(path)
	s = filepath.Base(s)
	return strings.TrimSuffix(s, ".md")
}
