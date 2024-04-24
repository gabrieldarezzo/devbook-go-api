package models

import (
	"errors"
	"strings"
	"time"
)

// Article represent a post
type Article struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (article *Article) Prepare() error {
	if erro := article.validate(); erro != nil {
		return erro
	}

	article.format()
	return nil
}

func (article *Article) validate() error {

	if article.Title == "" {
		return errors.New("O campo: 'title' é obrigatório e não pode estar em branco")
	}

	if article.Content == "" {
		return errors.New("O campo: 'content' é obrigatório e não pode estar em branco")
	}

	return nil
}

func (article *Article) format() error {

	article.Title = strings.TrimSpace(article.Title)
	article.Content = strings.TrimSpace(article.Content)

	return nil
}
