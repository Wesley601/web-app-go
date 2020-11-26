package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/wesley601/fundamentos-web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

//Index busca por todos os produtos e renderiza home
func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.All()
	templates.ExecuteTemplate(w, "Index", produtos)
}

// New renderiza página de adicionar produto
func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

// Insert insere um produto no banco e redireciona para home
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

// Delete deleta um produto pelo id re redireciona para home
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeletaProduto(id)
	http.Redirect(w, r, "/", 301)
}

// Edit Busca o produto pelo id e renderiza a página de edição
func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	p := models.Find(id)
	templates.ExecuteTemplate(w, "Edit", p)
}

// Update atualiza um produto no banco e redireciona para home
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		models.Update(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}
