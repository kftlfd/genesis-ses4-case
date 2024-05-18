package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var curRate = "39.00"

type RateObj struct {
	Id       uint   `json:"id"`
	Ask      string `json:"ask"`
	Desk     string `json:"desk"`
	AskSum   string `json:"askSum"`
	BidSum   string `json:"bidSum"`
	Currency string `json:"currency"`
	AscCount string `json:"ascCount"`
	BidCount string `json:"bidCount"`
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal([]RateObj{
		{Ask: curRate, Currency: "usd"},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(res)
}

func updateRate() {
	newRate := 3800 + rand.Intn(201)
	curRate = formatRate(newRate)
}

func formatRate(rate int) string {
	return fmt.Sprintf("%d.%02d", rate/100, rate%100)
}

func main() {
	app := http.NewServeMux()

	app.Handle("/", &homeHandler{})

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			updateRate()
		}
	}()

	if err := http.ListenAndServe(":8000", app); err != nil {
		log.Fatal(err)
	}
}
