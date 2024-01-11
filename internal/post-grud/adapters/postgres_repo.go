package adapters

import (
	"imantask/internal/post-grud/domain"

	"github.com/jackc/pgx"
)

type PostRepo struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &PostRepo{db: db}
}

func (b *PostRepo)GetByID(ID int)




