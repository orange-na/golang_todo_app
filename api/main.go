package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/rs/cors"
)



type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type Task struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Title string `json:"title"`
}


type GetTasksHandler struct {}
type AddTaskHandler struct {}
type EditTaskHandler struct {}
type DeleteTaskHandler struct {}

type SignUpHandler struct {}
type SignInHandler struct {}
type GetUsersHandler struct {}
type GetMeHandler struct {}


func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"Yuto8181nmb",
		"localhost",
		"3306",
		"todos",
	)
	return sql.Open("mysql", dsn)
}


func (h *GetUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	rows, err := connection.Query("select * from user;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}


	users := []User{}
	
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}
	
	rows.Close()

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *GetTasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	rows, err := connection.Query("select * from task;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	tasks := []Task{}

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.Id, &task.UserId, &task.Title)
		tasks = append(tasks, task)
	}

	rows.Close()

	response, _ := json.Marshal(tasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *AddTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    connection, err := connectDB()
    if err != nil {
        panic(err)
    }
    defer connection.Close()

    task.Id = uuid.New().String()

    _, err = connection.Exec("INSERT INTO task (id, userId, title) VALUES (?, ?, ?)", task.Id, task.UserId, task.Title)
    if err != nil {
        fmt.Println("Error executing query:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}


func (h *EditTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskId := r.URL.Path[len("/tasks/"):]
	if taskId == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	_, err = connection.Exec("UPDATE task SET title = ? WHERE id = ?", task.Title, taskId)
	if err != nil {
		fmt.Println("Error executing query:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *DeleteTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    taskId := r.URL.Path[len("/tasks/"):]
    if taskId == "" {
        http.Error(w, "Task ID is required", http.StatusBadRequest)
        return
    }

    connection, err := connectDB()
    if err != nil {
        panic(err)
    }
    defer connection.Close()

    _, err = connection.Exec("DELETE FROM task WHERE id = ?", taskId)
    if err != nil {
        fmt.Println("Error executing query:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}


func (h *SignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

    var user User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	user.Id = uuid.New().String()

	_, err = connection.Exec("INSERT INTO user (id, name, password) VALUES (?, ?, ?)", user.Id, user.Name, user.Password)
	if err != nil {
		fmt.Println("Error executing query:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func (h *SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

    var user User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	rows, err := connection.Query("select * from user where name = ? and password = ?", user.Name, user.Password)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	users := []User{}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}

	rows.Close()

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}





func main() {

	// HelloHandler 型の変数を宣言
	// handler := HelloHandler{}
	// hogeHandler := HogeHandler{}
	// fugaHandler := FugaHandler{}
	// getTasksHandler := GetTasksHandler{}

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 許可するオリジンを指定。"*"は全てのオリジンを許可
	    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッドを指定
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}, // 許可するHTTPヘッダを指定
	})

    http.Handle("/tasks", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            (&GetTasksHandler{}).ServeHTTP(w, r)
        case http.MethodPost:
            (&AddTaskHandler{}).ServeHTTP(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })))

    http.Handle("/tasks/", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            (&EditTaskHandler{}).ServeHTTP(w, r)
        case http.MethodDelete:
            (&DeleteTaskHandler{}).ServeHTTP(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })))

	http.Handle("/sign-up", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		(&SignUpHandler{}).ServeHTTP(w, r)
	})))

	http.Handle("/sign-in", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		(&SignInHandler{}).ServeHTTP(w, r)
	})))

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			(&GetUsersHandler{}).ServeHTTP(w, r)
		case http.MethodPost:
			// TODO: Implement AddTaskHandler
		case http.MethodPut:
			// TODO: Implement EditTaskHandler
		case http.MethodDelete:
			// TODO: Implement DeleteTaskHandler
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})


	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	server.ListenAndServe()
}