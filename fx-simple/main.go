package main

import (
	"client-with-fiber/clients"
	"client-with-fiber/database"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	//err := initDatabase2()
	//if err != nil {
	//	panic(err)
	//}
	//// defer database.DBConn.
	//app := fiber.New()
	//
	//setupRoutes(app)
	//
	//app.Get("/", home)
	//
	//log.Fatal(app.Listen(":3000"))

	//+++++++++++++++++++++++++++++++++++++


	app := fx.New(
		// here put all the constructors
		fx.Provide(
			NewApp,
			NewLogger,
		),

		//here we invoke the function to setup our work, which requires the previous constructed objects
		fx.Invoke(
			initDatabase,
		),
		fx.Invoke(setupRoutes),

		fx.WithLogger(
			func( l *log.Logger) fxevent.Logger {
				return &fxevent.ConsoleLogger{W: l.Writer()}
			},
		),

		// this is a special hook to add func on the startup or shutdown (accessing the lifecycle of fx)
		fx.Invoke(register),
	)

	app.Run()
}

func NewLogger() *log.Logger {
	//return log.New(os.Stdout,"client",0)
	return log.Default()
}

func register(lifecycle fx.Lifecycle, app *fiber.App) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go app.Listen(":3000")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("shutting down....")
				return app.Shutdown()
			},
		},
	)
}

func NewApp() *fiber.App {
	fmt.Println("instantiating APP.")
	return fiber.New()
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
	app.Post("/clients/:id", clients.UpdateClient)
	app.Delete("/clients/:id", clients.DeleteClient)
}

func initDatabase(l fxevent.Logger) error {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("can not opend database connection: %v", err)
	}
	err = database.DBConn.AutoMigrate(&clients.Client{})
	if err != nil {
		return fmt.Errorf("migration error: %v", err)
	}
	l.LogEvent(&fxevent.OnStartExecuting{
		FunctionName: "MYinitDatabase",
		CallerName:   "noIdea",
	})
	return nil
}

func initDatabase2() error {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("can not opend database connection: %v", err)
	}
	err = database.DBConn.AutoMigrate(&clients.Client{})
	if err != nil {
		return fmt.Errorf("migration error: %v", err)
	}
	return nil
}
