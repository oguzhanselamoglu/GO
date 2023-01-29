package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Product struct {
	Id   int
	Name string
}

var products []Product

func dataGet() {
	prd1 := Product{
		Id:   1,
		Name: "PRD 1",
	}
	prd2 := Product{
		Id:   2,
		Name: "PRD 2",
	}
	products = append(products, prd1, prd2)
}

func Api() {
	fmt.Println("Api")
	r := mux.NewRouter()
	dataGet()
	r.HandleFunc("/", index)
	r.HandleFunc("/post", post)
	r.HandleFunc("/post/{category}/{id}", post)
	r.HandleFunc("/products", getProducts)
	http.ListenAndServe(":8080", r)

}
func getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	j, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "appliscation/json")
	fmt.Fprintf(w, string(j))
	//w.Write(j)

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
