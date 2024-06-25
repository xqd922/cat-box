package task

import (
	"github.com/daifiyum/cat-box/singbox"
	"github.com/daifiyum/cat-box/subservice/database"
	"github.com/daifiyum/cat-box/subservice/models"
	"github.com/daifiyum/cat-box/subservice/parser"
	"github.com/daifiyum/cat-box/tray"
	"github.com/daifiyum/cat-box/utils"

	"github.com/robfig/cron/v3"
)

var scheduler *cron.Cron

func handleUpdate() {
	db := database.DB
	var subscriptions []models.Subscriptions
	db.Find(&subscriptions)
	for _, subscription := range subscriptions {
		if subscription.AutoUpdate {
			config, err := parser.Handler(subscription.Link)
			if err != nil {
				utils.LogError("Failed to generate configuration")
				continue
			}
			db.Model(&subscription).Where(subscription.ID).Update("data", config)
			if subscription.Active {
				if tray.GetIsProxy() {
					err = singbox.Reload(string(config))
					if err != nil {
						utils.LogError("Failed to reload configuration")
						continue
					}
				} else {
					singbox.SaveConfig(string(config))
				}
			}
			utils.LogInfo("Automatic update successful")
		}
	}
}

func InitScheduler() {
	db := database.DB
	options := new(models.Options)
	db.Model(options).Where("name=?", "options").First(options)
	scheduler = cron.New()
	scheduler.AddFunc("@every "+options.UpdateDelay, func() {
		handleUpdate()
	})
	scheduler.Start()
}

func Scheduler(delay string) {
	if scheduler != nil {
		scheduler.Stop()
	}

	scheduler = cron.New()
	scheduler.AddFunc("@every "+delay, func() {
		handleUpdate()
	})
	scheduler.Start()
}
