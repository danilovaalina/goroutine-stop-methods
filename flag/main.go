package main

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	var (
		stop atomic.Bool
		wg   sync.WaitGroup
	)

	wg.Go(func() {
		for {
			if stop.Load() {
				log.Info().Msg("stop goroutine on condition")
				return
			}
			log.Info().Msg("working...")
			time.Sleep(500 * time.Millisecond)
		}
	})

	time.Sleep(time.Second)
	stop.Store(true)
	wg.Wait()
}
