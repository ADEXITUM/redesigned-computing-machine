package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/index.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/articles/articles")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		//установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`,`anons`,`full_text`) VALUES(`%s`, `%s`, `%s`)", title, anons, full_text))

		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/create/", create)
	http.HandleFunc("/", index)
	http.HandleFunc("/save_article/", save_article)
	http.ListenAndServe(":3608", nil)
}

func main() {
	handleFunc()
}
