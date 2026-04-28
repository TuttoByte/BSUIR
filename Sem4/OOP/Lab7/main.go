package main

import (
	"Lab7/api"
	"Lab7/shared/configs"
	"github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"log"
)

// @title           Interwier
// @version         1.0
// @BasePath
// @securityDefinitions.apikey PASETOAuth
// @in header
// @name Authorization
// @description Введите "Bearer <PASETO токен>". Получите токен на /login и прикрепите его.
func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	cfg, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	routerApi := fiber.New()

	//Swagger
	cfgSwagger := swaggerui.Config{
		BasePath: "/",
		FilePath: "docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	app, err := api.NewApp(cfg, routerApi)
	if err != nil {
		log.Fatal(err)
	}

	routerApi.Use(swaggerui.New(cfgSwagger))

	err = app.Start()
	if err != nil {
		log.Fatal(err)
	}

}
