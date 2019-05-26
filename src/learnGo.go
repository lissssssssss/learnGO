//asdfasdf
package main
import (
	"fmt"
	//"os"
	//"sync"
	//_ "referenceType"
	//_ "usertype"
	//_ "forrange"
	//"escapeAnalysis"
	//_ "testpanic"
	//_ "my_slice"
	//_ "concurrent"
	"testmethod"
	"reflect"
)


func main() {
	//taskCount := 40
	//var wg sync.WaitGroup
	//task := make(chan string)

	//for i := 1; i < taskCount; i++ {
	//	go func() {
	//		s := <-task
	//		fmt.Println(s, os.Getpid())
	//		wg.Done()
	//	}()
	//}
	//for i := 1; i < taskCount; i++ {
	//	wg.Add(1)
	//	task <- fmt.Sprintf("echo %dth string!",i)
	//}
	//wg.Wait()

	fmt.Println("end!")
	var a testmethod.T
	t := reflect.TypeOf(a)
	m := t.Method(0)
	fmt.Println(m.Name, m.Type)
	fmt.Println("*******************")
}
