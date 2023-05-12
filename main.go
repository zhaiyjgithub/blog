package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//// Use sync.Mutex
	processForMutex()
	//time.Sleep(time.Second * 2)
	//// Use sync.Cond
	//processForCond()
	//processRWMutex()
}

type Bucket struct {
	Mutex  sync.Mutex
	Tokens []int
	Done   chan int
}

func processForMutex() {
	bucket := &Bucket{
		Mutex:  sync.Mutex{},
		Tokens: nil,
		Done:   make(chan int),
	}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			go bucket.producer(i)
			//time.Sleep(time.Second * 1)
		}
	}()
	//go bucket.consumer()
	go func() {
		time.Sleep(time.Second * 15)
		fmt.Println("close producer")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("done")
}

func (b *Bucket) producer(idx int) {
	b.Mutex.Lock()
	tokens := []int{rand.Intn(100), rand.Intn(100)}
	fmt.Printf("Put: Index: %d tokens: %v\r\n", idx, tokens)
	b.Tokens = append(b.Tokens, tokens...)
	b.Done <- idx
	fmt.Println("done sent")
	b.Mutex.Unlock()

}

func (b *Bucket) consumer() {
	for {
		fmt.Println("waiting for producer...")
		select {
		case idx := <-b.Done:
			fmt.Println("done received")
			b.Mutex.Lock()
			fmt.Printf("Get: Index: %d tokens: %v\r\n", idx, b.Tokens)
			// Clean tokens
			b.Tokens = nil
			b.Mutex.Unlock()
		}
	}
}

type BucketRW struct {
	Mutex  sync.RWMutex
	Tokens []int
	Done   chan int
}

func processRWMutex() {
	bucket := &BucketRW{
		Mutex:  sync.RWMutex{},
		Tokens: nil,
		Done:   make(chan int),
	}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			go bucket.producer(i)
			time.Sleep(time.Second * 1)
		}
	}()

	go bucket.consumer()

	go func() {
		time.Sleep(time.Second * 15)
		fmt.Println("close producer")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
}

func (b *BucketRW) producer(idx int) {
	b.Mutex.Lock()
	defer func() {
		b.Done <- idx
		b.Mutex.Unlock()
	}()

	tokens := []int{rand.Intn(100), rand.Intn(100)}
	fmt.Printf("Cond-Put: Index: %d tokens: %v\r\n", idx, tokens)
	b.Tokens = append(b.Tokens, tokens...)
}

func (b *BucketRW) consumer() {
	for {
		select {
		case <-b.Done:
			b.Mutex.RLock()
			fmt.Println("Cond-Get: tokens", b.Tokens)
			b.Tokens = nil
			b.Mutex.RUnlock()
		}
	}
}

type BucketCond struct {
	Cond   *sync.Cond
	Tokens []int
	Done   chan int
}

func processForCond() {
	bucket := &BucketCond{
		Cond:   sync.NewCond(&sync.Mutex{}),
		Tokens: nil,
	}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			go bucket.producer(i)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			bucket.consumer()
		}
	}()

	go func() {
		time.Sleep(time.Second * 15)
		fmt.Println("close producer")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("done")
}

func (b *BucketCond) producer(idx int) {
	b.Cond.L.Lock()
	defer b.Cond.L.Unlock()
	tokens := []int{rand.Intn(100), rand.Intn(100)}
	fmt.Printf("Cond-Put: Index: %d tokens: %v\r\n", idx, tokens)
	b.Tokens = append(b.Tokens, tokens...)
	b.Cond.Signal()
}

func (b *BucketCond) consumer() {
	b.Cond.L.Lock()
	defer b.Cond.L.Unlock()

	// Note: must use Wait() in a for express.
	//  Wait() will automatically unlock the b.Cond.L and suspends the current goroutine
	for len(b.Tokens) == 0 {
		b.Cond.Wait()
	}
	fmt.Println("Cond-Get: tokens", b.Tokens)
	// Note: clean tokens. It's a very important code that here is using the token as
	// the condition to wait for lock.
	b.Tokens = nil
}
