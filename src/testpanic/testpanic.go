package testpanic
import (
	"runtime/debug"
)

func test() {
	panic("i am dead")
}

func init() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	test()
}
