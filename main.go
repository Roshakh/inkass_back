package main

import (
	"inkass/inkassback/database"
	"inkass/inkassback/env"
	"inkass/inkassback/router"
	"net/http"

	"github.com/gorilla/handlers"
	language "github.com/moemoe89/go-localization"
)

func bootstrap() {
	// initiate the go-localization & bind some config
	cfg := language.New()
	// json file location
	cfg.BindPath("./locales/language.json")
	// default language
	cfg.BindMainLocale("uz")
	var err error
	env.Lang, err = cfg.Init()
	if err != nil {
		panic(err)
	}
}

func init() {

	bootstrap()
}

func main() {
	database.InitialMigration()
	router.CreateRouter()
	http.ListenAndServe(":4000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization", "Token"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(env.Router))
}
