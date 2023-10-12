package todo

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"log"
	"strconv"
	"sync"
	"todo/ent"
	"todo/middleware/auth"
)

var todoLists = make(map[int]map[*websocket.Conn]struct{})
var mu sync.Mutex

// type: "notify" "saved" "updated"
type Message struct {
	Type       string `json:"type"`
	Content    string `json:"content"`
	TodoItemID int    `json:"todo_item_id"`
}

func connectTodoList(c *websocket.Conn, repo TodoRepository) {
	var todolist_id string = c.Params("id")
	todolistID, _ := strconv.Atoi(todolist_id)
	todolist, _ := repo.FindTodoListByID(todolistID)
	log.Println("WEBSOCKET connect to todolist: ", todolist_id)

	tokenString := c.Query("token")
	if tokenString == "" {
		log.Println("WEBSOCKET no token")
		c.Close()
		return
	}
	_, err := auth.VerifyToken(tokenString)
	if err != nil {
		log.Println("WEBSOCKET invalid token")
		c.Close()
		return
	}

	mu.Lock()
	if todoLists[todolistID] == nil {
		todoLists[todolistID] = make(map[*websocket.Conn]struct{})
	}
	todoLists[todolistID][c] = struct{}{}
	mu.Unlock()
	broadcast(todolistID, "notify", []byte("New user connected"), -1)

	for {
		var (
			msg []byte
			err error
		)
		if _, msg, err = c.ReadMessage(); err != nil {
			log.Println("WEBSOCKET read error:", err)
			break
		}
		log.Printf("WEBSOCKET recv: %s", msg)
		todoItemID, messageType := readMessage(repo, msg, todolist)

		broadcast(todolistID, messageType, msg, todoItemID)
	}

	mu.Lock()
	delete(todoLists[todolistID], c)
	mu.Unlock()
	broadcast(todolistID, "notify", []byte("User disconnected"), -1)
}

func broadcast(todolistID int, messageType string, message []byte, todoItemID int) {
	mu.Lock()
	defer mu.Unlock()
	msg := Message{
		Type:       messageType,
		Content:    string(message),
		TodoItemID: todoItemID,
	}

	if clients, ok := todoLists[todolistID]; ok {
		for client := range clients {
			sendMessage(client, msg)
		}
	}
}

func sendMessage(c *websocket.Conn, msg Message) {
	err := c.WriteJSON(msg)
	if err != nil {
		log.Println("WEBSOCKET write:", err)
	}
}

func readMessage(repo TodoRepository, msg []byte, todolist *ent.TodoList) (int, string) {
	var data map[string]interface{}
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Println("WEBSOCKET unmarshal:", err)
		return -1, ""
	}

	messageType, ok := data["type"].(string)
	if !ok {
		log.Println("WEBSOCKET invalid type")
		return -1, ""
	}

	switch messageType {
	case "saved":
		content, ok := data["content"].(string)
		if !ok {
			log.Println("WEBSOCKET invalid content")
			return -1, ""
		}
		todoItem, _ := repo.SaveTodoItem(content, todolist)
		return todoItem.ID, "saved"
	case "updated":
		todoitemID, ok := data["todo_item_id"].(int)
		if !ok {
			log.Println("WEBSOCKET invalid todo_item_id")
			return -1, ""
		}
		content, ok := data["content"].(string)
		if !ok {
			log.Println("WEBSOCKET invalid content")
			return -1, ""
		}

		todoitemID, _ = repo.UpdateTodoItem(todoitemID, content)
		return todoitemID, "updated"
	case "deleted":
		todoitemID, ok := data["todo_item_id"].(int)
		if !ok {
			log.Println("WEBSOCKET invalid todo_item_id")
			return -1, ""
		}

		repo.DeleteTodoItem(todoitemID)
		return todoitemID, "deleted"
	default:
		log.Println("WEBSOCKET invalid type")
		return -1, ""
	}
}
