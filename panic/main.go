package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func worker() {
	defer func() {
		if r := recover(); r != nil {
			log.Info().Msgf("caught panic in worker:: %v", r)
		}
	}()

	for i := 0; i < 5; i++ {
		log.Info().Msgf("work: %d", i)
		if i == 2 {
			panic("stop goroutine by panic")
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
