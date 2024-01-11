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

func (p *PostRepo) GetPage(offset,limit int)(pb.PostResponseList, error){
	query := `
		SELECT id, user_id, title, body
		FROM post
		LIMIT $1 OFFSET $2
	`
	rows, err := p.db.Query(query, limit, offset)
	if err != nil {
		return pb.PostResponseList{}, domain.ThisPageDoesNotExist
	}
	defer rows.Close()

	posts :=pb.PostResponseList{}

	for rows.Next() {

		var post pb.PostResponse

		err := rows.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)

		if err !=nil{
			return pb.PostResponseList{},err
		}
		posts.Posts =append(posts.Posts, &post)
	}
	return posts,rows.Err()
}
