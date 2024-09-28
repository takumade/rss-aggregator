package main

import (
	"net/http"

	"github.com/rss-aggregator/internal/database"
)


type authedHandler func(http.ResponseWriter, *http.Request, database.User)


func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}