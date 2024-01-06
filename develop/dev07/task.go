package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

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
