package controller

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// lire les template
func RenderTemplate(w http.ResponseWriter, tmpl string) error {
    path := filepath.Join("template", tmpl)
    log.Println("Rendu:", path)

    t, err := template.ParseFiles(path)
    if err != nil {
        log.Println("Erreur ParseFiles:", err)
        http.Error(w, "Erreur serveur (template introuvable)", http.StatusInternalServerError)
        return err
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    if err := t.Execute(w, nil); err != nil {
        log.Println("Erreur Execute:", err)
        http.Error(w, "Erreur serveur (render)", http.StatusInternalServerError)
        return err
    }

    return nil
}

// dire a renderTemplate de lire les bonne template en fontion des routes
func APropos(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /aPropos")
    if err := RenderTemplate(w, "aPropos.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Categories(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /categories")
    if err := RenderTemplate(w, "categories.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Collection(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /collection")
    if err := RenderTemplate(w, "collection.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Favoris(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /favoris")
    if err := RenderTemplate(w, "favoris.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Home(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /")
    if err := RenderTemplate(w, "home.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Recherche(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /recherche")
    if err := RenderTemplate(w, "recherche.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}

func Ressources(w http.ResponseWriter, r *http.Request) {
    log.Println("GET /ressources")
    if err := RenderTemplate(w, "ressources.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}
