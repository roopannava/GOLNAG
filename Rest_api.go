package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	ID        string `json:"id"`
	FIRSTNAME string `json:"FIRSTNAME"`
}
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getHello(context *gin.Context) {
	fmt.Printf("Welcome to the HomePage!")
}

func getUsers(context *gin.Context) {
	fmt.Printf("Getting users")
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "user1:pw1234@tcp(127.0.0.1:3306)/testdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("select ID,FIRSTNAME from ns_user")
	fmt.Println("Results")
	fmt.Println(results)
	for results.Next() {

		var tag User
		// for each row, scan the result into our tag composite object
		fmt.Println("ID", results.Scan(&tag.FIRSTNAME))
		err = results.Scan(&tag.ID, &tag.FIRSTNAME)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.FIRSTNAME)
	}
	defer db.Close()
}

func main() {

	// defer the close till after the main function has finished
	// executing

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/hello", getUsers)

	router.Run("localhost:9090")

}
