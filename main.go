package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

type gomikuji struct {
	t time.Time
}

type omikujiResult struct {
	Fortune  string `json:"fortune"`
	ImageURL string `json:"imageURL"`
}

func main() {
	http.HandleFunc("/", handle)
	rand.Seed(time.Now().UnixNano())
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	g := &gomikuji{
		t: time.Now(),
	}
	result := g.omikuji()

	jsonString, err := result2JsonString(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, jsonString)
}

//おみくじ結果をJOSN文字列に変換
func result2JsonString(o *omikujiResult) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(o); err != nil {
		return "", err
	}
	return buf.String(), nil
}

//おみくじロジック
func (g *gomikuji) omikuji() *omikujiResult {
	castNum := g.cast()
	var result, imgURL string
	switch castNum {
	case 1:
		result = "凶"
		imgURL = ""
	case 2, 3:
		result = "吉"
		imgURL = ""
	case 4, 5:
		result = "中吉"
		imgURL = ""
	case 6:
		result = "大吉"
		imgURL = ""
	}
	return &omikujiResult{
		Fortune:  result,
		ImageURL: imgURL,
	}

}

func (g *gomikuji) cast() int {
	//日にちを取得
	//_, m, d := g.t.Date()
	//if m == time.January &&
	return rand.Intn(6) + 1
}
