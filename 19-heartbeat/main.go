package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 11500*time.Millisecond)
	defer cancel()
	const timeout = 2 * time.Second
	const beatInterval = 1 * time.Second
	heartbeat, v := doTask(ctx, beatInterval)
loop:
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				break loop
			}
			fmt.Println("beat pulse ⚡️")
		case r, ok := <-v:
			if !ok {
				break loop
			}
			t := strings.Split(r.String(), "m=")
			fmt.Printf("value: %v [s]\n", t[1])
		case <-time.After(timeout):
			fmt.Println("doTask goroutine's heartbeat stopped")
			break loop
		}
	}
}
func doTask(
	ctx context.Context,
	beatInterval time.Duration,
) (<-chan struct{}, <-chan time.Time) {
	heartbeat := make(chan struct{})
	out := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(2 * beatInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendValue := func(r time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- r:
					return
				}
			}
		}
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				if i == 3 {
					time.Sleep(3000 * time.Millisecond)
				}
				sendPulse()
				i++
			case r := <-task.C:
				sendValue(r)
			}
		}
	}()
	return heartbeat, out
}
