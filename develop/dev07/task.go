package main

import (
	"fmt"
	"math/rand"
	"time"
)
//comment
func getChan(name string) <- chan string{
	ch := make(chan string, 10)
	go func(){
		for i := 0; i < 5; i++{
			time.Sleep(time.Second * time.Duration((rand.Intn(10) + 1)))
			ch <- name
		}
		close(ch)
	}()
	return ch
}

func or(channels ... <- chan string) chan string{
	ch := make(chan string)
	closeChan := make(chan struct{})
	for _, c := range channels{
		go func(c <-chan string){
			for v := range c{
				ch <- v
			}
			closeChan <- struct{}{}
		}(c)
	}
	go func(){
		count := 0
		for range closeChan {			
			count++
			if count == len(channels){
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	ch1 := getChan("n1")
	ch2 := getChan("n2")
	ch3 := getChan("n3")

	m := or(ch1, ch2, ch3)

	for v := range m{
		fmt.Println(v)
	}
}
