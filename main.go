package main

import (
	"fmt"
	"net/http"
)

// TODO urlDatabase
var urlDatabase = map[string]string{
	"google": "https://www.google.com",
	"github": "https://github.com/adaurideveloper-sudo",
}

// TODO http.HandleFunc
func main() {

	http.HandleFunc("/", redirect)

	fmt.Println("Servidor rodando em htpp://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

// TODO func redirect
func redirect(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Path[1:]

	fmt.Printf("Buscamos pelo codigo: '%s'\n", code)

	if code == "" {
		fmt.Fprint(w, "Bem-vindo! Digite um codigo na URL, ex: /google")
		return
	}

	originalURL, exists := urlDatabase[code]

	if exists {
		fmt.Printf("Sucesso! Redirecionado para: %s\n, code")
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		fmt.Printf("erro: O codigo '%s' nao existe no mapa \n", code)
		http.Error(w, "Link nao encontrado! verifique se digitou corretamente", http.StatusNotFound)

	}

}
