package controller

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/SpikeIsUp/pissonChat_groupie_tracker/ApiMemeMakerinternal/meme"
	"github.com/SpikeIsUp/pissonChat_groupie_tracker/SQLiteinternal/storage"
)

func Home(w http.ResponseWriter, r *http.Request) {
	memes, err := meme.GetMemes()
	if err != nil {
		http.Error(w, "Erreur API", http.StatusInternalServerError)
		log.Println("Erreur GetMemes:", err)
		return
	}
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	if err := tmpl.Execute(w, memes); err != nil {
		log.Println("Erreur template Home:", err)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	memes, err := meme.GetMemes()
	if err != nil {
		http.Error(w, "Erreur API", http.StatusInternalServerError)
		log.Println("Erreur GetMemes Search:", err)
		return
	}
	var result []meme.Meme
	for _, m := range memes {
		if strings.Contains(strings.ToLower(m.Name), strings.ToLower(query)) {
			result = append(result, m)
		}
	}
	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println("Erreur template Search:", err)
	}
}

func Favorites(w http.ResponseWriter, r *http.Request) {
	rows, err := storage.DB.Query("SELECT id, name, blank FROM favorites")
	if err != nil {
		http.Error(w, "Erreur DB", http.StatusInternalServerError)
		log.Println("Erreur DB Favorites:", err)
		return
	}
	defer rows.Close()

	var memes []meme.Meme
	for rows.Next() {
		var m meme.Meme
		if err := rows.Scan(&m.ID, &m.Name, &m.Blank); err != nil {
			log.Println("Erreur scan Favorites:", err)
			continue
		}
		memes = append(memes, m)
	}
	log.Println("Nombre de favoris récupérés:", len(memes))
	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	if err := tmpl.Execute(w, memes); err != nil {
		log.Println("Erreur template Favorites:", err)
	}
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erreur ParseForm AddFavorite:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	_, err := storage.DB.Exec(
		"INSERT OR IGNORE INTO favorites (id, name, blank) VALUES (?, ?, ?)",
		r.FormValue("id"),
		r.FormValue("name"),
		r.FormValue("blank"),
	)
	if err != nil {
		log.Println("Erreur Insert Favorite:", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Erreur ParseForm RemoveFavorite:", err)
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
		return
	}
	_, err := storage.DB.Exec(
		"DELETE FROM favorites WHERE id = ?",
		r.FormValue("id"),
	)
	if err != nil {
		log.Println("Erreur Remove Favorite:", err)
	}
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/aPropos.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Erreur template About:", err)
	}
}
