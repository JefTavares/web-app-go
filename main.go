package main

import (
	"net/http"

	"github.com/jeftavares/web-app-go/routes"
)

func main() {

	routes.CarregaRotas() //carrega as rotas inicialmente

	//?Sobe o servidor http na porta 8000
	http.ListenAndServe(":8000", nil)
}
