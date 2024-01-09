package adapters

import (
	domain "imantask/internal/collector/domain"

	"github.com/jackc/pgx"
)

type PostRepo struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &PostRepo{db: db}
}

func (b *PostRepo) Save(post domain.Data) error {
	_, err := b.db.Exec("INSERT INTO data (id, user_id, title, body) values ($1, $2, $3, $4)",
		post.ID, post.UserID, post.Title, post.Body)

	return err
}
