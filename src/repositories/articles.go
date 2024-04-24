package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represent of repository
type ArticlesRepository struct {
	db *sql.DB
}

// NewRepositoryOfUsers create
func NewRepositoryOfArticles(db *sql.DB) *ArticlesRepository {
	return &ArticlesRepository{db}
}

// CreateUser insert a new user in database
func (articlesRepository ArticlesRepository) CreateArticles(article models.Article) (uint64, error) {

	statement, erro := articlesRepository.db.Prepare("INSERT INTO articles (title, content, author_id) VALUES (?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(article.Title, article.Content, article.AuthorId)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}
