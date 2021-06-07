package repositories

import (
	"api/src/modules"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepositories(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user modules.User) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into users(name, nick, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil

}

func (repo users) Find(nameOrNick string) ([]modules.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	results, erro := repo.db.Query(
		"select id, name, nick, email, createAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer results.Close()

	var users []modules.User

	for results.Next() {
		var user modules.User

		if erro = results.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo users) FindID(ID uint64) (modules.User, error) {
	results, erro := repo.db.Query(
		"select id, name, nick, email, createAt from users where id = ?",
		ID,
	)
	if erro != nil {
		return modules.User{}, erro
	}
	defer results.Close()

	var user modules.User

	if results.Next() {
		if erro = results.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); erro != nil {
			return modules.User{}, erro
		}
	}

	return user, nil

}

func (repo users) UserUpdate(ID uint64, user modules.User) error {
	statement, erro := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}

	_, erro = statement.Exec(user.Name, user.Nick, user.Email, ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repo users) UserDelete(ID uint64) error {
	statement, erro := repo.db.Prepare(
		"delete from users where id = ?",
	)
	if erro != nil {
		return erro
	}

	_, erro = statement.Exec(ID)
	if erro != nil {
		return erro
	}

	return nil
}
