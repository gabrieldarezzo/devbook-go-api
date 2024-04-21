package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
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

	// _, err := usersRepository.db.Exec("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?", "Atualizado 33", "gabrieldarezzo", "darezzo.gabriel@example.com", 1)
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

// FindUserById in DataBase a couple of users math users
func (usersRepository UsersRepository) UpaUpa(userId uint64) (models.User, error) {

	_, err := usersRepository.db.Exec("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?", "Atualizado", "gabrieldarezzo", "darezzo.gabriel@example.com", 1)
	if err != nil {
		log.Fatal(err)
	}

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
func (usersRepository UsersRepository) UpdateUserById(userId uint64, userNewData models.User) error {
	// statement, erro := usersRepository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")

	// if erro != nil {
	// 	return erro
	// }
	// defer statement.Close()

	// // fmt.Print(userNewData.Name)
	// // fmt.Print(userNewData.Nick)
	// // fmt.Print(userNewData.Email)
	// // fmt.Print(userId)
	// // fmt.Print("\n####\n####\n####____\n####\n")
	// // fmt.Print(userId)
	// // fmt.Print("\n####\n####\n####____\n####\n")

	// _, erro = statement.Exec(userNewData.Name, userNewData.Nick, userNewData.Email, userId)
	// if erro != nil {
	// 	fmt.Print("\n####\n####\n####____\n####\n")
	// 	fmt.Print(erro)
	// 	fmt.Print("\n####\n####\n####____\n####\n")
	// 	return erro
	// }

	// return nil

	// Inicia uma transação
	// tx, err := usersRepository.db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	// Em caso de erro, faz rollback da transação
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return
	// 	}
	// 	// Caso contrário, commita a transação
	// 	err = tx.Commit()
	// }()

	// // Prepara a instrução SQL para atualizar o usuário
	// stmt, err := tx.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()

	// // Executa a instrução SQL com os novos dados do usuário
	// // _, err = stmt.Exec(userNewData.Name, userNewData.Nick, userNewData.Email, userId)
	// _, err = stmt.Exec("XXX", "NICKX", "EMAILX", 3)
	// if err != nil {
	// 	return err
	// }

	// return nil

	_, err := usersRepository.db.Exec("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?", "Gabriel Sousa Darezzo", "gabrieldarezzo", "darezzo.gabriel@example.com", 1)
	if err != nil {
		log.Fatal(err)

	}
	return nil

}
