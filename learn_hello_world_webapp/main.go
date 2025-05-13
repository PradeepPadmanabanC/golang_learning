package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is a about page and 2+2 is %d", sum))

}

func AddValues(x, y int) int {
	return x + y
}

// test add
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
