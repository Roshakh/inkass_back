package env

import (
	"html/template"

	"github.com/gorilla/mux"
	language "github.com/moemoe89/go-localization"
)

var (
	Router    *mux.Router
	Lang      *language.Config
	Templates *template.Template
)
