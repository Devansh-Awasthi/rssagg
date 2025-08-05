package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func responseWithErr(w http.ResponseWriter,code int,msg string){
	if code>499 {
		log.Println("Error Occured",msg)
	}
	type err struct{
		Error string `json:"errors"`
	}
	responseWithJSON(w,code,err{
		Error : msg,
	})
}
func responseWithJSON(w http.ResponseWriter,code int,payload interface{}){
  data,err := json.Marshal(payload)
  if err != nil {
	log.Fatal("Marshal Failed",err)
	w.WriteHeader(500)
	return
  }
  w.Header().Add("Content-Type","application/json")
  w.WriteHeader(code)
  w.Write(data)

}