package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("stop goroutine by context with timeout")
				return
			default:
				log.Info().Msg("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	<-ctx.Done()
}
