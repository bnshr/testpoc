package main

import (
	"context"
	"fmt"
)

func work(ctx context.Context) {
	fmt.Println("Doing something")
}

func main_context() {
	ctx := context.TODO()
	work(ctx)
}
