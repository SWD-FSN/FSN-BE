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
	likeRepo repo.ILikeRepo
	userRepo repo.IUserRepo
	postRepo repo.IPostRepo
	logger   *log.Logger
}

func InitializePostService(db *sql.DB, logger *log.Logger) service.IPostService {
	return &postService{
		likeRepo: repository.InitializeLikeRepo(db, logger),
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
	if err := verifyUser(id, id_validate, p.userRepo, ctx); err != nil {
		return nil, err
	}

	return p.postRepo.GetPostsByUser(id, ctx)
}

// RemovePost implements service.IPostService.
func (p *postService) RemovePost(id string, actorId string, ctx context.Context) error {
	var cmtRepo repo.ICommentRepo = nil
	if !isObjectBelongActor(id, post_object, actorId, p.likeRepo, cmtRepo, p.postRepo, ctx) {
		return errors.New(noti.GenericsRightAccessWarnMsg)
	}

	return p.postRepo.RemovePost(id, ctx)
}

// UpPost implements service.IPostService.
func (p *postService) UpPost(req dto.UpPostReq, ctx context.Context) error {
	if err := verifyUser(req.AuthorId, id_validate, p.userRepo, ctx); err != nil {
		return err
	}

	// Set status
	if req.IsPrivate == nil {
		*req.IsPrivate = false
	}

	if req.IsHidden == nil {
		*req.IsHidden = false
	}

	var curTime time.Time = time.Now()
	return p.postRepo.CreatePost(business_object.Post{
		PostId:     util.GenerateId(),
		AuthorId:   req.AuthorId,
		Content:    req.Content,
		Attachment: req.Attachment,
		IsPrivate:  *req.IsPrivate,
		IsHidden:   *req.IsHidden,
		CreatedAt:  curTime,
		UpdatedAt:  curTime,
		Status:     true,
	}, ctx)
}

// GetPosts implements service.IPostService.
func (p *postService) GetPosts(ctx context.Context) *[]dto.PostResponse {
	var res []dto.PostResponse
	posts, _ := p.postRepo.GetPosts(ctx)

	for _, post := range *posts {
		likes, _ := p.likeRepo.GetLikesFromObject(post.PostId, business_object.GetPostTable(), ctx)
		author, _ := p.userRepo.GetUser(post.AuthorId, ctx)

		if author != nil {
			var likeAmounts int = 0
			if likes != nil {
				likeAmounts = len(*likes)
			}
			res = append(res, dto.PostResponse{
				PostId:        post.PostId,
				Content:       post.Content,
				Attachment:    post.Attachment,
				IsPrivate:     post.IsPrivate,
				IsHidden:      post.IsHidden,
				LikeAmount:    likeAmounts,
				CreatedAt:     post.CreatedAt,
				AuthorId:      author.UserId,
				Username:      author.Username,
				ProfileAvatar: author.ProfileAvatar,
			})
		}
	}

	return &res
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

	if req.Attachment != "" {
		post.Attachment = req.Attachment
	}

	post.UpdatedAt = time.Now()

	return p.postRepo.UpdatePost(*post, ctx)
}
