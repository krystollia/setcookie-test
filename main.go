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
		toGenerate := false
		if err != nil {
			toGenerate = true
			log.Println("Get cookie err:", err)
		} else {
			if s := c.Value; s == "" {
				toGenerate = true
			} else {
				log.Println("Got cookie:", s)
			}
		}

		if toGenerate {
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
