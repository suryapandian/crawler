package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/suryapandian/crawler/config"
	"github.com/suryapandian/crawler/logger"
	"github.com/suryapandian/crawler/worker"
)

func main() {
	logger.SetupLog(config.LOG_LEVEL)

	jobs := []worker.Job{
		worker.NewCrawler(config.URL_TO_CRAWL),
	}
	worker.RunJobs(jobs)

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh

	logger.LogEntryWithRef().Infof("stopping crawler gracefully")
	worker.StopJobs(jobs)
}
