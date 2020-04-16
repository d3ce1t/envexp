package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	exp "github.com/d3ce1t/envexp/exporter"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {

	var pretty bool
	var outputType string
	var prefixes arrayFlags
	output := os.Stdout

	flag.StringVar(&outputType, "t", "json", "Output type: web, json")
	flag.BoolVar(&pretty, "p", false, "Enable pretty output")
	flag.Var(&prefixes, "prefix", "Only export env variables matching this prefix")
	flag.Parse()

	env := exp.NewEnvironment()
	filterEnv(env, prefixes)

	exporter := exp.NewExporter(outputType, output, pretty)

	if err := exporter.Export(env); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func filterEnv(env exp.Environment, prefixes []string) {
	for key := range env {
		if len(prefixes) > 0 && !hasPrefixes(prefixes, key) {
			env.Unset(key)
		}
	}
}

func hasPrefixes(a []string, x string) bool {
	for _, n := range a {
		if strings.HasPrefix(x, n) {
			return true
		}
	}
	return false
}
