package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"btpgo/models"
)

func main() {
	app := fiber.New()
	dsn := "host=postgres-05b4967f-26e8-46c6-9236-60de13fcb364.cqryblsdrbcs.us-east-1.rds.amazonaws.com user=ca44c9e5bd49 password=851bd5f7dbc5508a9d2da dbname=zLgbjsopIMDP port=2643"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, BTP!")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User;
		db.Find(&users);
		return c.JSON(users)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		user := new(models.User);
		if err := c.BodyParser(user); err != nil {
			return c.Status(503).SendString(err.Error())
		}
	
		db.Create(&user)
		return c.Status(201).JSON(user)
	});

	


	app.Listen(":8080")
}