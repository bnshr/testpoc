package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	requestId := ctx.Value("request-id")
	fmt.Println(requestId)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out inside doSomethingCool")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go context tutorial")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("oh no ..")
	}

	time.Sleep(5 * time.Second)
}
