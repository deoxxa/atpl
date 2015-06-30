package main

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
