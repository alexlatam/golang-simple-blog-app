package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseGlob("views/**/*"))

func DBConection() (conecction *sql.DB) {
	Driver := "mysql"
	User := "root"
	Password := ""
	Host := "localhost"
	Port := "3306"
	Database := "go-simple-blog"

	conection, err := sql.Open(Driver, User+":"+Password+"@tcp("+Host+":"+Port+")/"+Database)

	if err != nil {
		panic(err.Error())
	}
	return conection
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/store", Store)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	log.Println("Listen on the port 8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))

}

type Post struct {
	Id        int
	Title     string
	Content   string
	Author    string
	CreatedAt string
}

// Functions
func Home(w http.ResponseWriter, r *http.Request) {

	DBConection := DBConection()
	posts, err := DBConection.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}

	post := Post{}
	postArrays := []Post{}

	for posts.Next() {
		var id int
		var title, content, author, created_at string
		err = posts.Scan(&id, &title, &content, &author, &created_at)
		if err != nil {
			panic(err.Error())
		}
		post.Id = id
		post.Title = title
		post.Content = content[:450] + "..."
		post.Author = author
		post.CreatedAt = FormatDate(created_at)

		postArrays = append(postArrays, post)
	}
	templates.ExecuteTemplate(w, "home", postArrays)
}
func Show(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	DBConection := DBConection()
	result, err := DBConection.Query("SELECT * FROM posts WHERE id = " + id + " LIMIT 1; ")
	if err != nil {
		panic(err.Error())
	}
	post := Post{}

	for result.Next() {
		var id int
		var title, content, author, created_at string

		err = result.Scan(&id, &title, &content, &author, &created_at)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = content
		post.Author = author
		post.CreatedAt = created_at

	}
	defer result.Close()

	templates.ExecuteTemplate(w, "show", post)
}
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}
func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		author := r.FormValue("author")

		if title == "" || content == "" || author == "" {
			http.Redirect(w, r, "/create", 301)
		}

		DBConection := DBConection()
		insert, err := DBConection.Prepare("INSERT INTO posts (title, content, author) VALUES (?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(title, content, author)
		defer insert.Close()

	}
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	DBConection := DBConection()
	result, err := DBConection.Query("SELECT * FROM posts WHERE id = " + id + " LIMIT 1; ")
	if err != nil {
		panic(err.Error())
	}
	post := Post{}

	for result.Next() {
		var id int
		var title, content, author, created_at string

		err = result.Scan(&id, &title, &content, &author, &created_at)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = content
		post.Author = author
		post.CreatedAt = created_at

	}
	defer result.Close()

	templates.ExecuteTemplate(w, "edit", post)
}
func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Println(id)

	if r.Method == "POST" {

		title := r.FormValue("title")
		content := r.FormValue("content")
		author := r.FormValue("author")

		if title == "" || content == "" || author == "" {
			http.Redirect(w, r, "/edit?id="+id, 301)
		}
		DBConection := DBConection()
		update, err := DBConection.Prepare("UPDATE posts SET title = ?, content = ?, author = ? WHERE id = ?;")
		if err != nil {
			panic(err.Error())
		}
		update.Exec(title, content, author, id)
		defer update.Close()

	}
	http.Redirect(w, r, "/show?id="+id, 301)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	DBConection := DBConection()
	delete, err := DBConection.Prepare("DELETE FROM posts WHERE id = ?;")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer delete.Close()

	http.Redirect(w, r, "/", 301)
}

// Function to format date on this format Month day, year
func FormatDate(date string) string {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Format("January 2, 2006")
}
