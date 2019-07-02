package utdfgo

import "fmt"

func main() {
	utdf := run("")
	for _, p := range utdf {
		fmt.Println(p.toString())
	}
}
