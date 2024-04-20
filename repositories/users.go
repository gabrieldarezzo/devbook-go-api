package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represent of repository
type UsersRepository struct {
	db *sql.DB
}

// NewRepositoryOfUsers create
func NewRepositoryOfUsers(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

// CreateUser insert a new user in database
func (usersRepository UsersRepository) CreateUser(user models.User) (uint64, error) {

	// fmt.Print("\n\n####Pass here!\n")
	statement, erro := usersRepository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}
