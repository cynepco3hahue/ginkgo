package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func BuildUnfocusCommand() *Command {
	return &Command{
		Name:         "unfocus",
		AltName:      "blur",
		FlagSet:      flag.NewFlagSet("unfocus", flag.ExitOnError),
		UsageCommand: "ginkgo unfocus (or ginkgo blur)",
		Usage: []string{
			"Recursively unfocuses any focused tests under the current directory",
		},
		Command: unfocusSpecs,
	}
}

func unfocusSpecs([]string, []string) {
	unfocus("Describe")
	unfocus("Context")
	unfocus("It")
	unfocus("Measure")
	unfocus("DescribeTable")
	unfocus("Entry")
	unfocus("Specify")
	unfocus("When")
}

func unfocus(component string) {
	fmt.Printf("Removing F%s...\n", component)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		println(err.Error())
		return
	}
	for _, f := range files {
		if f.Name() == "vendor" {
			continue
		}
		cmd := exec.Command("gofmt", fmt.Sprintf("-r=F%s -> %s", component, component), "-w", f.Name())
		out, _ := cmd.CombinedOutput()
		if string(out) != "" {
			println(string(out))
		}
	}
}
