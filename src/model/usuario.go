package model

import (
	"api-gestar-bem/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

// Usuario - vai representar um usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

func (Usuario *Usuario) Preparar(etapa string) error {
	if erro := Usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := Usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (Usuario *Usuario) validar(etapa string) error {
	if Usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if Usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if Usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(Usuario.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if etapa == "cadastro" && Usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaHash)
	}
	return nil
}
