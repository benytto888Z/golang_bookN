package handlers

import (
	"net/http"

	"github.com/benytto888Z/bookingN/pkg/renders"
)



func Home(w http.ResponseWriter, r *http.Request){
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request){
renders.RenderTemplate(w, "about.page.tmpl")
}