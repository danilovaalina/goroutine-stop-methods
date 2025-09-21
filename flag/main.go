package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	var (
		stop bool
		wg   sync.WaitGroup
	)

	wg.Go(func() {
		for {
			if stop {
				log.Info().Msg("stop goroutine on condition")
				return
			}
			log.Info().Msg("working...")
			time.Sleep(500 * time.Millisecond)
		}
	})

	time.Sleep(time.Second)
	stop = true
	wg.Wait()
}
