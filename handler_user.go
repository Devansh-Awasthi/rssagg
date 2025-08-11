package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/google/uuid"
)
func (apiCfg *apiConfig) handlerUser(w http.ResponseWriter , r *http.Request){
	type parameter struct{
		Name string `json:"name"` 
	}
	decoder := json.NewDecoder(r.Body)
	param := parameter{}
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w,400,fmt.Sprintf("Error on Parsing:%s",err))
		return
	}
	user,err:=apiCfg.db.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
	 Name   :  param.Name,
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		responseWithErr(w,400,fmt.Sprintf("Error on Creating User:%s",err))
		return
	}
	

	responseWithJSON(w,200,user)
}