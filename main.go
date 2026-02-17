package main

import (
	"fmt"
	"net/http"
)

// Nosso "banco de dados" temporario da memoria
var urlDatabase = map[string]string{
	"google": "https://www.google.com",
	"github": "https://github.com/adaurideveloper-sudo",
}

func main() {
	//Rota para encurtar
	http.HandleFunc("/", redirect)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func redirect(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path traz algo como "/google"
	// Usamos [1:] para tirar a primeira barra e so ficar so comn "google"
	code := r.URL.Path[1:]

	fmt.Printf("Buscamos pelo codigo: '%s'\n", code)

	if code == "" {
		fmt.Fprint(w, "Bem-vindo! Digite um codigo na URL, ex: /google")
		return
	}

	//Busca no mapa
	originalURL, exists := urlDatabase[code]

	if exists {
		// Se existir, mandamos o usuario para la (Status 302: Found)
		fmt.Printf("Sucesso! Redirecionado para: %s\n, code")
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		//Se nao existir, avisamos
		fmt.Printf("erro: O codigo '%s' nao existe no mapa \n", code)
		http.Error(w, "Link nao encontrado! verifique se digitou corretamente", http.StatusNotFound)

	}

}
