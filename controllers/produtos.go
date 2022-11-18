package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jeftavares/web-app-go/models"
)

// cria os templates (as páginas html)
// must encapsula todos os templates e devolve dois retornos.
// ParseGlob tras todos com *.html
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	// passa vazio sem param temp.ExecuteTemplate(w, "Index", nil)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome") //pega lá da tag <input type="text" name="nome" class="form-control">
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//conversão de string para float
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		//converte para inteiro
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301) //volta para a página principal
}
