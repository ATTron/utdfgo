package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/attron/utdfgo"
)

func main() {
	// setup cli flags
	fname := flag.String("filename", "", "file that contains utdf data")
	rcall := flag.String("run", "ToString", "run specific command")
	ocall := flag.String("output", "", "save output of cli to file")
	flag.Parse()

	if *fname == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	// run commands and output to file if requested
	utdf, err := utdfgo.Run(*fname)
	if err != nil {
		log.Fatal("Unable to process:", err)
	}
	for i, p := range utdf {
		cmd := reflect.ValueOf(p).MethodByName(*rcall)
		output := cmd.Call([]reflect.Value{})
		fmt.Printf("%v\n", output[0])
		if *ocall != "" {
			if i == 0 {
				_, err := os.Create(*ocall)
				if err != nil {
					log.Fatal("Cannot create file", *ocall)
				}
			}
			f, err := os.OpenFile(*ocall, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			fmt.Fprintln(f, output[0])
		}
	}
	// useful for debugging
	// fmt.Printf("fname: %s, rcall: %s, ocall: %s", *fname, *rcall, *ocall)
}
