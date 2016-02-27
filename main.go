package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	cookieName = "dsdscookie"
)

var (
	count int = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(cookieName)
		if err != nil {
			panic(err)
		}
		s := c.Value
		if s != "" {
			log.Println("Got cookie:", s)
		} else {
			cookieValue := strconv.Itoa(count)
			count = count + 1
			log.Println("No cookie, generate new:", cookieValue)
			cookie := &http.Cookie{
				Name:  cookieName,
				Value: cookieValue,
			}
			http.SetCookie(w, cookie)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
