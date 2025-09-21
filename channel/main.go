package main

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {

	stop := make(chan struct{})
	var wg sync.WaitGroup

	wg.Go(func() {
		for {
			select {
			case <-stop:
				log.Info().Msg("stop a goroutine through a channel")
				return
			default:
				log.Info().Msg("working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	})

	time.Sleep(time.Second)
	close(stop)
	wg.Wait()
}
