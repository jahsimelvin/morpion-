package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Game struct {
	Board         [][]string // Tableau du jeu
	CurrentPlayer string     // Joueur actuel
	Winner        string     // Vainqueur du jeu (le cas échéant)
	IsGameOver    bool       // Indique si le jeu est terminé
}

// Fonction pour initialiser un nouveau jeu
func NewGame() *Game {
	// Créer un tableau vide de 3x3
	board := make([][]string, 3)
	for i := range board {
		board[i] = make([]string, 3)
		for j := range board[i] {
			board[i][j] = " " // Initialiser chaque case avec un espace vide
		}
	}
	// Créer un nouvel objet Game avec le tableau initialisé et d'autres valeurs par défaut
	return &Game{
		Board:         board,
		CurrentPlayer: "X", // Commencer par le joueur X
		Winner:        "",
		IsGameOver:    false,
	}
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "game.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func winHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "win.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func looseHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "loose.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func equalityHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "equality.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed for random number generation

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/win", winHandler)
	http.HandleFunc("/loose", looseHandler)
	http.HandleFunc("/equality", equalityHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("picture"))))

	log.Println("Server started on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
