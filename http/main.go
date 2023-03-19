package main

import "net/http"

func main() {
	httpDemo()
}

func httpDemo() {
	http.Handle("/index.html", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello word"))
	}))

	http.ListenAndServe(":8088", nil)
}
