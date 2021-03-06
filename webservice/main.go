package main

import (
	"fmt"
	"net/http"

	"github.com/gilbertek/go-sandbox/webservice/controllers"
)

func main() {
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Tricia",
	// 	LastName:  "McMillan",
	// }
	//
	// fmt.Println(u)
	//
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	//
	// port := 3000
	// _, err := startWebServer(port, 2)
	// fmt.Println(err)
}

func startWebServer(port, numberOfretries int) (int, error) {
	fmt.Println("Starting server...")

	// do something important
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfretries)
	return port, nil
	// errors.New("Something went wrong")
}
