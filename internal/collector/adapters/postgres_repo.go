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

func (b *CollectorRepo) Save(post domain.Data) error {
	_, err := b.db.Exec("INSERT INTO post (id, user_id, title, body) values ($1, $2, $3, $4)",
		post.ID, post.UserID, post.Title, post.Body)

	return err
}

