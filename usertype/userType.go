//keyword type can create user type base on basic type, struct, function, other user type and so on.
package usertype
import (
	"fmt"
)

//a user type based on byte
type flags byte

const (
	//iota is a index of const
	read flags = 1 << iota
	//const can omit type and value, and it is the same as last const('s expression)
	write
	exec
)

func init() {
	f := read | exec
	fmt.Printf("%b\n", f)
}

//test doc
func TestDoc() {
	fmt.Println("TestDoc")
}
