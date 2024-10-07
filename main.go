package main

import (
	"database/sql"
	"github/dhqanh/bosu-project/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiCOnfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't not cnont to dtb")
	}

	storeDB := database.New(conn)
	// if err != nil {
	// 	log.Fatal("Can't connect to database: ", err)
	// }

	apiCfg := apiCOnfig{
		DB: storeDB,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthy", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/user", apiCfg.handlerUsersCreate)
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Server starting in port %v", portStr)

	log.Fatal(server.ListenAndServe())

}
