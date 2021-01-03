package clients

import (
	"client-with-fiber/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name      string `json:"name,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	CPF       string `json:"cpf,omitempty"`
}

func ListClients(c *fiber.Ctx) error {
	db := database.DBConn
	clients := []Client{}

	db.Find(&clients)

	return c.JSON(clients)
}

func CreateClient(c *fiber.Ctx) error {
	var client Client
	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DBConn.Create(&client)
	return c.JSON(client)
}

func GetClient(c *fiber.Ctx) error {
	id := c.Params("id")

	client := Client{}
	database.DBConn.First(&client, id)
	if client.Name == "" {
		return c.Status(404).JSON("Not Found")
	}
	return c.JSON(client)
}

func UpdateClient(c *fiber.Ctx) error {
	var newData Client
	if err := c.BodyParser(&newData); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	id := c.Params("id")
	var client Client

	database.DBConn.First(&client, id)
	if client.Name == "" {
		return c.Status(404).JSON("Not Found")
	}
	database.DBConn.Model(&client).Updates(newData)
	return c.JSON(client)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	var client Client
	database.DBConn.First(&client, id)

	if client.Name == "" {
		return c.Status(404).JSON("Not Found")
	}
	database.DBConn.Delete(&client)
	return c.JSON("client deleted")
}
