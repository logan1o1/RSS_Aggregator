package main

import "net/http"

func handlerError(resp http.ResponseWriter, req *http.Request) {
	respondWithError(resp, 400, "Something went wrong")
}
