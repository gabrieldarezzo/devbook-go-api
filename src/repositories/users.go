package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// FindUser in DataBase a couple of users math users
func (usersRepository UsersRepository) FindUser(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, erro := usersRepository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick)

	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// FindUserById in DataBase a couple of users math users
func (usersRepository UsersRepository) FindUserById(userId uint64) (models.User, error) {

	// _, err := usersRepository.db.Exec("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?", "Atualizado 44", "gabrieldarezzo", "darezzo.gabriel@example.com", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	row, erro := usersRepository.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE id = ?", userId)

	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if erro = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// UpdateUserById Update user in DataBase passing his ID
func (usersRepository UsersRepository) UpdateUserById(userId uint64, userToUpdate models.User) error {

	_, erro := usersRepository.db.Exec("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
		userToUpdate.Name,
		userToUpdate.Nick,
		userToUpdate.Email,
		userId,
	)
	if erro != nil {
		return erro
	}

	return nil
}

// UpdateUserById Update user in DataBase passing his ID
func (usersRepository UsersRepository) DeleteUser(userId uint64) error {

	_, erro := usersRepository.db.Exec("DELETE FROM users WHERE id = ?", userId)
	if erro != nil {
		return erro
	}

	return nil
}