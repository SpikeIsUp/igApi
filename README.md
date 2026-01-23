# 🎉 Meme Universe

**Meme Universe** est une application web développée en **Go** permettant de découvrir des memes, de les rechercher et de les enregistrer en favoris grâce à une base de données **SQLite**.

Le projet repose sur une architecture claire (MVC léger), utilise des templates HTML, du CSS pour le style et une API interne pour récupérer les memes.

---

## 🚀 Fonctionnalités

- 🖼️ Affichage de memes depuis une API interne  
- 🔍 Recherche de memes par nom  
- ⭐ Ajout de memes aux favoris  
- 🗑️ Suppression des favoris  
- 💾 Sauvegarde persistante des favoris avec SQLite  
- 🎨 Interface web stylée et responsive  
- 📄 Page “À propos”

---

## 🧱 Architecture du projet

pissonChat_groupie_tracker/
│
├── main.go
├── go.mod
├── go.sum
│
├── controller/
│ └── controller.go
│
├── router/
│ └── router.go
│
├── SQLiteinternal/
│ └── storage/
│ └── storage.go
│
├── ApiMemeMakerinternal/
│ └── meme/
│
├── template/
│ ├── home.html
│ ├── recherche.html
│ ├── favoris.html
│ └── aPropos.html
│
├── assets/
│ └── css/
│ ├── home.css
│ ├── recherche.css
│ └── aPropos.css
│
└── database.db

---

## 🛠️ Technologies utilisées

- **Go 1.24**
- **net/http**
- **html/template**
- **SQLite** (`modernc.org/sqlite`)
- **HTML5 / CSS3**
- Architecture MVC simplifiée

---

## ⚙️ Installation et lancement

### 1️⃣ Cloner le dépôt
```bash
git clone https://github.com/SpikeIsUp/pissonChat_groupie_tracker.git
cd pissonChat_groupie_tracker
