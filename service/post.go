package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
	"sync"
	"time"
)

type postService struct {
	userRepo repo.IUserRepo
	postRepo repo.IPostRepo
	logger   *log.Logger
}

func InitializePostService(db *sql.DB, logger *log.Logger) service.IPostService {
	return &postService{
		userRepo: repository.InitializeUserRepo(db, logger),
		postRepo: repository.InitializePostRepo(db, logger),
		logger:   logger,
	}
}

func GeneratePostService() (service.IPostService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializePostService(cnn, logger), nil
}

// GetAllPosts implements service.IPostService.
func (p *postService) GetAllPosts(ctx context.Context) (*[]business_object.Post, error) {
	return p.postRepo.GetAllPosts(ctx)
}

// GetPost implements service.IPostService.
func (p *postService) GetPost(id string, ctx context.Context) (*business_object.Post, error) {
	return p.postRepo.GetPost(id, ctx)
}

// GetPostsByUser implements service.IPostService.
func (p *postService) GetPostsByUser(id string, ctx context.Context) (*[]business_object.Post, error) {
	if err := verifyUser(id, p.userRepo, ctx); err != nil {
		return nil, err
	}

	res, err := p.postRepo.GetPostsByUser(id, ctx)
	if err != nil {
		return nil, err
	}

	// Sort
	util.SortByTime(*res, func(item business_object.Post) time.Time {
		return item.CreatedAt
	}, false)

	return res, nil
}

// RemovePost implements service.IPostService.
func (p *postService) RemovePost(id string, ctx context.Context) error {
	return p.postRepo.RemovePost(id, ctx)
}

// UpPost implements service.IPostService.
func (p *postService) UpPost(req dto.UpPostReq, ctx context.Context) error {
	if err := verifyUser(req.AuthorId, p.userRepo, ctx); err != nil {
		return err
	}

	// Set status
	if req.IsPrivate == nil {
		*req.IsPrivate = false
	}

	if req.IsHidden == nil {
		*req.IsHidden = false
	}

	return p.postRepo.CreatePost(business_object.Post{
		PostId:    util.GenerateId(),
		AuthorId:  req.AuthorId,
		Content:   req.Content,
		IsPrivate: *req.IsPrivate,
		IsHidden:  *req.IsHidden,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Status:    true,
	}, ctx)
}

// UpdatePost implements service.IPostService.
func (p *postService) UpdatePost(req dto.UpdatePostReq, ctx context.Context) error {
	post, err := p.postRepo.GetPost(req.PostId, ctx)
	if err != nil {
		return err
	}

	if post == nil {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	// Initialize waiting variable for 3 goroutines
	var wg sync.WaitGroup
	wg.Add(3)

	// Setting status
	go func() {
		defer wg.Done()
		if !util.IsBooleanRemain(req.IsPrivate, post.IsPrivate) {
			post.IsPrivate = *req.IsPrivate
		}
	}()

	go func() {
		defer wg.Done()
		if !util.IsBooleanRemain(req.IsHidden, post.IsHidden) {
			post.IsHidden = *req.IsHidden
		}
	}()

	go func() {
		defer wg.Done()
		if !util.IsBooleanRemain(req.Status, post.Status) {
			post.Status = *req.Status
		}
	}()

	// Wait for 3 goroutines to be done
	wg.Wait()

	post.UpdatedAt = time.Now().UTC()

	return p.postRepo.UpdatePost(*post, ctx)
}
