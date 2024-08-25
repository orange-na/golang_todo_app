package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type HelloHandler struct {}
type HogeHandler struct {}
type FugaHandler struct {}


func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello world 🍣")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"Yuto8181nmb",
		"localhost",
		"3306",
		"test",
	)

	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	rows, err := connection.Query("select * from students;")
    // rows, err := connection.Query("INSERT INTO students (id, name) VALUES (1, 'Yuto');")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}


	type Student struct {
		ID int `json:"id"`
		Name string `json:"name"`
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

func (h *HogeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello world 🍣")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"root",
		"Yuto8181nmb",
		"localhost",
		"3306",
		"test",
	)

	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	rows, err := connection.Query("select * from students;")
    // rows, err := connection.Query("INSERT INTO students (id, name) VALUES (1, 'Yuto');")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}


	type Student struct {
		ID int `json:"id"`
		Name string `json:"name"`
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


	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	server.ListenAndServe()
}