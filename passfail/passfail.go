//pass_fail reports whether a grade is passing or failing.
package main

import (
	"fmt"
	"log"

	"github.com/xuexj2021/go/keyboard"
)

//main 可以输入用户的分数，并返回是否及格
func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
