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

// CreateArticles insert a new article in database
func (repository ArticlesRepository) CreateArticles(article models.Article) (uint64, error) {

	statement, erro := repository.db.Prepare("INSERT INTO articles (title, content, author_id) VALUES (?,?,?)")

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

// FindArticle
func (repository ArticlesRepository) FindArticle(articleId uint64) (models.Article, error) {

	row, erro := repository.db.Query(`
	SELECT 
		articles.id,
		title,
		content,
		likes,
		author_id,
		articles.created_at,
		users.nick
	FROM articles 
	INNER JOIN users ON (
		articles.author_id = users.id
	)
	WHERE 
		articles.id = ?
	`, articleId)

	if erro != nil {
		return models.Article{}, erro
	}
	defer row.Close()

	var article models.Article
	if row.Next() {

		if erro = row.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.Likes,
			&article.AuthorId,
			&article.CreatedAt,
			&article.AuthorNick,
		); erro != nil {
			return models.Article{}, erro
		}
	}

	return article, nil
}
