package todo

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"todo/ent"
	"todo/middleware/auth"
	"todo/module/user"
)

func Routes(app fiber.Router, db *ent.Client) {
	repo := NewTodoRepository(db)
	userRepo := user.NewUserRepository(db)
	serv := NewTodoService(repo, userRepo)

	//app.Get("/todo", getUserHandler(serv))
	app.Post("/todo", auth.HeaderAuthorization(), createTodoHandler(serv))
	//app.Get("/ws/todo-list/:id", websocket.New(connectTodoList))
	app.Get("/ws/todo-list/:id", websocket.New(func(conn *websocket.Conn) {
		connectTodoList(conn, repo)
	}))

	//app.Put("/todo/:id", updateTodoHandler(serv))
	//app.Delete("/todo/:id", deleteTodoHandler(serv))
}

func createTodoHandler(serv TodoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return serv.CreateTodoList(c)
	}
}
