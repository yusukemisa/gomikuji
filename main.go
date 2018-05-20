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

// const (
// 	daikichi
// )

type gomikuji struct {
	t time.Time
}

type omikujiResult struct {
	Text string      `json:"text"`
	Img  attachments `json:"attachments"`
	//ImageURL string `json:"imageURL"`
}

type attachments []attachment

type attachment struct {
	Title    string `json:"title"`
	Color    string `json:"color"`
	ImageURL string `json:"image_url"`
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
	var result, image string
	switch castNum {
	case 1:
		result = "凶"
		image = "omikuji_kyou.png"
	case 2, 3:
		result = "吉"
		image = "omikuji_kichi.png"
	case 4, 5:
		result = "中吉"
		image = "omikuji_chuukichi.png"
	case 6:
		result = "大吉"
		image = "omikuji_daikichi.png"
	}
	a := attachments{
		{
			ImageURL: "https://raw.githubusercontent.com/yusukemisa/gomikuji/master/" + image,
			Title:    "あなたの運勢",
			Color:    "#bce2e8",
		},
	}
	return &omikujiResult{
		Text: result,
		Img:  a,
	}
}

func (g *gomikuji) cast() int {
	//日にちを取得
	//_, m, d := g.t.Date()
	//if m == time.January &&
	return rand.Intn(6) + 1
}

/*
{
  "text": "I found a MONSTER! ",
  "channel": "#incoming-test",
"color": "#2eb886",
  "icon_emoji": ":baymax:"
}
*/
