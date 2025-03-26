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
	notiRepo    repo.INotificationRepo
	likeRepo    repo.ILikeRepo
	logger      *log.Logger
}

func InitializeLikeService(db *sql.DB, logger *log.Logger) service.ILikeService {
	return &likeService{
		userRepo:    repository.InitializeUserRepo(db, logger),
		postRepo:    repository.InitializePostRepo(db, logger),
		commentRepo: repository.InitializeCommentRepo(db, logger),
		notiRepo:    repository.InitializeNotiRepo(db, logger),
		likeRepo:    repository.InitializeLikeRepo(db, logger),
		logger:      logger,
	}
}

func GenerateLikeService() (service.ILikeService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeLikeService(cnn, logger), nil
}

const (
	post_obj string = "POST_OBJ"
	cmt_obj  string = "CMT_OBJ"
)

// DoLike implements service.ILikeService.
func (l *likeService) DoLike(req dto.DoLikeReq, ctx context.Context) error {
	var capturedErr error
	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	var actor dto.UserDBResModel
	wg.Add(2)

	// Verify author
	go func() {
		defer wg.Done()
		if err := verifyAccount(req.ActorId, id_validate, &actor, l.userRepo, ctx); err != nil {
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
		defer wg.Done()
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

	var curTime time.Time = time.Now()

	if err := l.likeRepo.CreateLike(business_object.Like{
		LikeId:     util.GenerateId(),
		AuthorId:   req.ActorId,
		ObjectId:   req.ObjectId,
		ObjectType: req.ObjectType,
		CreatedAt:  curTime,
	}, ctx); err != nil {
		return err
	}

	// Create noti
	var actionType string = "like"
	l.notiRepo.CreateNotification(business_object.Notification{
		NotificationId: util.GenerateId(),
		ActorId:        req.ActorId,
		ObjectId:       req.ObjectId,
		ObjectType:     req.ObjectType,
		Action:         actionType,
		IsRead:         false,
		CreatedAt:      curTime,
	}, ctx)

	// Call socket hub to noti online user
	sendMsgSocket(req.ObjectId, req.ObjectType, actor.Username, actor.ProfileAvatar, actionType, "", curTime, nil, l.commentRepo, l.postRepo, nil, ctx)

	return nil
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
	like, err := l.likeRepo.GetLike(id, ctx)
	if err != nil {
		return err
	}

	notification, err := l.notiRepo.GetNotificationOnAction(dto.GetNotiOnActionRequest{
		ActorId:    like.AuthorId,
		ObjectId:   like.ObjectId,
		ObjectType: like.ObjectType,
		Action:     "like",
		CreatedAt:  like.CreatedAt,
	}, ctx)
	if err != nil {
		return err
	}

	if notification == nil {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	if err := l.notiRepo.RemoveNotification(notification.NotificationId, ctx); err != nil {
		return err
	}

	return l.likeRepo.CancelLike(id, ctx)
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

func getAuthorOfObject(objectId, objectType string, cmtRepo repo.ICommentRepo, postRepo repo.IPostRepo, conversationRepo repo.IConversationRepo, ctx context.Context) string {
	var res string

	switch objectType {
	case post_object:
		if post, _ := postRepo.GetPost(objectId, ctx); post != nil {
			res = post.AuthorId
		}
	case comment_object:
		if cmt, _ := cmtRepo.GetComment(objectId, ctx); cmt != nil {
			res = cmt.AuthorId
		}
	case conversation_object:
		if conversation, _ := conversationRepo.GetConversation(objectId, ctx); conversation != nil {
			res = conversation.Members
		}
	case user_object:
		res = objectId
	}

	return res
}

func sendMsgSocket(objectId, objectType, actorUsername, actorAvatar, actionType, orgContent string, createdAt time.Time, containerId *string, cmtRepo repo.ICommentRepo, postRepo repo.IPostRepo, conversationRepo repo.IConversationRepo, ctx context.Context) {
	if userId := getAuthorOfObject(objectId, objectType, cmtRepo, postRepo, conversationRepo, ctx); userId != "" {
		var users []string = util.ToSliceString(userId, sepChar)

		for _, user := range users {
			if cnn, isExist := clients[user]; isExist {
				content, contentType := generateContentAndContentTypeOfMsg(actorUsername, actionType, objectType, orgContent)

				sendMessage(dto.WSSendMessageRequest{
					UserId:             user,
					UserAvatar:         actorAvatar,
					Content:            content,
					ContentType:        contentType,
					CreatedAt:          createdAt,
					ContentContainerId: containerId,
				}, nil, cnn)
			}
		}
	}
}
