package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msdelay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(msdelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	webServer := sampleOperation(ctx, "webserver", 200)
	microService := sampleOperation(ctx, "microService", 200)
	database := sampleOperation(ctx, "database", 200)
MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microService:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}
	}
	fmt.Println(<-database)
}
