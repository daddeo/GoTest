package main

import (
	"fmt"
	"net/http"
)

/*
	testing this:
	1. from command line, AND in test/webtest use:
		go build
	2. run with:
		./webtest
	3. go to browser and hit:
		http://localhost:8080/hello?name=Jason

	should see in the browser:
	Hello, Jason

	Stop command line with CTRL-C (VS code) to stop program
*/

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	// nil handler tells Go to use the default HTTP handler built into the HTTP package
	http.ListenAndServe(":8080", nil)
}
