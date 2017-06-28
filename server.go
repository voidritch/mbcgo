package main

import (
        "fmt"
	"http"
	"mbcgo/config"
)

func main() {
        // fs := http.FileServer(http.Dir("static"))
	root := func (w http.ResponseWriter, r *http.Request) {
    		fmt.Fprintf(w, "You have reached %s!", r.URL.Path[1:])
	}
        http.Handle("/", root)

        log.Println("Listening...")
        http.ListenAndServe(":8080", nil)
}
