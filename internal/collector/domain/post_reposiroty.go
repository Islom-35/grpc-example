package domain

type CollectorRepository interface {
	Save(data Post) error
}

type PostProvider interface {
	GetPosts(page int) (PostsResponse, error)
}