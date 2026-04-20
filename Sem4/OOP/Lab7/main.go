package main

import (
	"Lab7/api"
	"Lab7/shared/configs"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {

	cfg, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	routerApi := fiber.New()
	app := api.NewApp(cfg, routerApi)
	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
