package domain

import (
	"errors"
	pb "imantask/internal/genproto/ppb"
)

var (
	ErrorPostNotFound = errors.New("post not found")
	ThisPageDoesNotExist = errors.New("this page does not exist") 
)

type PostRepository interface {
	GetByID(ID int)(pb.PostResponse,error) 
	GetPage(offset,limit int)(pb.PostResponseList, error)
}