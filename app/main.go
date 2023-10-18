package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Embed struct {
	Title   string
	Message string
	Users   map[int]User
	Time    time.Time
}

type User struct {
    ID       int
    Name     string
    Password string
}

const (
    DriverName = "mysql"
    DataSourceName = "root:root@tcp(mysql-container:3306)/golang_db"
)


var usr = make(map[int]User)
var templates = make(map[string]*template.Template)
var dbErr error
var queryErr error
var err error


func main() {
    // docker起動時にmysqlコンテナが立ち上がるよりも先に接続しにいってしまうので、少し時間を置く
    // todo:もう少しスマートな実装にしたい
   time.Sleep(3 * time.Second)
	db, dbErr := sql.Open(DriverName, DataSourceName)
	if dbErr != nil {
	    log.Fatal("error connecting to database:", dbErr)
	}
	defer db.Close()

	rows, queryErr := db.Query("SELECT * FROM users")
    if queryErr != nil {
        log.Fatal("query error:", queryErr)
    }
    defer rows.Close()

    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Password); err != nil {
            log.Fatal(err)
        }
        usr[u.ID] = User{
            ID:       u.ID,
            Name:     u.Name,
            Password: u.Password,
        }
    }

    port := "8080"
    templates["index"] = loadTemplate("index")
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", handleIndex)
    log.Printf("server listening on http://localhost:%s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    temp := Embed{"Hello Golang!", "こんにちは！", usr, time.Now()}
    if err := templates["index"].Execute(w, temp); err != nil {
        log.Print("failed to execute template: %v", err)
    }
}

func loadTemplate(name string) *template.Template {
    t, err := template.ParseFiles(
        "root/"+name+".html",
        "root/template/header.html",
        "root/template/footer.html",
    )
    if err != nil {
        log.Fatalf("template error: %v", err)
    }
    return t
}
