package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w,"hello world")
	})
	http.ListenAndServe(":8080",nil)
}
