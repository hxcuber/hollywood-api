package helloWorld

import (
	"fmt"
	"net/http"
)

func GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}
