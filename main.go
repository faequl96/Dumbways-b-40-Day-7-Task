package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")

	route.HandleFunc("/project", myProject).Methods("GET")
	route.HandleFunc("/project/{name}", myProjectDetail).Methods("GET")
	route.HandleFunc("/form-project", myProjectForm).Methods("GET")
	route.HandleFunc("/add-project", myProjectDataForm).Methods("POST")

	route.HandleFunc("/contact", contact).Methods(("GET"))

	fmt.Println("Server running at localhost port 8080")

	http.ListenAndServe("localhost:8080", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}

func myProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("views/myProject.html")

	if err == nil {
		tmpl.Execute(w, nil)
	} else {
		w.Write([]byte("Message: "))
		w.Write([]byte(err.Error()))
	}

	// w.WriteHeader(http.StatusOK)
}

func myProjectDetail(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/myProjectDetail.html")

	// id, _ := strconv.Atoi(mux.Vars(r)["name"])
	// fmt.Println(id)

	response := map[string]interface{}{
		"Title":    "Pemrograman Web",
		"Contents": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minus quae similique alias obcaecati quas aut atque voluptatem quibusdam consequatur repudiandae sequi distinctio aliquid magnam, nemo, sapiente quo non rem deserunt quis praesentium vero quasi eum voluptates qui? Unde, similique, alias obcaecati accusamus voluptatum atque necessitatibus asperiores voluptates, quo perferendis perspiciatis!",
		"Id":       "id",
	}

	if err == nil {
		tmpl.Execute(w, response)
	} else {
		panic(err)
	}

}

func myProjectForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/myProjectForm.html")

	if err == nil {
		tmpl.Execute(w, nil)
	} else {
		panic(err)
	}
}

func myProjectDataForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.PostForm.Get("projectName"))
	fmt.Println(r.PostForm.Get("startDate"))
	fmt.Println(r.PostForm.Get("endDate"))
	fmt.Println(r.PostForm.Get("description"))

}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err == nil {
		tmpl.Execute(w, nil)
	} else {
		panic(err)
	}
}
