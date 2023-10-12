package todo

import (
	"context"
	"todo/ent"
	"todo/ent/todoitem"
	"todo/ent/todolist"
)

type TodoRepository interface {
	SaveTodoList(dto *CreateTodoListRequest, user *ent.User) (*ent.TodoList, error)
	SaveTodoItem(title string, todoList *ent.TodoList) (*ent.TodoItem, error)
	FindTodoListByID(id int) (*ent.TodoList, error)
	ShareTodoList(id int, user *ent.User) (int, error)
	FindTodoItemByID(id int) (*ent.TodoItem, error)
	UpdateTodoItem(id int, title string) (int, error)
	DeleteTodoItem(id int) (int, error)
}

type todoRepository struct {
	db *ent.Client
}

func NewTodoRepository(db *ent.Client) TodoRepository {
	return todoRepository{db: db}
}

func (repo todoRepository) SaveTodoList(dto *CreateTodoListRequest, user *ent.User) (*ent.TodoList, error) {

	savedTodoList, err := repo.db.TodoList.Create().
		SetTitle(dto.Title).
		AddUsers(user).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return savedTodoList, nil
}

func (repo todoRepository) SaveTodoItem(title string, todoList *ent.TodoList) (*ent.TodoItem, error) {

	savedTodoItem, err := repo.db.TodoItem.Create().
		SetTitle(title).
		SetList(todoList).
		Save(context.TODO())
	if err != nil {
		return nil, err
	}

	return savedTodoItem, nil
}

func (repo todoRepository) FindTodoListByID(id int) (*ent.TodoList, error) {
	todoList, err := repo.db.TodoList.Query().
		Where(todolist.ID(id)).
		First(context.TODO())
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (repo todoRepository) ShareTodoList(id int, user *ent.User) (int, error) {
	todoList, err := repo.db.TodoList.Update().
		Where(todolist.ID(id)).
		AddUsers(user).
		Save(context.TODO())
	if err != nil {
		return -1, err
	}

	return todoList, nil
}

func (repo todoRepository) FindTodoItemByID(id int) (*ent.TodoItem, error) {
	todoItem, err := repo.db.TodoItem.Query().
		Where(todoitem.ID(id)).
		First(context.TODO())
	if err != nil {
		return nil, err
	}

	return todoItem, nil
}

func (repo todoRepository) UpdateTodoItem(id int, title string) (int, error) {
	todoItem, err := repo.db.TodoItem.Update().
		Where(todoitem.ID(id)).
		SetTitle(title).
		Save(context.TODO())
	if err != nil {
		return -1, err
	}

	return todoItem, nil
}

func (repo todoRepository) DeleteTodoItem(id int) (int, error) {
	todoItem, err := repo.db.TodoItem.Delete().
		Where(todoitem.ID(id)).
		Exec(context.TODO())
	if err != nil {
		return -1, err
	}

	return todoItem, nil
}
