//make declear a slice,map or channel
//new(var) declear all type but,not initializate reference type.
package referenceType
import (
	"fmt"
	"reflect"
)

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 100)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}

func mkmapByVar() map[string]int {
	//var m map[string]int
	m := map[string]int{"b":1}
	m["a"] = 2
	return m
}

func mksliceBynew() []int {
	s := new([]int)
	fmt.Println("type:", reflect.TypeOf(s))
	return *s
}

func mksliceByVar() []int {
	var s []int
	s = append(s, 100)
	return s
}

func init() {
	fmt.Println("START!")
	m := mkmapByVar()
	print(m["a"])
	print(m["b"])

	s := mksliceByVar()
	println(s[0])
	s = mksliceBynew()
	fmt.Println("START2!")
	newInt := *new(int)
	newInt = 1
	println(newInt)
}
