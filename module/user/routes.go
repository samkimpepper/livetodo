package user

import (
	"github.com/gofiber/fiber/v2"
	"todo/ent"
	"todo/middleware/auth"
)

func Routes(app fiber.Router, db *ent.Client) {
	repo := NewUserRepository(db)
	serv := NewUserService(repo)

	app.Post("/register", registerHandler(serv))
	app.Post("/login", loginHandler(serv))
	app.Get("", auth.HeaderAuthorization(), getUserHandler(serv))
}

func registerHandler(s UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return s.Register(c)
	}
}

func loginHandler(s UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return s.Login(c)
	}
}

func getUserHandler(s UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return s.GetUser(c)
	}
}
