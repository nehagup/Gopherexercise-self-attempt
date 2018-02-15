package main

import (
	"net/http"
	"io"
)

func main(){
http.HandleFunc("/", foo)
http.HandleFunc("/some" , some)
http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"Hello world! You're back?")
}
func some(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r , "https://www.nehas.ml", http.StatusTemporaryRedirect)
}