package models

import (
	"tfalc/GolangStudies/WebApp/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaComBanco()

	selectDeProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()

	insertData, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectaComBanco()

	deleteData, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Produto {
	db := db.ConectaComBanco()

	produtoDB, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic("Produto não encontrado!")
	}

	produto := Produto{}

	for produtoDB.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDB.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}
