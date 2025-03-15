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

type socialRequestService struct {
	userRepo    repo.IUserRepo
	notiRepo    repo.INotificationRepo
	requestRepo repo.ISocialRequestRepo
	logger      *log.Logger
}

func InitializeSocialRequestService(db *sql.DB, logger *log.Logger) service.ISocialRequestService {
	return &socialRequestService{
		userRepo:    repository.InitializeUserRepo(db, logger),
		notiRepo:    repository.InitializeNotiRepo(db, logger),
		requestRepo: repository.InitializeSocialRequestRepo(db, logger),
		logger:      logger,
	}
}

func GenerateSocialRequestService() (service.ISocialRequestService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeSocialRequestService(cnn, logger), nil
}

const (
	follow_request     string = "follow"
	add_friend_request string = "add_friend"
)

// AcceptRequest implements service.ISocialRequestService.
func (s *socialRequestService) AcceptRequest(requestId, actorId string, ctx context.Context) error {
	var req *business_object.SocialRequest

	// Verify request
	if err := verifyRequest(requestId, s.requestRepo, req, ctx); err != nil {
		return err
	}

	// Verify reciever
	if actorId != req.AccountId {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	// Verify accounts
	var account *dto.UserDBResModel
	var author *dto.UserDBResModel

	var capturedErr error
	_, cancel1 := context.WithCancel(ctx)
	var wg1 sync.WaitGroup
	var mu1 sync.Mutex

	wg1.Add(2)

	// Verify author
	go func() {
		defer wg1.Done()

		if err := verifyAccount(req.AuthorId, id_validate, author, s.userRepo, ctx); err != nil {
			mu1.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel1()         // Cancel the other goroutine
			}

			mu1.Unlock()
		}
	}()

	// Verify requested account
	go func() {
		defer wg1.Done()

		if err := verifyAccount(req.AccountId, id_validate, account, s.userRepo, ctx); err != nil {
			mu1.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel1()         // Cancel the other goroutine
			}

			mu1.Unlock()
		}
	}()

	// Wait for 2 goroutines done
	wg1.Wait()

	if capturedErr != nil {
		return capturedErr
	}

	// Update friends
	_, cancel2 := context.WithCancel(ctx)
	var wg2 sync.WaitGroup
	var mu2 sync.Mutex

	wg2.Add(2)

	// Set friends
	account.Friends += sepChar + req.AuthorId
	author.Friends += sepChar + req.AccountId

	// Save changes to db
	go func() {
		defer wg2.Done()

		if err := s.userRepo.UpdateUser(*author, ctx); err != nil {
			mu2.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel2()         // Cancel the other goroutine
			}

			mu2.Unlock()
		}
	}()

	go func() {
		defer wg2.Done()

		if err := s.userRepo.UpdateUser(*account, ctx); err != nil {
			mu2.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel2()         // Cancel the other goroutine
			}

			mu2.Unlock()
		}
	}()

	wg2.Wait()

	if capturedErr != nil {
		return capturedErr
	}

	// Remove request
	return s.requestRepo.RemoveRequest(requestId, ctx)
}

// CancelRequest implements service.ISocialRequestService.
func (s *socialRequestService) CancelRequest(requestId string, actorId string, ctx context.Context) error {
	var req *business_object.SocialRequest

	// Verify request
	if err := verifyRequest(requestId, s.requestRepo, req, ctx); err != nil {
		return err
	}

	// Verify sender
	if actorId != req.AccountId && actorId != req.AuthorId {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	// Remove request
	return s.requestRepo.RemoveRequest(requestId, ctx)
}

// GetRequest implements service.ISocialRequestService.
func (s *socialRequestService) GetRequest(id string, ctx context.Context) (*business_object.SocialRequest, error) {
	return s.requestRepo.GetRequest(id, ctx)
}

// GetRequestsToUser implements service.ISocialRequestService.
func (s *socialRequestService) GetRequestsToUser(id string, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error) {
	if requestType != follow_request && requestType != add_friend_request {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	return s.requestRepo.GetRequestsToUser(id, requestType, ctx)
}

// GetAllRequests implements service.ISocialRequestService.
func (s *socialRequestService) GetAllRequests(ctx context.Context) (*[]business_object.SocialRequest, error) {
	return s.requestRepo.GetAllRequests(ctx)
}

// GetUserRequests implements service.ISocialRequestService.
func (s *socialRequestService) GetUserRequests(id string, requestType string, ctx context.Context) (*[]business_object.SocialRequest, error) {
	if requestType != follow_request && requestType != add_friend_request {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	return s.requestRepo.GetUserRequests(id, requestType, ctx)
}

// ProcessRequest implements service.ISocialRequestService.
func (s *socialRequestService) ProcessRequest(req dto.SocialRequest, ctx context.Context) error {
	if req.AuthorId == "" || req.AccountId == "" {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	if req.ActionType != follow_request && req.ActionType != add_friend_request {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	var capturedErr error
	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	var account *dto.UserDBResModel

	// Validate requested account
	go func() {
		defer wg.Done()

		if err := verifyAccount(req.AccountId, id_validate, account, s.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr != nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	// Validate author
	go func() {
		defer wg.Done()

		if err := verifyUser(req.AuthorId, s.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr != nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	wg.Wait()

	if capturedErr != nil {
		return capturedErr
	}

	// Account status as private -> can't request
	if account.IsPrivate {
		return errors.New("")
	}

	var curTime time.Time = time.Now()

	if err := s.requestRepo.CreateRequest(business_object.SocialRequest{
		RequestId:   util.GenerateId(),
		AuthorId:    req.AuthorId,
		AccountId:   req.AccountId,
		RequestType: req.ActionType,
		CreatedAt:   curTime,
	}, ctx); err != nil {
		return err
	}

	s.notiRepo.CreateNotification(business_object.Notification{
		NotificationId: util.GenerateId(),
		ActorId:        req.AuthorId,
		ObjectId:       req.AccountId,
		ObjectType:     "user",
		Action:         req.ActionType,
		IsRead:         false,
		CreatedAt:      curTime,
	}, ctx)

	return nil
}

func verifyRequest(id string, repo repo.ISocialRequestRepo, req *business_object.SocialRequest, ctx context.Context) error {
	var res error

	req, res = repo.GetRequest(id, ctx)
	if res != nil {
		return res
	}

	if req == nil {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	return nil
}
