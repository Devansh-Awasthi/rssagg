package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)
func main(){
	fmt.Println("RSS Agregator")
	godotenv.Load(".env")
	 PortVal := os.Getenv("PORT")
	 if(PortVal == ""){
		log.Fatal("port not found")
	 }
	 router := chi.NewRouter()
	 router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string {"http://*","https://*"},
			AllowedMethods: []string {"GET","POST","PUT","DELETE","OPTIONS" },
			AllowedHeaders: []string {"*"},
			ExposedHeaders: []string {"Link"},
			AllowCredentials: false,
			MaxAge: 300,
		}))
		v1Router := chi.NewRouter()
		v1Router.HandleFunc("/ready",handlerReadyness)
		router.Mount("/v1",v1Router)
	 srv := &http.Server{
		Handler: router,
		Addr: ":"+PortVal,
	 }
	 fmt.Println("Port:" , PortVal)
	 err := srv.ListenAndServe()
	 if err != nil {
		log.Fatal(err)
	 }

}