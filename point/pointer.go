package point

import (
	"fmt"
)

func main() {
	var myInt int = 42
	var myIntPointer *int = &myInt
	*myIntPointer = myInt
	fmt.Println(*myIntPointer)
}
