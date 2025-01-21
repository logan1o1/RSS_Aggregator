package main

import "net/http"

func handlerReadiness(resp http.ResponseWriter, req *http.Request) {
	respondWithJSON(resp, 200, struct{}{})
}
