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

// FindArticle
func (repository ArticlesRepository) FindArticles(userId uint64) ([]models.Article, error) {

	// fmt.Printf("\nGetting articles of who user is following and himself articles: %d\n", userId)

	rows, erro := repository.db.Query(`
	select 
		DISTINCT 
		a.*,
		u.nick
	from 
	followers 
	inner join users u on (
		followers.user_id = u.id
	)
	inner join articles a on (
		u.id = a.author_id 
	)
	where 
		followers.follower_id = ?
		or followers.user_id  = ?
	order by a.created_at desc
	`, userId, userId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article
		if erro = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.AuthorId,
			&article.Likes,
			&article.CreatedAt,
			&article.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// CREATE TABLE articles(
//     id          int auto_increment primary key,
//     title       varchar(100) not null,
//     content     varchar(300) not null,
//     likes       INT default 0,
//     author_id   INT NOT NULL,
//     created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
//     FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
// ) ENGINE=InnoDB;

// FindArticle
func (repository ArticlesRepository) UpdateArticle(articleId uint64, article models.Article) error {

	_, erro := repository.db.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?",
		article.Title,
		article.Content,
		articleId,
	)
	if erro != nil {
		return erro
	}

	return nil
}
