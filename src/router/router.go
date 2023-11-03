package router

import (
	"api-gestar-bem/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar - vai retornar as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
