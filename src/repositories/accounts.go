package repositories

import (
	"database/sql"

	"github.com/danilo-lopes/simple-bank/src/models"
)

type accountsRepository struct {
	db *sql.DB
}

// NewUsersRepository creates a Users repository
func NewUsersRepository(db *sql.DB) *accountsRepository {
	return &accountsRepository{db}
}

func (repository accountsRepository) CreateAccount(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO users (name, email, cpf) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Email, user.CPF)
	if erro != nil {
		return 0, erro
	}

	userID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	statement, erro = repository.db.Prepare(
		"INSERT INTO accounts (id, active, available_limit) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}

	result, erro = statement.Exec(userID, user.Account.Active, user.Account.AvailableLimit)
	if erro != nil {
		return 0, erro
	}

	return uint64(userID), nil
}

func (repository accountsRepository) GetAccount(userID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"SELECT * FROM users WHERE id = ?",
		userID,
	)
	defer lines.Close()
	if erro != nil {
		return models.User{}, erro
	}

	var user models.User
	for lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CPF,
		); erro != nil {
			return models.User{}, erro
		}
	}

	lines, erro = repository.db.Query(
		"SELECT * FROM accounts WHERE id = ?",
		userID,
	)
	defer lines.Close()
	if erro != nil {
		return models.User{}, erro
	}

	var account models.Account
	var active int
	for lines.Next() {
		if erro = lines.Scan(
			&account.ID,
			&active,
			&account.AvailableLimit,
		); erro != nil {
			return models.User{}, erro
		}
	}
	if active == 0 {
		account.Active = false
	} else {
		account.Active = true
	}

	user.Account = account

	return user, nil
}
