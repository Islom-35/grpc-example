package domain

type CollectorRepository interface {
	Save(data Data) error
}

type PostProvider interface {
	GetPosts(page int) (PostsResponse, error)
}