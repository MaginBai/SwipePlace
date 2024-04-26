package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	configFile = "config.json"
)

var db *sql.DB
var config Config

type User struct {
	Name             string
	Age              int
	Courses          string
	Email            string
	Login            string
	PasswordHash     string
	DateRegistration time.Time
}

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func main() {
	var err error
	config, err = LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/register", registerHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.PostFormValue("name")
	ageStr := r.PostFormValue("age")
	courses := r.PostFormValue("courses")
	email := r.PostFormValue("email")
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	// Проверка наличия обязательных полей
	if name == "" || ageStr == "" || courses == "" || email == "" || login == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	passwordHash := password

	user := User{
		Name:             name,
		Age:              age,
		Courses:          courses,
		Email:            email,
		Login:            login,
		PasswordHash:     passwordHash,
		DateRegistration: time.Now(),
	}

	err = createUser(user)
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User created successfully")
}

func createUser(user User) error {
	sqlStatement := `
	INSERT INTO users (name, age, courses, email, login, password_hash, date_registration)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, user.Name, user.Age, user.Courses, user.Email, user.Login, user.PasswordHash, user.DateRegistration).Scan(&id)
	if err != nil {
		return err
	}
	log.Printf("New record ID is %v\n", id)
	return nil
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
