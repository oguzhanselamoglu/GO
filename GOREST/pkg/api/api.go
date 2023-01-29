package api

import (
	"fmt"
	"net/http"
)

func Api() {
	fmt.Println("Api")
	http.HandleFunc("/", index)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}
func post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post page"))
}
