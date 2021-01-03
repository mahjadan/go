package main

import (
	"client-with-fiber/clients"
	"client-with-fiber/database"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	initDatabase()
	// defer database.DBConn.
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", home)

	log.Fatal(app.Listen(":3000"))
}

func home(c *fiber.Ctx) error {
	return c.Send([]byte("hello there.."))
}

func setupRoutes(app *fiber.App) {

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${pid} tid: ${locals:requestid} latency:${latency} ${status} - ${method} ${path}​\n",
		TimeFormat: "02-Jan-2006-24:60:60:00000",
	}))

	app.Get("/clients", clients.ListClients)
	app.Get("/clients/:id", clients.GetClient)
	app.Post("/clients", clients.CreateClient)
	app.Post("clients/:id", clients.UpdateClient)
	app.Delete("/clients/:id", clients.DeleteClient)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("can not opend database connection,")
	}
	database.DBConn.AutoMigrate(&clients.Client{})
	fmt.Println("database migrated successfully..")
}
