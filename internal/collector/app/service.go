package app

import (
	"context"
	"fmt"
	domain "imantask/internal/collector/domain"
	"imantask/internal/genproto/pb"

	"strconv"
)

type PostService interface {
	Save(context.Context, *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error)
}

type postService struct {
	repo         domain.PostRepository
	postProvider domain.PostProvider
}

func NewPostService(repo domain.PostRepository, postProvider domain.PostProvider) PostService {
	return &postService{
		repo:         repo,
		postProvider: postProvider,
	}
}

func (p *postService) Save(context.Context, *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error) {
	allPosts := make([]domain.Data, 0)

	for i := 1; i <= 50; i++ {
		posts, err := p.postProvider.GetPosts(strconv.Itoa(i))
		if err != nil {
			fmt.Println("Error:", err)
			return &pb.CollectPostsResponse{}, err
		}
		allPosts = append(allPosts, posts.Data...)
	}
	count := 0

	for _, post := range allPosts {

		fmt.Println(post)
		err := p.repo.Save(post)
		if err != nil {
			fmt.Println("Error saving post:", err)
		}else{
			count++
		}
		
	}

	fmt.Println("Total Posts:", count)
	return &pb.CollectPostsResponse{}, nil
}
