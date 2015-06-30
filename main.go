// command atpl parses yaml and executes a go text/template against the
// contents of that yaml.
//
//     atpl -input data.yml -template config.tpl
// or
//     atpl -template config.tpl < data.yml
//
// This will parse the yaml data into a generic structure (i.e. interface{},
// []interface{}, and map[string]interface{}) and execute the template
// config.tpl against that data.
package main // import fknsrs.biz/p/atpl

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	inputFile    = flag.String("input", "-", "YAML/JSON file to read input from (default stdin).")
	templateFile = flag.String("template", "", "Template file.")
)

func main() {
	flag.Parse()

	if *templateFile == "" {
		fmt.Fprintln(os.Stderr, "template argument must be provided")
		os.Exit(1)
	}

	templateData, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		panic(err)
	}

	t, err := template.New(*templateFile).Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	input := os.Stdin
	if *inputFile != "-" {
		f, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}

		input = f
	}

	d, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	var v interface{}
	if err := yaml.Unmarshal(d, &v); err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, v); err != nil {
		panic(err)
	}
}
