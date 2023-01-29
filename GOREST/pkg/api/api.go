package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Api() {
	fmt.Println("Api")
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/post", post)
	r.HandleFunc("/post/{category}/{id}", post)
	http.ListenAndServe(":8080", r)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}
func post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := vars["category"]

	a := r.Method

	w.Write([]byte("post page, post id:" + id + " category: " + category + " method : " + a))
}
