package adapters

import (
	pb "imantask/internal/genproto/ppb"
	"imantask/internal/post-grud/domain"

	"github.com/jackc/pgx"
)

type PostRepo struct {
	db *pgx.Conn
}

func NewPostRepository(db *pgx.Conn) domain.PostRepository {
	return &PostRepo{db: db}
}

func (p *PostRepo) GetByID(ID int) (pb.PostResponse, error) {
	var post pb.PostResponse
	err := p.db.QueryRow("SELECT id, user_id, title, body FROM post WHERE id=$1", ID).
		Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return pb.PostResponse{}, domain.ErrorPostNotFound
	}
	return post, nil
}
