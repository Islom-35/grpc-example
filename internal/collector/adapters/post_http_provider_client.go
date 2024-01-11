package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	domain "imantask/internal/collector/domain"
	"io"
	"net/http"
	"strconv"

	"time"
)

type Client struct {
	Client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		Client: &http.Client{
			Timeout: timeout,
		},
	}, nil

}

func (c Client) GetPosts(page int) (domain.PostsResponse, error) {
	url := "https://gorest.co.in/public/v1/posts?page=" + strconv.Itoa(page)
	resp, err := c.Client.Get(url)
	if err != nil {
		return domain.PostsResponse{}, err
	}

	defer resp.Body.Close()
	fmt.Println("Response status: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.PostsResponse{}, err
	}

	var r domain.PostsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return domain.PostsResponse{}, err
	}

	return r, nil
}

