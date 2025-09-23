package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		ticker, timeout := time.NewTicker(500*time.Millisecond), time.After(3*time.Second)

		for count := 1; ; count++ {
			select {
			case <-ticker.C:
				log.Info().Msgf("tick: %d", count)
			case <-timeout:
				log.Info().Msg("stop goroutine by timeout")
				return
			}
		}
	})

	wg.Wait()
}
