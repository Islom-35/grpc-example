package domain

type PostRepository interface {
	Save(data Data) error
}

type PostProvider interface {
	GetPosts(page string) (PostsResponse, error)
}