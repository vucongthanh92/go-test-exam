package handler

import (
	"github.com/vucongthanh92/go-test-exam/config"
	"github.com/vucongthanh92/go-test-exam/internal/application/cronjob"
)

func Gracefully(cfg *config.AppConfig, cronService cronjob.CronJobService) (err error) {
	return nil
}
