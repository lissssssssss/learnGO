package main

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
	}
	done <- true
}

func productor(data chan int) {
	for i := 0; i < 3; i++ {
		println("send:", i)
		data <- i
	}
	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)

	go productor(data)
	go consumer(data, done)

	<-done
}
