/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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

	index = flag.String("index", "docs/index.md", "index file name")
)

func init() {
	flag.Parse()
}

func main() {
	if err := os.Rename(*index, filepath.Join(filepath.Dir(*index), "Home.md")); err != nil {
		log.Fatalf("cannot rename index file: %v", err)
	}

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
				path := strings.Replace(trimLink(v.Value), "index", "Home", 1)
				fmt.Fprintf(f, "**[%s](%s)**\n", v.Key, path)

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
