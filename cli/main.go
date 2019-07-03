package main

import (
	"flag"
	"fmt"
)

func main() {
	fname := flag.String("filename", "", "file that contains utdf data")
	rcall := flag.String("run", "ToString", "run specific command(s) seperated by comma")
	ocall := flag.String("output", "utdfgo", "save output of cli to file")
	flag.Parse()

	fmt.Printf("fname: %s, rcall: %s, ocall: %s", *fname, *rcall, *ocall)
}
