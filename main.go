package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/wesley601/fundamentos-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
