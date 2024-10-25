package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Mod
// Go offers several tools like go build (production ready code), run, mod
// go mod init github.com/bhavyansh001/mymodules
// Semantic versioning is done, seen in the mod file
// indirect goes away when you use it in your project
// go env is another tool, can see where the packages go in dir
// go mod tidy
// go build .
// go mod verify
// go list all			: list all modules
// go list -m all
// go list -m -versions github.com/gorilla/mux		: all available versions which you can choose from to use
// go mod why github.com/gorilla/mux		: why am I dependent on this package
// go mod graph
// go mod edit -go 1.15		, 	go mod edit -module
// go mod vendor		: node_modules
// go run -mod=vendor main.go		: runs from vendor folder that has all the required modules

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
    r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
	// log.Fatal() is for error handling
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang, hello!</h1>"))
}