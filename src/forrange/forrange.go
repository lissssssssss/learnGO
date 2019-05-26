package forrange
import (
	"fmt"
)

func init() {
	data := [3]string{"a", "b", "c"}
	for i, s := range data {
		fmt.Println(i, s)
	}

	for s := range data {
		fmt.Println(s)
	}
}
