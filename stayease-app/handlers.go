package stayease

import "net/http"

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
