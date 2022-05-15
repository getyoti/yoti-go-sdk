package main

import (
	"html/template"
	"log"
	"net/http"
)

func errorPage(w http.ResponseWriter, r *http.Request) {
	templateVars := map[string]interface{}{
		"yotiError": r.Context().Value(contextKey("yotiError")).(string),
	}
	log.Printf("%s", templateVars["yotiError"])
	t, err := template.ParseFiles("error.html")
	if err != nil {
		panic(errParsingTheTemplate + err.Error())
	}

	err = t.Execute(w, templateVars)
	if err != nil {
		panic(errApplyingTheParsedTemplate + err.Error())
	}

}
