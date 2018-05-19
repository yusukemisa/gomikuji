package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	rand.Seed(time.Now().UnixNano())
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, omikuji())
}

func omikuji() string {
	castNum := rand.Intn(6) + 1
	var result string
	switch castNum {
	case 1:
		result = "凶"
	case 2:
		result = "吉"
	case 3:
		result = "吉"
	case 4:
		result = "中吉"
	case 5:
		result = "中吉"
	case 6:
		result = "大吉"
	}
	return fmt.Sprintf("おみくじ結果：%v,さいのめ：%v\n", result, castNum)
}
