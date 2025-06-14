package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>bleh</title>
			</head>
			<body>
				yooo my n word
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
