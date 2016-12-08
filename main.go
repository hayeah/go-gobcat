package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"

	"fmt"

	"github.com/kr/pretty"
)

func main() {
	var input io.Reader

	argc := len(os.Args)
	if argc == 1 {
		input = os.Stdin
	} else if argc == 2 {
		if os.Args[1] == "-h" {
			help()
			os.Exit(0)
		}

		filepath := os.Args[1]
		f, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}
		input = f
	} else {
		fmt.Fprintln(os.Stderr, "Can dump only one file at a time")
		help()
		os.Exit(1)
	}

	d := gob.NewDecoder(input)

	var o map[string]interface{}
	err := d.Decode(&o)

	if err != nil {
		log.Fatalln(err)
	}

	pretty.Println(o)

}

func help() {
	msg := `
	gobcat [file]

	Pretty print Golang gob from file or standard input
	`

	fmt.Fprintln(os.Stdout, msg)
}
