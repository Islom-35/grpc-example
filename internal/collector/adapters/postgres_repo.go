package adapters

import (
	domain "imantask/internal/collector/domain"

	"github.com/jackc/pgx"
)

type CollectorRepo struct {
	db *pgx.Conn
}

func NewCollectorRepository(db *pgx.Conn) domain.CollectorRepository {
	return &CollectorRepo{db: db}
}

func (b *CollectorRepo) Save(post domain.Post) error {
	_, err := b.db.Exec("INSERT INTO post (id, user_id, title, body,page) values ($1, $2, $3, $4, $5)",
		post.ID, post.UserID, post.Title, post.Body,post.Page)

	return err
}

