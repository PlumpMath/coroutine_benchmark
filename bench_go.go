package main

import "fmt"
import "time"
import "os"
import "strconv"

func worker(begin chan int, quit chan int) {
	boom := time.After(600 * time.Second)
	begin<-0
	for {
		select {
		case <-boom:
			quit<-0
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func schedule(n int, begin chan int, quit chan int) {
	for i:=0; i<n; i++ {
		time.Sleep(1 * time.Millisecond)
		go worker(begin, quit)
	}
}

func info(begin chan int, quit chan int) {
	running := 0
	tick := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-begin:
			running++
		case <-quit:
			running--
		case <-tick:
			fmt.Println("RUNNING ", running)
		}
	}
}

func main() {

	concurrency := 1000
	if len(os.Args)>1 {
		concurrency, _ = strconv.Atoi( os.Args[1] )
	}
	begin := make(chan int)
	quit := make(chan int)
	go schedule(concurrency, begin, quit)
	go info(begin, quit)

	for {
		time.Sleep(10 * time.Second)
	}
}
