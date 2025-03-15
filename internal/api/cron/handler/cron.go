package handler

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jasonlvhit/gocron"
	"github.com/vucongthanh92/go-base-utils/logger"
	"github.com/vucongthanh92/go-test-exam/config"
	"github.com/vucongthanh92/go-test-exam/helper/utils"
	"github.com/vucongthanh92/go-test-exam/internal/application/cronjob"
)

func Crawl(
	cfg *config.AppConfig,
	cronService cronjob.CronJobService,
) {
	ctx, cancel := context.WithCancel(context.Background())
	shutdownCh := make(chan struct{})

	// Listen for interrupt signals to initiate graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh

		logger.Info("Received interrupt signal. Shutting down gracefully...")
		close(shutdownCh) // Signal the shutdown
		cancel()
	}()

	utils.SafeGo(func() {
		StartServices(cronService)
	})

	// Wait for the context to be canceled (either by the scheduler or interrupt signal)
	<-ctx.Done()
	// Ensure that the scheduler is stopped before updating UserJob
	gocron.Clear()

	if err := Gracefully(cfg, cronService); err != nil {
		logger.Error(`error endGracefully`)
		return
	}

	logger.Info("Shutting down Crawl Data gracefully...")
}
