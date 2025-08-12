package main

import (
	"fmt"
	"net/http"

	"github.com/Devansh-Awasthi/rssagg/internal/auth"
	"github.com/Devansh-Awasthi/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ApiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithErr(w, 403, fmt.Sprintf("Error on getting User:%s", err))
			return
		}
		user, err := apiCfg.db.GetUserByAPIKey(r.Context(), ApiKey)
		if err != nil {
			responseWithErr(w, 403, fmt.Sprintf("Error on getting user by api key:%s", err))
			return
		}
		h(w,r,user)
	}
	

}
