package main

import (
	"fmt"
	"net/http"

	"github.com/benytto888Z/bookingN/pkg/handlers"
)


const PORT = ":8080"

func main()  {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about", handlers.About)

	
	fmt.Println("Server start on port "+PORT)
	http.ListenAndServe(PORT, nil)
}