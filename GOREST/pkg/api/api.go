package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Id   int
	Name string
}
type Option struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Question struct {
	Key         string   `json:"key"`
	Label       string   `json:"label"`
	Required    bool     `json:"required"`
	Order       int      `json:"order"`
	Value       string   `json:"value"`
	ControlType string   `json:"controlType"`
	Type        string   `json:"type"`
	Options     []Option `json:"options"`
}

var products []Product
var questions []Question

func CreateQuestions() {
	q := Question{
		Key:         "name",
		Label:       "Kullanıcı Adı",
		Required:    true,
		Order:       0,
		Value:       "oguzhan",
		ControlType: "textbox",
		Type:        "",
		Options:     nil,
	}
	q1 := Question{
		Key:         "password",
		Label:       "Şifre",
		Required:    true,
		Order:       0,
		Value:       "",
		ControlType: "textbox",
		Type:        "",
		Options:     nil,
	}
	q2 := Question{
		Key:         "mail",
		Label:       "email adresi",
		Required:    false,
		Order:       0,
		Value:       "",
		ControlType: "textbox",
		Type:        "email",
		Options:     nil,
	}

	q3 := Question{
		Key:         "properties",
		Label:       "şirket",
		Required:    false,
		Order:       0,
		Value:       "",
		ControlType: "selectbox",
		Type:        "email",
		Options: []Option{
			{
				Key:   "1",
				Value: "Deneme",
			},
			{
				Key:   "2",
				Value: "Deneme 1",
			},
			{
				Key:   "3",
				Value: "Deneme 2",
			},
		},
	}
	questions = append(questions, q, q1, q2, q3)
}

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
	CreateQuestions()
	r.HandleFunc("/", index)
	r.HandleFunc("/post", post)
	r.HandleFunc("/post/{category}/{id}", post)
	r.HandleFunc("/products", getProducts)
	r.HandleFunc("/user", user)
	r.HandleFunc("/GetQuestions", getQuestions)
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
func getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println(r.URL.Path)
	j, err := json.Marshal(questions)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "appliscation/json")
	fmt.Fprintf(w, string(j))
	//w.Write(j)

}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func user(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	log.Println(r.FormValue("username"))
	log.Println(r.FormValue("password"))

	fmt.Fprintf(w, "success")
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
