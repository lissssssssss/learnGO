package escapeAnalysis

import (
	"fmt"
)

func test() *int {
	a := 0x100
	return &a
}

func Init() {
	var a *int = test()
	fmt.Println(a, *a) 
}
