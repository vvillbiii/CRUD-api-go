package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){
	Connection()


	//init router
	r := httprouter.New()
	http.ListenAndServe("localhost:5000", r)

}

