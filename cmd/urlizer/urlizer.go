package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if err := checkArgs(args); err != nil {
		panic(err.Error())
	}

	inputs := strings.Split(args[0], ",")
	template := strings.Trim(args[1], "")

	fmt.Print(strings.TrimSuffix(inject(inputs, template), ","), "")
}

func inject(inputs []string, template string) string {
	var url string
	for _, v := range inputs {
		url += fmt.Sprintf("%s,\n", strings.Replace(template, "{}", v, -1))
	}
	return url
}

func checkArgs(args []string) error {
	if len(args) != 2 {
		return errors.New("You must give 2 inputs, a list of inputs, and a url template ")
	}
	if len(strings.Split(args[0], ",")) < 1 {
		return errors.New("this must be a comma seperated list")
	}
	return nil
}
