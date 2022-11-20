package routes

import (
	"net/http"

	"github.com/jeftavares/web-app-go/controllers"
)

func CarregaRotas() {
	/*Cria uma rota*/
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)       //exibe a tela de cadastro novo produto
	http.HandleFunc("/insert", controllers.Insert) //rota que faz o insert ne um novo produto
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)     //exibe e carrega a tela de edição
	http.HandleFunc("/update", controllers.Update) //atualiza no banco
}
