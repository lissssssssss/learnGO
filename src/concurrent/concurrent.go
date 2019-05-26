package concurrent 
import (
	"fmt"
)

func init() {
	c := make(chan string)
	go func() {
		for i := 1; i < 4; i++ {
			c <- fmt.Sprintf("echo %dth string!",i)
		}
	}()
	go func() {
		s := <- c
		fmt.Println(s)
	}()
}
