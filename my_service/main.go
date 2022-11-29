package main

import (
	"my_service/solver"
	"net/http"
)

func main() {

	err := http.ListenAndServe("localhost:3000", solver.NewService())
	if err != nil {
		panic(err)
	}

}
