package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Advert struct {
	Id      int
	Text    string
	Caption string
	Cost    int
	Urlfoto string
	Date    time.Time
}

var db *sql.DB

func SelectAll() string {
	rows, err := db.Query("select id, text, caption, cost, Urlfoto, date from avito.Adverts ")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	Adverts := []Advert{}

	for rows.Next() {
		p := Advert{}
		err := rows.Scan(&p.Id, &p.Text, &p.Caption, &p.Cost, &p.Urlfoto, &p.Date)
		if err != nil {
			fmt.Println(err)
			continue
		}

		Adverts = append(Adverts, p)
		fmt.Println(p.Date)
	}

	b, err := json.Marshal(&Adverts)
	println(string(b))
	return string(b)
}

func Select1(id int) string {
	advert := Advert{}
	_ = db.QueryRow("select id, text, caption, cost, Urlfoto, date from avito.Adverts  where id = ?", id).Scan(&advert.Id, &advert.Text, &advert.Caption, &advert.Cost, &advert.Urlfoto, &advert.Date)
	b, _ := json.Marshal(&advert)
	println(string(b))
	return string(b)
}

func Add1(advert Advert) {
	_, err := db.Exec("insert into avito.Adverts (Text, Caption, Cost, Urlfoto, Date) values (?, ?, ?, ?, ?)",
		advert.Text, advert.Caption, advert.Cost, advert.Urlfoto, time.Now())
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:4605421QWqw@/avito?parseTime=true")

	if err != nil {
		log.Println(err)
	}

}

func main() {

	SelectAll()

	//advert := Advert{0,"1", "2", 3, "4", time.Now()}
	//Add1(advert)

	//defer db.Close()
	//Select1(2)

	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9]+}", HandlerId)
	router.HandleFunc("/", HandlerKoren)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":80", nil)

}
func HandlerKoren(w http.ResponseWriter, r *http.Request) {
	otvet := SelectAll()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, otvet)
}

func HandlerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["id"]
	id, _ := strconv.Atoi(sid)

	otvet := Select1(id)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, otvet)
}
