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

type likeService struct {
	userRepo    repo.IUserRepo
	postRepo    repo.IPostRepo
	commentRepo repo.ICommentRepo
	likeRepo    repo.ILikeRepo
	logger      *log.Logger
}

func InitializeLikeService(db *sql.DB, logger *log.Logger) service.ILikeService {
	return &likeService{
		userRepo: repository.InitializeUserRepo(db, logger),
		postRepo: repository.InitializePostRepo(db, logger),
		likeRepo: repository.InitializeLikeRepo(db, logger),
		logger:   logger,
	}
}

const (
	post_obj string = "POST_OBJ"
	cmt_obj  string = "CMT_OBJ"
)

func GenerateLikeService() (service.ILikeService, error) {
	cnn, err := db.ConnectDB(business_object.GetLikeTable())

	if err != nil {
		return nil, err
	}

	return InitializeLikeService(cnn, &log.Logger{}), nil
}

// DoLike implements service.ILikeService.
func (l *likeService) DoLike(req dto.DoLikeReq, ctx context.Context) error {
	var capturedErr error
	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	// Verify author
	go func() {
		if err := verifyUser(req.AuthorId, l.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel()          // Cancel the other goroutine
			}

			mu.Unlock()
		}
	}()

	// Verify object
	go func() {
		if err := verifyObject(req.ObjectId, req.ObjectType, l.commentRepo, l.postRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel()          // Cancel the other goroutine
			}

			mu.Unlock()
		}
	}()

	// Wait for 2 goroutines done
	wg.Wait()

	if capturedErr != nil {
		return capturedErr
	}

	return l.likeRepo.CreateLike(business_object.Like{
		LikeId:     util.GenerateId(),
		AuthorId:   req.AuthorId,
		ObjectId:   req.ObjectId,
		ObjectType: req.ObjectType,
		CreatedAt:  time.Now().UTC(),
	}, ctx)
}

// GetAllLikes implements service.ILikeService.
func (l *likeService) GetAllLikes(ctx context.Context) (*[]business_object.Like, error) {
	return l.likeRepo.GetAllLikes(ctx)
}

// GetLike implements service.ILikeService.
func (l *likeService) GetLike(id string, ctx context.Context) (*business_object.Like, error) {
	return l.likeRepo.GetLike(id, ctx)
}

// GetLikesFromObject implements service.ILikeService.
func (l *likeService) GetLikesFromObject(id string, kind string, ctx context.Context) (*[]business_object.Like, error) {
	return l.GetLikesFromObject(id, getObjectType(kind), ctx)
}

// UndoLike implements service.ILikeService.
func (l *likeService) UndoLike(id string, ctx context.Context) error {
	return l.likeRepo.CancelLike(id, ctx)
}

func verifyUser(id string, repo repo.IUserRepo, ctx context.Context) error {
	var errRes error = errors.New(noti.GenericsErrorWarnMsg)

	// Empty ID
	if id == "" {
		return errRes
	}

	user, err := repo.GetUser(id, ctx)
	if err != nil { // Error accessing db
		return err
	}

	if user == nil { // Not exists
		return errRes
	}

	return nil
}

func verifyObject(id, kind string, cmtRepo repo.ICommentRepo, postRepo repo.IPostRepo, ctx context.Context) error {
	var errRes error = errors.New(noti.GenericsErrorWarnMsg)

	// Empty ID
	if id == "" {
		return errRes
	}

	switch kind {
	case cmt_obj:
		cmt, err := cmtRepo.GetComment(id, ctx)
		if err != nil { // Error accessing db
			errRes = err
			break
		}

		if cmt != nil { // Exists
			errRes = nil
		}

	case post_obj:
		post, err := postRepo.GetPost(id, ctx)
		if err != nil { // Error accessing db
			errRes = err
			break
		}

		if post != nil { // Exists
			errRes = nil
		}
	}

	return errRes
}

func getObjectType(kind string) string {
	var res string

	switch kind {
	case cmt_obj:
		res = "comment"
	case post_obj:
		res = "post"
	}

	return res
}
