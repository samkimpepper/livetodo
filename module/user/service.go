package user

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"todo/middleware/auth"
)

type UserService interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(c *fiber.Ctx) error {
	var dto RegisterRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	if err := hashPassword(&dto.Password); err != nil {
		return err
	}

	_, err := s.repo.Save(&dto)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Successfully registered",
	})
}

func (s *userService) Login(c *fiber.Ctx) error {
	var dto LoginRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	user, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return err
	}

	accessToken, err := auth.GenerateAccessToken(strconv.Itoa(user.ID), user.Email)
	if err != nil {
		return err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.Email)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		Expires:  refreshToken.ExpAt,
		HTTPOnly: true,
		Path:     "/",
		SameSite: "None",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":          "Successfully logged in",
		"access_token": accessToken.Token,
	})
}

func hashPassword(pw *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*pw), 14)
	*pw = string(bytes)
	return err
}

func (s *userService) GetUser(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Locals("userID").(string))
	user, err := s.repo.FindByID(userId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
