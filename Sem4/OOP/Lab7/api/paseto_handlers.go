package api

import (
	"Lab7/controllers"
	entities "Lab7/models"
	"Lab7/shared/configs"
	"Lab7/shared/crypto"
	"Lab7/shared/tockens/auth"
	"Lab7/shared/tockens/models"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type App struct {
	token     *auth.PasetoAuth
	routerApi *fiber.App
	config    configs.Config

	db       *controllers.Dbontroller
	promlems *controllers.ProblemsController
	slots    *controllers.AvalTimeController
	session  *controllers.SessionController

	ctx context.Context
}

const (
	authHeader = "Authorization"
	typeBearer = "bearer"
)

func NewApp(config configs.Config, api *fiber.App) (*App, error) {

	ctx := context.Background()

	pasetoToken, err := auth.NewPasseto([]byte(config.TokenKey))
	if err != nil {
		return nil, err
	}

	db, err := controllers.NewDbontroller(ctx)
	if err != nil {
		return nil, err
	}

	problems := controllers.NewProblemsController(db.Problems)

	slots := controllers.NewAvalTimeController(db.Avalavility)

	session := controllers.NewSessionController(db.Sessions, db.Problems)

	app := &App{
		token:     pasetoToken,
		routerApi: api,
		config:    config,
		db:        db,
		ctx:       ctx,
		promlems:  problems,
		slots:     slots,
		session:   session,
	}
	app.SetApi()
	return app, nil
}

// Login godoc
// @Summary      Login user
// @Description  Login user and return pasetto tocken
// @Tags         log in
// @Accept       json
// @Produce      json
//
// @Param        request  body  models.Credentials  true  "Sign in info"
//
// @Success      200
// @Failure      400
// @Failure      500
//
// @Router      /login [post]
func (a *App) Login(c fiber.Ctx) error {
	creds := new(models.Credentials)
	if err := c.Bind().JSON(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashPass, err := a.db.Interwiewers.GetPassHash(a.ctx, creds.Username)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if hashPass == "" {
		return c.Status(fiber.StatusUnauthorized).JSON("You are not signed in")
	}

	if !crypto.CheckPasswordHash(creds.Password, hashPass) {
		return c.Status(fiber.StatusUnauthorized).JSON("Invalid password")
	}

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

// Register godoc
// @Summary      Register user
// @Description  Returns nil error of success
// @Tags         sign in
// @Accept       json
// @Produce      json
//
// @Param        request  body  models.RegisterInfo  true  "Sign in info"
//
// @Success      200
// @Failure      400
// @Failure      500
//
// @Router      /register [post]
func (a *App) Register(c fiber.Ctx) error {
	creds := new(models.RegisterInfo)

	if err := c.Bind().JSON(creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hash, err := crypto.HashPassword(creds.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	err = a.db.Interwiewers.Update(a.ctx, &entities.Interviewer{
		0,
		creds.Username,
		hash,
		creds.Email,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (a *App) CheckAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		authVal := c.Get(authHeader)

		if len(authVal) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		splitHeader := strings.Fields(authVal)

		claims, err := a.token.VerifyTocken(splitHeader[0])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}

func (a *App) SetApi() {
	a.routerApi.Post("/login", a.Login)
	a.routerApi.Post("/register", a.Register)

	protectedApi := a.routerApi.Group("api", a.CheckAuth())
	protectedApiSession := protectedApi.Group("session")
	protectedApiCandidates := protectedApi.Group("candidates")
	protectedApiProblmes := protectedApi.Group("problems")
	protectedApiSlots := protectedApi.Group("slots")

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

	protectedApiSession.Post("/create", a.CreateSession)
	protectedApiSession.Get("/all", a.GetAllSessions)
	protectedApiSession.Post("/get", a.GetSessionById)
	protectedApiSession.Delete("/:id", a.DeleteSessionById)
	protectedApiSession.Post("/problems", a.AddProblemsToSession)
	protectedApiSession.Post("/result", a.SetResultHandler)

	protectedApiSession.Post("/:id/start", a.StartSession)
	protectedApiSession.Post("/:id/stop", a.StopSession)

	protectedApiCandidates.Post("/add", a.AddCandidateHandler)
	protectedApiCandidates.Delete("/:id", a.DeleteCandidateHandler)

	protectedApiProblmes.Post("/add", a.AddProblemHandler)
	protectedApiProblmes.Delete("/:id", a.DeleteProblemHandler)
	protectedApiProblmes.Get("/all", a.GetAllProblems)
	protectedApiProblmes.Get("/:id", a.GetProblemByIdHandler)

	protectedApiSlots.Post("/book", a.BookSlotsHandler)
	protectedApiSlots.Post("/unbook", a.UnookSlotsHandler)
	protectedApiSlots.Delete("/:id", a.DeleteSlotsByIdHandler)
	protectedApiSlots.Post("/add", a.AddSlotHandler)

}

func (a *App) Start() error {
	return a.routerApi.Listen(a.config.Address)
}
