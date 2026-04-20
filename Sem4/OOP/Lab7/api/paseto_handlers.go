package api

import (
	"Lab7/shared/configs"
	"Lab7/shared/tockens/auth"
	"Lab7/shared/tockens/models"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type App struct {
	token     *auth.PasetoAuth
	routerApi *fiber.App
	config    configs.Config
}

const (
	authHeader = "Authorization"
	typeBearer = "bearer"
)

func NewApp(config configs.Config, api *fiber.App) *App {

	pasetoToken, err := auth.NewPasseto([]byte(config.TokenKey))
	if err != nil {
		return nil
	}

	app := &App{
		token:     pasetoToken,
		routerApi: api,
		config:    config,
	}
	app.SetApi()
	return app
}

func (a *App) Login(c fiber.Ctx) error {
	creds := new(models.Credentials)

	if err := c.Bind().JSON(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//TODO: Implement login logic

	pasetoToken, err := a.token.NewTocken(models.TockenData{
		Subject:  "for user",
		Duration: a.config.TokenDuration,
		AdditionalClaims: models.AdditionalClaims{
			Name: creds.Username,
			Role: creds.Username,
		},
		Footer: models.Footer{MetaData: "fotter for" + creds.Username},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": pasetoToken})

}

func (a *App) CheckAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		authVal := c.Get(authHeader)

		if len(authVal) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		splitHeader := strings.Fields(authVal)

		if len(splitHeader) < 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		authType := strings.ToLower(splitHeader[0])
		if authType != typeBearer {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		claims, err := a.token.VerifyTocken(splitHeader[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}

func (a *App) SetApi() {
	a.routerApi.Post("/login", a.Login)

	protectedApi := a.routerApi.Group("api", a.CheckAuth())

	//api is protected group
	protectedApi.Get("/account", func(c fiber.Ctx) error {
		val := c.Locals("claims")

		v, ok := val.(models.ServiceClaims)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		fmt.Println(v)
		owner := fmt.Sprintf("<h3>Account owner - %s</h3>", v.Name)
		role := fmt.Sprintf("<h3>Account role - %s</h3>", v.Role)
		footer := fmt.Sprintf("<h3>Account footer - %s</h3>", v.MetaData)

		return c.SendString(owner + role + footer)

	})
}

func (a *App) Start() error {
	return a.routerApi.Listen(a.config.Address)
}
