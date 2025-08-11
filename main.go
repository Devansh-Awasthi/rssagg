package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
type apiConfig struct{
	db *database.Queries
}
func main(){
	fmt.Println("RSS Agregator")
	godotenv.Load(".env")
	 PortVal := os.Getenv("PORT")

	 if(PortVal == ""){
		log.Fatal("port not found")
	 }
	 db_URL := os.Getenv("DB_URL")
	 if(db_URL == ""){
		log.Fatal("db_URL not found")
	 }
	 connection,err := sql.Open("postgres",db_URL)
	 if err != nil {
		 log.Fatal("cant connect to db",err)
		}
		queries:= database.New(connection)
	 apiCfg := apiConfig{
        db:queries,
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
		v1Router.Get("/ready",handlerReadyness)
		v1Router.Get("/errHandle",handlerErr)
		v1Router.Post("/users",apiCfg.handlerUser)
		router.Mount("/v1",v1Router)
	 srv := &http.Server{
		Handler: router,
		Addr: ":"+PortVal,
	 }
	 fmt.Println("Port:" , PortVal)
	 err1 := srv.ListenAndServe()
	 if err1 != nil {
		log.Fatal(err1)
	 }

}