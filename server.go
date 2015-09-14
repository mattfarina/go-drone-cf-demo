package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// message being done this way is to force an actual test case that can be
// tested in CI/CD. See the simple test for more detail.
func message() string {
	return "Hello World"
}

func hello(w http.ResponseWriter, req *http.Request) {
	msg := message()
	fmt.Fprintln(w, msg)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
