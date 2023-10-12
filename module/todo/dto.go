package todo

type CreateTodoListRequest struct {
	Title string `json:"title"`
}

type ShareTodoListRequest struct {
	TodoListID int `json:"todo_list_id"`
	UserID     int `json:"user_id"`
}
