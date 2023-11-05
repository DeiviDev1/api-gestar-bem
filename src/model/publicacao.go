package model

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorid,omitempty"`
	AutorNick string    `json:"autornick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaem,omitempty"`
}

func (p *Publicacao) Preparar() error {
	if erro := p.validar(); erro != nil {
		return erro
	}
	p.formatar()
	return nil
}

func (p *Publicacao) validar() error {
	if p.Titulo == "" {
		return errors.New("O titulo é obrigatório e não pode estar em branco")
	}
	if p.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}
