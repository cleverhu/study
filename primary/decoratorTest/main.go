package main

import "net/http"

func decorator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("token") == "" {
			w.Write([]byte("token error"))
		} else {
			h(w, r)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", decorator(index))
	http.ListenAndServe(":8899", nil)

}
