package main

import (
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func worker() {
	defer log.Info().Msg("deferred operation")
	for i := 0; i < 5; i++ {
		log.Info().Msgf("work: %d", i)
		if i == 2 {
			runtime.Goexit()
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Go(worker)
	wg.Wait()

	log.Info().Msg("done")
}
