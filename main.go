package utdfgo

import "fmt"

func main() {
	utdf := Run("")
	for _, p := range utdf {
		fmt.Println(p.ToString())
	}
}
