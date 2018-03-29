package gofontwoff_test

import (
	"net/http"

	"github.com/shurcooL/gofontwoff"
	"github.com/shurcooL/httpgzip"
)

func Example() {
	fontsHandler := http.FileServer(gofontwoff.Assets)
	http.Handle("/assets/fonts/", http.StripPrefix("/assets/fonts", fontsHandler))
}

func Example_httpGzip() {
	fontsHandler := httpgzip.FileServer(gofontwoff.Assets, httpgzip.FileServerOptions{ServeError: httpgzip.Detailed})
	http.Handle("/assets/fonts/", http.StripPrefix("/assets/fonts", fontsHandler))
}
