package domain

type PostRepository interface {
	Save(data Post) error
}

type PostProvider interface {
	GetPosts(page int) (PostsResponse, error)
}