package main

import (
	"Auto_Reload/controllers"
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	port := PORT

	go controllers.AutoReload()
	http.HandleFunc("/", controllers.ReloadWeb)
	fmt.Println("Auto Reload app is listening on port", port)
	http.ListenAndServe(port, nil)
}
