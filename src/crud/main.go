package main

import (
        "database/sql"
        "encoding/json"
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux" 
        "github.com/gorilla/handlers"
        _ "github.com/go-sql-driver/mysql"
)

type User struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
}

var db *sql.DB

func init() {
        var err error
        db, err = sql.Open("mysql", "root:arathi@tcp(localhost:3306)/hari")
        if err != nil {
                log.Fatal(err)
        }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
        enableCors(&w)
        rows, err := db.Query("SELECT id, name FROM users")
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
                var u User
                if err := rows.Scan(&u.ID, &u.Name); err != nil {
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                }
                users = append(users, u)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
        enableCors(&w)
        vars := mux.Vars(r)
        id := vars["id"]

        var u User
        err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name)
        if err != nil {
                if err == sql.ErrNoRows {
                        http.Error(w, "User not found", http.StatusNotFound)
                        return
                }
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
      // Handle CORS preflighted request sent by browser.
        enableCors(&w)
        if (*r).Method == "OPTIONS" {
                return
        }
        var u User
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        result, err := db.Exec("INSERT INTO users (name) VALUES (?)", u.Name)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        lastID, err := result.LastInsertId()
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        u.ID = int(lastID)

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
        enableCors(&w)
        if (*r).Method == "OPTIONS" {
                return
        }
        vars := mux.Vars(r)
        id := vars["id"]

        var u User
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        _, err = db.Exec("UPDATE users SET name = ? WHERE id = ?", u.Name, id)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
        enableCors(&w)
        if (*r).Method == "OPTIONS" {
                return
        }
        vars := mux.Vars(r)
        id := vars["id"]

        _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusNoContent)
}

func enableCors(w *http.ResponseWriter) {
        (*w).Header().Set("Access-Control-Allow-Origin", "*")
        (*w).Header().Set("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
        // We need to allow the Authorization header to be sent to the backend.
        (*w).Header().Set("Access-Control-Allow-Headers", "*")
        (*w).Header().Set("Access-Control-Max-Age", "86400")
}

func main() {
        router := mux.NewRouter()

        router.HandleFunc("/users", getUsers).Methods("GET")
        router.HandleFunc("/users/{id}", getUser).Methods("GET")
        router.HandleFunc("/users", createUser).Methods("POST")
        router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
        router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

        corsObj := handlers.AllowedOrigins([]string{"*"}) // Replace with your allowed origins
        corsObj = append(corsObj, handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}))
        corsObj = append(corsObj, handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}))
        handler := handlers.CORS(corsObj)(router)

        fmt.Println("Server listening on :8080")
        log.Fatal(http.ListenAndServe(":8080", handler))
}