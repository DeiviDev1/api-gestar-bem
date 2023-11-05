package controllers

import (
	"api-gestar-bem/src/autentication"
	"api-gestar-bem/src/banco"
	"api-gestar-bem/src/model"
	"api-gestar-bem/src/repositorys"
	"api-gestar-bem/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao cria uma nova publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autentication.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao model.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositorys.NewRepositoryPublicacoes(db)
	publicacao.ID, erro = repository.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

// BuscarPublicacoes traz as publicacoes de um usuario que apareceriam no feed
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao traz uma unica publicacao do banco de dados
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao altera os dados de uma publicacao no banco de dados
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// DeletarPublicacao exclui uma publicacao do banco de dados
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
