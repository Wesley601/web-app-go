package models

import (
	"database/sql"

	"github.com/wesley601/fundamentos-web/db"
)

// Produto modelo para tabela de produtos
type Produto struct {
	ID              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

// Find Busca por um produto pelo id
func Find(id string) Produto {
	database := db.DbConnect()
	defer database.Close()
	rows, err := database.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	p := Produto{}

	for rows.Next() {
		p = scanProduto(rows)
	}
	return p
}

// All busca por produtos na tabela de produtos
func All() []Produto {
	database := db.DbConnect()
	defer database.Close()
	selectProdutos, err := database.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	defer selectProdutos.Close()

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		p = scanProduto(selectProdutos)
		produtos = append(produtos, p)
	}
	return produtos
}

// CriarNovoProduto insere um novo produto no banco de dados
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	database := db.DbConnect()
	defer database.Close()
	insereProduto, err := database.Prepare(
		"insert into produtos (nome, descricao, preco, quantidade) values($1, $2, $3, $4)",
	)
	if err != nil {
		panic(err.Error())
	}

	insereProduto.Exec(nome, descricao, preco, quantidade)
}

// DeletaProduto deleta um produto através do id
func DeletaProduto(id string) {
	database := db.DbConnect()
	defer database.Close()
	deletarProduto, err := database.Prepare(
		"delete from produtos where id=$1",
	)
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
}

// Update insere um novo produto no banco de dados
func Update(id int, nome, descricao string, preco float64, quantidade int) {
	database := db.DbConnect()
	defer database.Close()
	insereProduto, err := database.Prepare(
		"update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5",
	)
	if err != nil {
		panic(err.Error())
	}

	insereProduto.Exec(nome, descricao, preco, quantidade, id)
}

func scanProduto(selectProdutos *sql.Rows) Produto {
	var id, quantidade int
	var nome, descricao string
	var preco float64

	err := selectProdutos.Scan(
		&id,
		&nome,
		&descricao,
		&preco,
		&quantidade,
	)

	if err != nil {
		panic(err.Error())
	}

	p := Produto{id, nome, descricao, preco, quantidade}

	return p
}
