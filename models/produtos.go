package models

import "github.com/felipesbreve/db"

type Produto struct {
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func buscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}