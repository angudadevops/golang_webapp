package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type userinfo struct {
	fname string
	lname string
}

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/users", 302)
}

func main() {
	tmpl := template.Must(template.ParseFiles("users.html"))
	mux := http.NewServeMux()
        //mux.Handler(r)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		if r.Method == "POST" {
			//mux := http.NewServeMux()
			//mux.Handler(r)
			mux.HandleFunc("/users", redirect)
			tmpl1 := template.Must(template.ParseFiles("db.html"))
			if r.Method != http.MethodPost {
				tmpl1.Execute(w, nil)
				return
			}
			//tmpl.Execute(w, struct{ users []user }{users})

			details := userinfo{
				fname: r.FormValue("fname"),
				lname: r.FormValue("lname"),
			}
			db, err := sql.Open("mysql", "root:root@(mysql:3306)/test_db?parseTime=true")
			if err != nil {
				log.Fatal(err)
				return
			}
			if err := db.Ping(); err != nil {
				log.Fatal(err)
				return
			}
			if db != nil {
				fmt.Println("Connected to MySQL DB")
			}

			firstname := details.fname
			lastname := details.lname

			result, err := db.Exec(`INSERT INTO MyUsers (firstname, lastname) VALUES (?, ?)`, firstname, lastname)
			if err != nil {
				log.Fatal(err)
			}

			id, err := result.LastInsertId()
			fmt.Println(id)

			type user struct {
				firstname string
				lastname  string
			}
			//query := `SELECT * FROM MyUsers`
			rows, err := db.Query(`SELECT * FROM MyUsers`)
			if err != nil {
				log.Fatal(err)
				//return
			}
			defer rows.Close()

			var users []user
			for rows.Next() {
				var u user

				err := rows.Scan(&u.firstname, &u.lastname)
				if err != nil {
					log.Fatal(err)
				}
				users = append(users, u)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}

			type TodoPageData struct {
				PageTitle string
				Users     []user
			}
			data := TodoPageData{
				PageTitle: "My SQL DB Users list",
				Users:     users,
			}
			fmt.Printf("%q\n", users)
			tmpl1.Execute(w, data)

			//tmpl.Execute(w, struct{ Success bool }{true})
		}
	})
	http.ListenAndServe(":8080", nil)
}

