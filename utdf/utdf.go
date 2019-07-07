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
	fname := flag.String("filename", "", "file that contains utdf data")
	rcall := flag.String("run", "ToString", "run specific command")
	ocall := flag.String("output", "utdfgo", "save output of cli to file")
	flag.Parse()

	if *fname == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	f, err := os.Create(*ocall)
	if err != nil {
		log.Fatal("Error writing to file:", *ocall)
	}
	defer f.Close()
	utdf := utdfgo.Run(*fname)
	for _, p := range utdf {
		cmd := reflect.ValueOf(p).MethodByName(*rcall)
		output := cmd.Call([]reflect.Value{})
		fmt.Printf("%v\n", output[0])
		fmt.Fprintln(f, output[0])
	}
	fmt.Printf("fname: %s, rcall: %s, ocall: %s", *fname, *rcall, *ocall)
}
