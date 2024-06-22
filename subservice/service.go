package subservice

import (
	"os"

	"github.com/daifiyum/cat-box/config"
	"github.com/daifiyum/cat-box/subservice/database"
	"github.com/daifiyum/cat-box/subservice/router"
	"github.com/daifiyum/cat-box/task"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SubService() {
	// write logs to app.log
	file, _ := os.OpenFile(config.Config("LOG_PATH"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// -H=windowsgui 输出到终端后无法打印日志，故注释下面一行
	// iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(file)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	if err := database.ConnectDB(); err != nil {
		return
	}

	router.SetupRoutes(app)
	task.InitScheduler()

	app.Listen(":3000")
}
