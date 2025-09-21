package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info().Msgf("stop goroutine by context with deadline")
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
