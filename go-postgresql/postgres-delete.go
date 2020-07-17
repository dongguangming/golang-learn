package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
        //_ "github.com/bmizerany/pq"
        
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:12345678@192.168.8.200:5432/douyin?sslmode=disable"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func main() {

	// Create DB pool
	//db, err := sql.Open("postgres", "host=192.168.8.200 port=5432 user=postgres password=12345678 dbname=douyin sslmode=disable")
	db, err := sql.Open("postgres",DB_DSN)
        if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()


        var user User = User{ID:4,Email:"dong@qq.com",Password:"abcdedf120"}
         _, err = db.Exec("DELETE FROM  users  where id=$1", user.ID)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("delete  ok!!!")

	// Create an empty user and make the sql query (using $1 for the parameter)
	var myUser User
	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	err = db.QueryRow(userSql, 4).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	} else {

	fmt.Printf("hello email: %s, password: %s, welcome back!\n", myUser.Email,myUser.Password)
    
        }
}
