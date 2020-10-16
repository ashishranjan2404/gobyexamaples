package main

import (
	"fmt"
	"sync/atomic"
	"sync" 
	"runtime"
)

func main() {
	c := make (chan int64, 5)
	go send(c)
	receive(c)
}

// receive Channel
func receive (c <- chan int64){
	for v:= range c{
		fmt.Println(v)
	}
}


// send channel
func send(c chan<- int64){
	var wg sync.WaitGroup 
	var counter int64
	gor := 10
	
	wg.Add(gor)
	for i:= 0 ; i < gor ; i++ {
		go func (){
			for i:= 0; i<gor ; i++{
				atomic.AddInt64(&counter,1)			
				c <- atomic.LoadInt64(&counter)
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)	
	
}
