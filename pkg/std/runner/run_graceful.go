package runner

import (
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
)

type GracefulRunner interface {
	Init() error
	Run() error
	Shutdown()
}

func RunGraceful(runner GracefulRunner) {
	log.Info("Starting...")

	if err := runner.Init(); err != nil {
		log.Fatal("Init", "error", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	go func() {
		if err := runner.Run(); err != nil {
			log.Fatal("Run", "error", err)
		}
		sigs <- os.Interrupt
	}()

	done := make(chan struct{})
	go func() {
		<-sigs
		runner.Shutdown()
		done <- struct{}{}
	}()
	<-done

	log.Info("Bye!")
}
