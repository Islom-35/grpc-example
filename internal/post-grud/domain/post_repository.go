package domain

import (
	"errors"
	pb "imantask/internal/genproto/ppb"
)

var (
	ErrorPostNotFound =errors.New("post not found")
)

type PostRepository interface {
	GetByID(ID int)(pb.PostResponse,error) 
}