package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//CarregarTemplates insere os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//ExecutarTemplate renderiza uma pagina html na tela
func ExecutarTemplate(rw http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(rw, template, dados)
}
