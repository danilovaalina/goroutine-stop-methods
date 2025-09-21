package main

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("stop goroutine by context")
				return
			default:
				log.Info().Msg("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	})

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
}
