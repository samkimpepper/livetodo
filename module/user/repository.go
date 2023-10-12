package user

import (
	"context"
	"todo/ent"
	"todo/ent/user"
)

type UserRepository interface {
	Save(dto *RegisterRequest) (*ent.User, error)
	FindByEmail(email string) (*ent.User, error)
	FindByID(id int) (*ent.User, error)
}

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) Save(dto *RegisterRequest) (*ent.User, error) {

	savedUser, err := repo.db.User.Create().
		SetEmail(dto.Email).
		SetPassword(dto.Password).
		SetUsername(dto.Username).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (repo userRepository) FindByEmail(email string) (*ent.User, error) {
	selectedUser, err := repo.db.User.Query().
		Where(user.Email(email)).
		First(context.TODO())
	if err != nil {
		return nil, err
	}

	return selectedUser, nil
}

func (repo userRepository) FindByID(id int) (*ent.User, error) {
	selectedUser, err := repo.db.User.Query().
		Where(user.ID(id)).
		First(context.TODO())
	if err != nil {
		return nil, err
	}

	return selectedUser, nil
}
