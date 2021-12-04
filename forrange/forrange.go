package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}

	for _, node := range notes {
		fmt.Println(node)
	}
}
