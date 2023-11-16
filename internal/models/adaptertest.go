package models

import "net/http"

type HandlerFuncAdapter func(http.ResponseWriter, *http.Request)

func (h HandlerFuncAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}
