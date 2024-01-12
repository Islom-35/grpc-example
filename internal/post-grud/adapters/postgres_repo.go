package adapters

import (
	"fmt"
	pb "imantask/internal/genproto/ppb"
	"imantask/internal/post-grud/domain"
	"log"
	"strings"

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
		return pb.PostResponseList{}, domain.ErrorThisPageDoesNotExist
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

func (p *PostRepo) Update(inp pb.UpdateRequest) error{
	
	ID := int(inp.ID)

	setValues :=make([]string,0)
	args :=make([]interface{},0)
	argID :=1
	
	if inp.Title != ""{ 
		setValues = append(setValues, fmt.Sprintf("title=$%d",argID))
		args = append(args, inp.Title)
		argID++
	}

	if inp.Body != ""{ 
		setValues = append(setValues, fmt.Sprintf("body=$%d",argID))
		args = append(args, inp.Body)
		argID++
	}

	setQuery:=strings.Join(setValues,",")

	query := fmt.Sprintf("Update post Set %s where id=$%d", setQuery, argID)
	args = append(args, ID)
	_, err := p.db.Exec(query, args...)
		if err!=nil{
			log.Println(err)
			return err
		}
	return nil
}

func (p *PostRepo)Delete(ID int) error{
	_,err := p.db.Exec("DELETE FROM post WHERE id=$1", ID)
		if err!=nil{
			return err
		}
	return nil
}
