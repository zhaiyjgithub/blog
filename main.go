package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

func main()  {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ctx := context.Background()
	if cmd := rdb.Ping(ctx); cmd.Val() != "PONG" {
		log.Fatalln(cmd.Val())
	}

	// Set redis config to enable keyspace event notifications. And the event type is "KEA".
	// It means will publish most type of events.
	if _, err := rdb.Do(ctx, "CONFIG", "SET", "notify-keyspace-events", "KEA").Result(); err != nil {
		log.Fatalln(err)
	}

	// Set redis to subscribe the event with expired type.
	pSub := rdb.PSubscribe(ctx, "__keyevent@0__:expired")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(p *redis.PubSub) {
		for {
			receiveKeyEvent(p)
		}
	}(pSub)

	go func() {
		for i := 0; i < 5; i ++ {
			key := fmt.Sprintf("event_key_%d", i)
			rdb.Set(ctx, key, i, time.Second * 5)
		}
	}()

	go func() {
		time.Sleep(time.Second * 8)
		wg.Done()
		fmt.Println("Done")
	}()
	wg.Wait()
}

func receiveKeyEvent(p *redis.PubSub)  {
	message, err := p.ReceiveMessage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message.String(), message.Payload, message.Channel)
}
