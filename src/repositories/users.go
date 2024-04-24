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

// FindUserByEmail Find a User in Database using Email as criteria
func (usersRepository UsersRepository) FindUserByEmail(email string) (models.User, error) {

	row, erro := usersRepository.db.Query("SELECT id, email, password FROM users WHERE email = ? ", email)

	if erro != nil {
		return models.User{}, erro
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if erro = row.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

// FallowNewUser a user
func (usersRepository UsersRepository) FallowNewUser(userId uint64, followIdUser uint64) error {

	statement, erro := usersRepository.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES (?,?)")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userId, followIdUser); erro != nil {
		return erro
	}

	return nil
}

// UnFollowUser a user
func (usersRepository UsersRepository) UnFollowUser(userId uint64, followIdUser uint64) error {

	statement, erro := usersRepository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userId, followIdUser); erro != nil {
		return erro
	}

	return nil
}

// FollowersOfByUserId Get all followers of a user
func (usersRepository UsersRepository) FollowersOfByUserId(userId uint64) ([]models.User, error) {
	rows, erro := usersRepository.db.Query(`
		select 
			uf.id,
			uf.name,
			uf.nick,
			uf.created_at
		from 
		followers 
		inner join users uf on (
			followers.follower_id = uf.id
		)
		where followers.user_id = ?
	`, userId)
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
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}
	return users, nil
}

// FollowersOfByUserId Get who a user is fallowing
func (usersRepository UsersRepository) GetAllFollowingUsersOfUserId(userId uint64) ([]models.User, error) {
	rows, erro := usersRepository.db.Query(`
		select 
			uf.id,
			uf.name,
			uf.nick,
			uf.created_at
		from 
		followers 
		inner join users uf on (
			followers.user_id = uf.id
		)
		where followers.follower_id = ?
	`, userId)
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
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}
	return users, nil
}

// UpdatePasswordUserById Update user in DataBase passing his ID
func (usersRepository UsersRepository) UpdatePasswordUserById(userId uint64, userToUpdate models.User) error {

	_, erro := usersRepository.db.Exec("UPDATE users SET password = ? WHERE id = ?",
		userToUpdate.Password,
		userId,
	)

	if erro != nil {
		return erro
	}

	return nil
}

func (usersRepository UsersRepository) GetPaswordByIdUser(userId uint64) (string, error) {

	row, erro := usersRepository.db.Query("SELECT password FROM users WHERE id = ?", userId)

	if erro != nil {
		return "", erro
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if erro = row.Scan(
			&user.Password,
		); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}
