package agent

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func SignalHandler() {
	log.Print("Sitting in signal loop")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	exitChan := make(chan int)

	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGINT:
				log.Print("Exiting on SIGINT")
				exitChan <- 0
			case syscall.SIGTERM:
				log.Print("Exiting on SIGTERM")
				exitChan <- 0
			default:
				log.Error().
					Str("signal", fmt.Sprintf("%v", s)).
					Msg("Exiting on some crazy signal")
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	log.Info().
		Int("exit code", code).
		Msg("Exiting")
	os.Exit(code)
}
