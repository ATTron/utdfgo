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
	rcall := flag.String("run", "ToString", "run specific command(s) seperated by comma")
	ocall := flag.String("output", "utdfgo", "save output of cli to file")
	flag.Parse()

	f, err := os.OpenFile(*ocall, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", *ocall)
	}
	defer f.Close()
	utdf := utdfgo.Run(*fname)
	for _, p := range utdf {
		cmd := reflect.ValueOf(p).MethodByName(*rcall)
		output := cmd.Call([]reflect.Value{})
		line, _ := fmt.Printf("%v\n", output[0])
		fmt.Println(line)
		// TODO: handle saving to output file
		// if _, err := f.WriteString(string(line)); err != nil {
		// 	log.Fatal("Cannot write line(s) to file", err)
		// }
	}
	fmt.Printf("fname: %s, rcall: %s, ocall: %s", *fname, *rcall, *ocall)
}
