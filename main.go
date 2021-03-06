package main

//"topmusicstreaming/collector"

import (
	"net/http"
	"topmusicstreaming/api"
	"topmusicstreaming/hub"

	"github.com/robfig/cron/v3"
)

func main() {

	cUS := cron.New()
	cUS.AddFunc("CRON_TZ=Europe/Paris 30 15 * * *", func() { hub.Hub_US() })
	cUS.Start()

	cFR := cron.New()
	cFR.AddFunc("CRON_TZ=Europe/Paris 30 16 * * *", func() { hub.Hub_FR() })
	cFR.Start()

	cDE := cron.New()
	cDE.AddFunc("CRON_TZ=Europe/Paris 30 17 * * *", func() { hub.Hub_DE() })
	cDE.Start()

	cES := cron.New()
	cES.AddFunc("CRON_TZ=Europe/Paris 30 18 * * *", func() { hub.Hub_ES() })
	cES.Start()

	cPT := cron.New()
	cPT.AddFunc("CRON_TZ=Europe/Paris 30 19 * * *", func() { hub.Hub_PT() })
	cPT.Start()

	cIT := cron.New()
	cIT.AddFunc("CRON_TZ=Europe/Paris 30 20 * * *", func() { hub.Hub_IT() })
	cIT.Start()

	http.HandleFunc("/api", api.Api)
	http.ListenAndServe(":9990", nil)

}
