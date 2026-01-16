package controller

import (
	"html/template"
	"net/http"

	"pissonChat_groupie_tracker/internal/meme"
	"pissonChat_groupie_tracker/internal/storage"
)

func Home(w http.ResponseWriter, r *http.Request) {
	memes, _ := meme.GetMemes()
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	tmpl.Execute(w, memes)
}

func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	memes, _ := meme.GetMemes()

	var result []meme.Meme
	for _, m := range memes {
		if query == "" || contains(m.Name, query) {
			result = append(result, m)
		}
	}

	tmpl := template.Must(template.ParseFiles("template/recherche.html"))
	tmpl.Execute(w, result)
}

func Favorites(w http.ResponseWriter, r *http.Request) {
	rows, _ := storage.DB.Query("SELECT id, name, url FROM favorites")
	defer rows.Close()

	var memes []meme.Meme
	for rows.Next() {
		var m meme.Meme
		rows.Scan(&m.ID, &m.Name, &m.URL)
		memes = append(memes, m)
	}

	tmpl := template.Must(template.ParseFiles("template/favoris.html"))
	tmpl.Execute(w, memes)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	storage.DB.Exec(
		"INSERT OR IGNORE INTO favorites VALUES (?, ?, ?)",
		r.FormValue("id"),
		r.FormValue("name"),
		r.FormValue("url"),
	)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	storage.DB.Exec("DELETE FROM favorites WHERE id = ?", r.FormValue("id"))
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/aPropos.html"))
	tmpl.Execute(w, nil)
}

func contains(a, b string) bool {
	return len(a) >= len(b) && (a == b || len(b) == 0 || (len(a) > len(b) && contains(a[1:], b)))
}
