package repositorys

import (
	"api-gestar-bem/src/model"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NewRepositoryPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (r Publicacoes) Criar(publicacao model.Publicacao) (uint64, error) {
	statament, erro := r.db.Prepare("insert into publicacao (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statament.Close()

	resultado, erro := statament.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

func (r Publicacoes) BuscarPorID(publicacaoID uint64) (model.Publicacao, error) {
	linha, erro := r.db.Query(`
		select p.*, u.nick from 
	    publicacao p inner join usuarios u 
	    on u.id = p.autor_id where p.id = ?`,
		publicacaoID)
	if erro != nil {
		return model.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao model.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return model.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (r Publicacoes) Buscar(usuarioID uint64) ([]model.Publicacao, error) {
	linha, erro := r.db.Query(`
	 select distinct p.*, u.nick from publicacao p
	 inner join usuarios u on u.id = p.autor_id 
	 inner join seguidores s on p.autor_id = s.usuario_id
	 where u.id = ? or s.seguidor_id = ?
	 order by 1 desc`,
		usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linha.Close()

	var publicacoes []model.Publicacao
	for linha.Next() {
		var publicacao model.Publicacao
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (r Publicacoes) Atualizar(publicacaoID uint64, publicacao model.Publicacao) error {
	statament, erro := r.db.Prepare("update publicacao set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (r Publicacoes) Deletar(publicacaoID uint64) error {
	statament, erro := r.db.Prepare("delete from publicacao where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (r Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]model.Publicacao, error) {
	linha, erro := r.db.Query(`
		select p.*, u.nick from publicacao p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`,
		usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linha.Close()

	var publicacoes []model.Publicacao
	for linha.Next() {
		var publicacao model.Publicacao
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)

	}
	return publicacoes, erro
}

func (r Publicacoes) Curtir(publicacaoID uint64) error {
	statament, erro := r.db.Prepare("update publicacao set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statament.Close()

	if _, erro = statament.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

func (r Publicacoes) Descurtir(publicacaoID uint64) error {
	statament, erro := r.db.Prepare(`
	    update publicacao set curtidas =
		case when curtidas > 0 then curtidas - 1
	    else 0 end where id = ?`)

	if erro != nil {
		return erro
	}

	if _, erro = statament.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}
