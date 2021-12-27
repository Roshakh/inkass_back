package env

import (
	"github.com/gorilla/mux"
	"github.com/moemoe89/go-localization"

)

var (
	Router *mux.Router
	Lang *language.Config 
)
