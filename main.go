package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles("tmpl/index.html"))
var game = template.Must(template.ParseFiles("tmpl/game.html"))
var win = template.Must(template.ParseFiles("tmpl/index.html"))
var loose = template.Must(template.ParseFiles("tmpl/index.html"))
var equality = template.Must(template.ParseFiles("tmpl/index.html"))
var rules = template.Must(template.ParseFiles("tmpl/index.html"))
var pseudo = template.Must(template.ParseFiles("tmpl/pseudo.html"))

type tictactoe struct{
	pseudo		string
	win			bool
	equality	bool
	endgame		bool
}

const port = ":3000"

func main(){
	http.HandleFunc("/",Index)

	http.HandleFunc("/game",Game)

	http.HandleFunc("/win", Win)

	http.HandleFunc("/loose", Loose)

	http.HandleFunc("/equality", Equality)

	http.HandleFunc("/rules", Rules)

	http.HandleFunc("/pseudo", Pseudo)

	fmt.Println("Serveur en cours d'execution sur http://localhost:8080")
	http.ListenAndServe(port, nil)
}

func Index (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/index")
}

func Game (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/game")
}

func Pseudo(w http.ResponseWriter, r *http.Request) {
	// var data string
	// pseudo := r.FormValue("pseudo")

	// if len(pseudo) > 0 {
	// 	// Mettre à jour la variable pseudo

	// 	data.Pseudo = pseudo

	// 	http.Redirect(w, r, "/skin", http.StatusSeeOther)
	// }

	renderTemplate(w, "tmpl/pseudo")

}

func Win (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/win")
}

func Loose (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/loose")
}

func Equality (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/equality")
}

func Rules (w http.ResponseWriter, r*http.Request){
	renderTemplate(w, "tmpl/rules")
}

func renderTemplate(w http.ResponseWriter, tmpl string) { //Parse le fichier html et envoi les informations au client
	t, err := template.ParseFiles("./" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

