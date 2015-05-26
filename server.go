package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world!")
}

var t = template.Must(template.ParseFiles("index.html"))

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //処理の最後にBodyを閉じる

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		// ファイル名を{id}.txt
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンス201
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		// パラメータを取得
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// person作成
		person := Person{
			ID:   id,
			Name: string(b),
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}
