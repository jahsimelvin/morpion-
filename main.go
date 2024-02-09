package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var index = template.Must(template.ParseFiles("html/index.html"))
var game = template.Must(template.ParseFiles("html/game.html"))
var win = template.Must(template.ParseFiles("html/index.html"))
var loose = template.Must(template.ParseFiles("html/index.html"))
var equality = template.Must(template.ParseFiles("html/index.html"))
var rules = template.Must(template.ParseFiles("html/index.html"))


const port = ":3000"

func main(){
	http.HandleFunc("/",Index)

	http.HandleFunc("/game",Game)

	http.HandleFunc("/win", Win)

	http.HandleFunc("/loose", Loose)

	http.HandleFunc("/equality", Equality)

	http.HandleFunc("/rules", Rules)

	http.HandleFunc("/win", Win)

	fmt.Println("Serveur en cours d'execution sur http://localhost:8080")
	http.ListenAndServe(port, nil)
}

func Index (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/index")
}

func Game (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/game")
}

func Win (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/win")
}

func Loose (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/loose")
}

func Equality (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/equality")
}

func Rules (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "html/rules")
}

func renderTemplate(w http.ResponseWriter, tmpl string) { //Parse le fichier html et envoi les informations au client
	t, err := template.ParseFiles("./" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}