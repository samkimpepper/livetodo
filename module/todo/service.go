package todo

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"todo/module/user"
)

type TodoService interface {
	CreateTodoList(c *fiber.Ctx) error
}

type todoService struct {
	repo     TodoRepository
	userRepo user.UserRepository
}

func NewTodoService(repo TodoRepository, userRepo user.UserRepository) TodoService {
	return &todoService{repo: repo, userRepo: userRepo}
}

func (service todoService) CreateTodoList(c *fiber.Ctx) error {
	var dto CreateTodoListRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	userID, _ := strconv.Atoi(c.Locals("userID").(string))
	todoUser, nil := service.userRepo.FindByID(userID)

	_, err := service.repo.SaveTodoList(&dto, todoUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo list created successfully",
	})
}

func (service todoService) ShareTodoList(c *fiber.Ctx) error {
	var dto ShareTodoListRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	user, err := service.userRepo.FindByID(dto.UserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	_, err = service.repo.ShareTodoList(dto.TodoListID, user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo list shared successfully",
	})
}
