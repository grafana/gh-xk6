// Package main contains main function for gh xk6.
package main

import (
	"log"

	"github.com/grafana/gh-xk6/cmd"
)

func main() {
	root := cmd.New()

	if err := root.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
