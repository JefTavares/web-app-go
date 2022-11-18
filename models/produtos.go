package models

import "github.com/jeftavares/web-app-go/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	/* Gera o lista na unha
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, manga curta", Preco: 39.90, Quantidade: 5},
		{"Tenis", "estilo bota", 89.90, 33}, {"Fone", "grande", 59, 2},{"Novo", "legal", 1.99, 1},
	}*/

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		/* Declara variaveis para ser utilizadas no Scan e atribuir a um novo produto */
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade) //faz a leitura da linha e atribui as variaveis
		if err != nil {
			panic(err.Error())
		}
		//Seta uma nova struct de produto (declarada lá em cima)
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) //adiciona na lista de produtos
	}
	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	//Prepara a inserção e valida o script insert
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade) //efetiva o insert

}
