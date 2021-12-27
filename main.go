package main

import (
	"fmt"
	"inkass/inkassback/env"

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
	fmt.Println(env.Lang.Lookup("uz", "Good morning"))
	fmt.Println(env.Lang.Lookup("cr", "Good morning"))
	fmt.Println(env.Lang.Lookup("ru", "Good morning"))
}
