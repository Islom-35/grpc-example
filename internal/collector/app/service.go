package app

import (
	"context"
	"fmt"
	"imantask/internal/collector/domain"
	"imantask/internal/genproto/pb"
	"sync"
)

type CollectorService interface {
	Save(context.Context, *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error)
}

type collectorService struct {
	repo         domain.CollectorRepository
	postProvider domain.PostProvider
}

func NewCollectorService(repo domain.CollectorRepository, postProvider domain.PostProvider) CollectorService {
	return &collectorService{
		repo:         repo,
		postProvider: postProvider,
	}
}

func (p *collectorService) Save(ctx context.Context, req *pb.CollectPostsRequest) (*pb.CollectPostsResponse, error) {
	var wg sync.WaitGroup

	numWorkers := 50

	postAllChannel := make(chan []domain.Data, numWorkers)

	errorChannel := make(chan error, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			datas, err := p.postProvider.GetPosts(page)
			if err != nil {
				errorChannel <- err
				return
			}
			postAllChannel <- datas.Data
		}(i)
	}
	// Close postAllChannel after all goroutines have finished
	go func() {
		wg.Wait()
		close(postAllChannel)
	}()

	// Process data from postAllChannel
	count := 0

	for posts := range postAllChannel {
		for _, post := range posts {
			err := p.repo.Save(post)
			if err != nil {
				errorChannel <- err
			}
			count++
		}
	}

	defer fmt.Println(count)

	// Close errorChannel after all errors have been processed
	close(errorChannel)

	// Wait for all goroutines to finish
	wg.Wait()

	// Check for errors

	err := <-errorChannel
	if err != nil {
		return nil, err
	}

	return &pb.CollectPostsResponse{}, nil
}
