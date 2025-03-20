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

type commentService struct {
	likeRepo    repo.ILikeRepo
	postRepo    repo.IPostRepo
	notiRepo    repo.INotificationRepo
	userRepo    repo.IUserRepo
	commentRepo repo.ICommentRepo
	logger      *log.Logger
}

func InitializeCommentService(db *sql.DB, logger *log.Logger) service.ICommentService {
	return &commentService{
		likeRepo:    repository.InitializeLikeRepo(db, logger),
		postRepo:    repository.InitializePostRepo(db, logger),
		notiRepo:    repository.InitializeNotiRepo(db, logger),
		userRepo:    repository.InitializeUserRepo(db, logger),
		commentRepo: repository.InitializeCommentRepo(db, logger),
		logger:      logger,
	}
}

func GenerateCommentService() (service.ICommentService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeCommentService(cnn, logger), nil
}

const (
	user_object    string = "user"
	post_object    string = "post"
	like_object    string = "like"
	comment_object string = "comment"
)

// EditComment implements service.ICommentService.
func (c *commentService) EditComment(req dto.EditCommentRequest, ctx context.Context) error {
	if !isObjectBelongActor(req.CommentId, comment_object, req.ActorId, nil, c.commentRepo, nil, ctx) {
		return errors.New(noti.GenericsRightAccessWarnMsg)
	}

	cmt, err := c.commentRepo.GetComment(req.CommentId, ctx)
	if err != nil {
		return err
	}

	cmt.Content = req.Content
	cmt.UpdatedAt = time.Now()

	return c.commentRepo.EditComment(*cmt, ctx)
}

// GetCommentsFromPost implements service.ICommentService.
func (c *commentService) GetCommentsFromPost(id string, ctx context.Context) *[]dto.CommentDataResponse {
	tmpStorage, _ := c.commentRepo.GetCommentsFromPost(id, ctx)

	var res *[]dto.CommentDataResponse

	for _, cmt := range *tmpStorage {
		var user *dto.UserDBResModel
		var likes *[]business_object.Like
		var wg sync.WaitGroup
		wg.Add(2)

		// Get like(s) from post
		go func() {
			defer wg.Done()
			likes, _ = c.likeRepo.GetLikesFromObject(id, comment_object, ctx)
		}()

		// Get author
		go func() {
			defer wg.Done()
			user, _ = c.userRepo.GetUser(cmt.AuthorId, ctx)
		}()

		wg.Wait()

		// Add data
		*res = append(*res, dto.CommentDataResponse{
			AuthorId:      cmt.AuthorId,
			Author_avatar: user.ProfileAvatar,
			Username:      user.Username,
			CommentId:     cmt.CommentId,
			Content:       cmt.Content,
			CreatedAt:     cmt.CreatedAt,
			LikeAmount:    len(*likes),
		})
	}

	return res
}

// PostComment implements service.ICommentService.
func (c *commentService) PostComment(req dto.CreateCommentRequest, ctx context.Context) error {
	post, err := c.postRepo.GetPost(req.PostId, ctx)
	if err != nil {
		return err
	}

	actor, err := c.userRepo.GetUser(req.ActorId, ctx)
	if err != nil {
		return err
	}

	var res error = errors.New(noti.GenericsRightAccessWarnMsg)
	if post.IsHidden {
		return res
	}

	if post.IsPrivate {
		// Get friend list
		users, err := c.userRepo.GetInvoledAccountsAmountFromUser(dto.GetInvoledAccouuntsRequest{
			UserId:       post.AuthorId,
			InvolvedType: "friends",
		}, ctx)

		if err != nil {
			return err
		}

		// Check if actor is friend of post owner
		var isFriend bool = false
		for _, user := range users {
			if user == req.ActorId {
				isFriend = true
				break
			}
		}

		if !isFriend {
			return res
		}
	}

	// Create comment
	var curTime time.Time = time.Now()
	if err := c.commentRepo.CreateComment(business_object.Comment{
		CommentId: util.GenerateId(),
		AuthorId:  req.ActorId,
		PostId:    req.PostId,
		Content:   req.Content,
		CreatedAt: curTime,
		UpdatedAt: curTime,
	}, ctx); err != nil {
		return err
	}

	var objectType string = "post"
	var actionType string = "comment"
	c.notiRepo.CreateNotification(business_object.Notification{
		NotificationId: util.GenerateId(),
		ActorId:        req.ActorId,
		ObjectId:       req.PostId,
		ObjectType:     objectType,
		Action:         actionType,
		IsRead:         false,
		CreatedAt:      curTime,
	}, ctx)

	// Send message socket
	sendMsgSocket(req.PostId, objectType, actor.Username, actor.ProfileAvatar, actionType, "", curTime, nil, c.commentRepo, c.postRepo, ctx)

	return nil
}

// RemoveComment implements service.ICommentService.
func (c *commentService) RemoveComment(actorId string, commentId string, ctx context.Context) error {
	if !isObjectBelongActor(commentId, comment_object, actorId, nil, c.commentRepo, nil, ctx) {
		return errors.New(noti.GenericsRightAccessWarnMsg)
	}

	return c.commentRepo.RemoveComment(commentId, ctx)
}

func isObjectBelongActor(objectId, objectType, userId string, likeRepo repo.ILikeRepo, cmtRepo repo.ICommentRepo, postRepo repo.IPostRepo, ctx context.Context) bool {
	var res bool = false

	switch objectType {
	case post_object:
		obj, _ := postRepo.GetPost(objectId, ctx)
		res = obj.AuthorId == userId
	case like_object:
		obj, _ := likeRepo.GetLike(objectId, ctx)
		res = obj.AuthorId == userId
	case comment_object:
		obj, _ := cmtRepo.GetComment(objectId, ctx)
		res = obj.AuthorId == userId
	}

	return res
}
