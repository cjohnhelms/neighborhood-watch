package main

import (
	"fmt"

	"github.com/cjohnhelms/sentinel/pkg/config"
	"github.com/cjohnhelms/sentinel/pkg/notify"
	"github.com/cjohnhelms/sentinel/pkg/scraper"

	log "github.com/cjohnhelms/sentinel/pkg/logging"
)

func main() {
	cfg := config.New()

	log.Info("Service starting")
	log.Debug(fmt.Sprintf("Config: %+v", cfg))

	ch := make(chan scraper.Event, 1)
	go scraper.FetchEvents(ch)
	go notify.Notify(ch, cfg)
	//go display.Update(ch)

	select {}
}
