package domain

type PostRepository interface {
	GetByID(ID int) 
}