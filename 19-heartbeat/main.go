package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags)
	ctx, cancel := context.WithTimeout(context.Background(), 5100*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beatInterval = 500 * time.Millisecond
	heartbeat, v := task(ctx, beatInterval)
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
		case <-time.After(wdtTimeout):
			errorLogger.Println("doTask goroutine's heartbeat stopped")
			break loop
		}
	}
}
func task(
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
		sendValue := func(t time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- t:
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
					time.Sleep(1000 * time.Millisecond)
				}
				sendPulse()
				i++
			case t := <-task.C:
				sendValue(t)
			}
		}
	}()
	return heartbeat, out
}
