package repositorys

import (
	"api-gestar-bem/src/model"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NewRepositoryUsuarios - vai criar um novo repositório de usuários
func NewRepositoryUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar - vai criar um usuário no banco de dados
func (r usuarios) Criar(usuario model.Usuario) (uint64, error) {
	statement, erro := r.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

func (r usuarios) Buscar(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := r.db.Query(
		"select id, nome, nick, email, criadoem from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var usuarios []model.Usuario
	for linhas.Next() {
		var usuario model.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (r usuarios) BuscarPorID(ID uint64) (model.Usuario, error) {
	linhas, erro := r.db.Query(
		"select id, nome, nick, email, criadoem from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario model.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (r usuarios) Atualizar(ID uint64, usuario model.Usuario) error {
	statement, erro := r.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}

func (r usuarios) Deletar(ID uint64) error {
	statement, erro := r.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (r usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	linha, erro := r.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linha.Close()

	var usuario model.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (r usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := r.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (r usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := r.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

func (r usuarios) BuscarSeguidores(usuarioID uint64) ([]model.Usuario, error) {
	linhas, erro := r.db.Query(`
      select u.id, u.nome, u.nick, u.email, u.criadoem
      from usuarios u inner join seguidores s on u.id = s.seguidor_id
      where s.usuario_id = ?
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []model.Usuario
	for linhas.Next() {
		var usuario model.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (r usuarios) BuscarSeguindo(usuarioID uint64) ([]model.Usuario, error) {
	linhas, erro := r.db.Query(`
	  select u.id, u.nome, u.nick, u.email, u.criadoem
	  from usuarios u inner join seguidores s on u.id = s.usuario_id
	  where s.seguidor_id = ?`,
		usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []model.Usuario
	for linhas.Next() {
		var usuario model.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, erro
}
