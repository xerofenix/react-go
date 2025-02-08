package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xerofenix/react-go/database"
)

func init() {
	database.LoadEnv()
	database.ConnectDB()

}

func main() {

	sqlDb, err := database.DBCon.DB()
	if err != nil {
		println("Error while connecting Databse", err)
	}
	defer sqlDb.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html", false)
	})

	app.Listen(":8008")
}
