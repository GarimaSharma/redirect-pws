package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/pivotal-cf/greenhouse-private/wiki", 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
