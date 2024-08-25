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

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

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

type HelloHandler struct {}
type HogeHandler struct {}
type FugaHandler struct {}

type GetTasksHandler struct {}
type AddTaskHandler struct {}
type EditTaskHandler struct {}
type DeleteTaskHandler struct {}

type SignUpHandler struct {}
type SignInHandler struct {}
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

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	rows, err := connection.Query("select * from user;")
    // rows, err := connection.Query("INSERT INTO students (id, name) VALUES (1, 'Yuto');")
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

	_, err = connection.Exec("UPDATE task SET title = ? WHERE id = ?", task.Title, task.Id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *DeleteTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
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

	_, err = connection.Exec("DELETE FROM task WHERE id = ?", task.Id)
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

	id := uuid.New().String()
	name := r.FormValue("name")
	password := r.FormValue("password")

	_, err = connection.Exec("INSERT INTO user (id, name, password) VALUES (?, ?, ?)", id, name, password)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	connection, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	rows, err := connection.Query("select * from students;")
    // rows, err := connection.Query("INSERT INTO students (id, name) VALUES (1, 'Yuto');")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	students := []Student{}
	
	for rows.Next() {
		student := Student{}
		rows.Scan(&student.ID, &student.Name)
		students = append(students, student)
	}
	
	rows.Close()

	response, _ := json.Marshal(students)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}



func (h *FugaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func main() {

	// HelloHandler 型の変数を宣言
	handler := HelloHandler{}
	hogeHandler := HogeHandler{}
	fugaHandler := FugaHandler{}
	getTasksHandler := GetTasksHandler{}

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 許可するオリジンを指定。"*"は全てのオリジンを許可
	    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッドを指定
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}, // 許可するHTTPヘッダを指定
	})
	
	// ハンドラにCORSの設定を適用
	http.Handle("/", c.Handler(&handler))
	http.Handle("/hoge", c.Handler(&hogeHandler))
	http.Handle("/fuga", c.Handler(&fugaHandler))

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasksHandler.ServeHTTP(w, r)
		case http.MethodPost:
			(&AddTaskHandler{}).ServeHTTP(w, r) 
		case http.MethodPut:
			(&EditTaskHandler{}).ServeHTTP(w, r)
		case http.MethodDelete:
			(&DeleteTaskHandler{}).ServeHTTP(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasksHandler.ServeHTTP(w, r)
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