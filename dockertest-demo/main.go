package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", home)

	log.Fatal(app.Listen(":2000"))
}

func home(ctx *fiber.Ctx) error {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		return ctx.Send([]byte("fail" + err.Error()))
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return ctx.Send(b)
}
