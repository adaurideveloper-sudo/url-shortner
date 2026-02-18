package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

var (
	urlDatabase = make(map[string]string)
	dbMutex     = sync.RWMutex{}
)

func gerarCodigo(n int) string {
	const letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codigo := make([]byte, n)
	for i := range codigo {
		codigo[i] = letras[rand.Intn(len(letras))]
	}
	return string(codigo)
}

func main() {
	http.HandleFunc("/encurtar", func(w http.ResponseWriter, r *http.Request) {
		urlLonga := r.URL.Query().Get("url")
		if urlLonga == "" {
			fmt.Fprintf(w, "Erro: passe uma url. Ex: /encurtar?url=https://google.com")
			return
		}

		custom := r.URL.Query().Get("custom")
		var codigoFinal string

		if custom != "" {
			dbMutex.RLock()
			_, exists := urlDatabase[custom]
			dbMutex.RUnlock()
			if exists {
				http.Error(w, "Código customizado já em uso!", http.StatusBadRequest)
				return
			}
			codigoFinal = custom
		} else {
			for {
				codigoFinal = gerarCodigo(5)
				dbMutex.RLock()
				_, exists := urlDatabase[codigoFinal]
				dbMutex.RUnlock()
				if !exists {
					break
				}
			}
		}

		dbMutex.Lock()
		urlDatabase[codigoFinal] = urlLonga
		dbMutex.Unlock()

		fmt.Fprintf(w, "Link criado com sucesso!\n")
		fmt.Fprintf(w, "Acesse: http://localhost:8080/%s", codigoFinal)
	})

	http.HandleFunc("/", redirect)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	if code == "" {
		fmt.Fprint(w, "Bem-vindo ao encurtador!")
		return
	}

	dbMutex.RLock()
	originalURL, exists := urlDatabase[code]
	dbMutex.RUnlock()

	if exists {
		fmt.Printf("Sucesso! Redirecionado para: %s\n", code)
		http.Redirect(w, r, originalURL, http.StatusFound)
	} else {
		fmt.Printf("erro: O codigo '%s' não existe no mapa \n", code)
		http.Error(w, "Link não encontrado! verifique se digitou corretamente", http.StatusNotFound)
	}
}