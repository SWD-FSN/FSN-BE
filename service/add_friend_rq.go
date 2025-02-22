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

type addFriendRqService struct {
	userRepo  repo.IUserRepo
	addFrRepo repo.IAddFriendRqRepo
	logger    *log.Logger
}

func InitializeAddFriendService(db *sql.DB, logger *log.Logger) service.IAddFriendRqService {
	return &addFriendRqService{
		userRepo:  repository.InitializeUserRepo(db, logger),
		addFrRepo: repository.InitializeAddFriendRqRepo(db, logger),
		logger:    logger,
	}
}

func GenerateAddFriendService() (service.IAddFriendRqService, error) {
	cnn, err := db.ConnectDB(business_object.GetAddFriendRequestTable())

	if err != nil {
		return nil, err
	}

	return InitializeAddFriendService(cnn, &log.Logger{}), nil
}

// AcceptAddFrRq implements service.IAddFriendRqService.
func (a *addFriendRqService) AcceptAddFrRq(requestId, actorId string, ctx context.Context) error {
	var req *business_object.AddFrRequest

	// Verify request
	if err := verifyAddFrRq(requestId, a.addFrRepo, req, ctx); err != nil {
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

		if err := verifyAccount(req.AuthorId, id_validate, author, a.userRepo, ctx); err != nil {
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

		if err := verifyAccount(req.AccountId, id_validate, account, a.userRepo, ctx); err != nil {
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

		if err := a.userRepo.UpdateUser(*author, ctx); err != nil {
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

		if err := a.userRepo.UpdateUser(*account, ctx); err != nil {
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
	return a.addFrRepo.RemoveAddFrRq(requestId, ctx)
}

// CancelAddFrRq implements service.IAddFriendRqService.
func (a *addFriendRqService) CancelAddFrRq(requestId string, actorId string, ctx context.Context) error {
	var req *business_object.AddFrRequest

	// Verify request
	if err := verifyAddFrRq(requestId, a.addFrRepo, req, ctx); err != nil {
		return err
	}

	// Verify sender
	if actorId != req.AccountId && actorId != req.AuthorId {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	// Remove request
	return a.addFrRepo.RemoveAddFrRq(requestId, ctx)
}

// GetAddFrRq implements service.IAddFriendRqService.
func (a *addFriendRqService) GetAddFrRq(id string, ctx context.Context) (*business_object.AddFrRequest, error) {
	return a.addFrRepo.GetAddFrRq(id, ctx)
}

// GetAddFrRqsToUser implements service.IAddFriendRqService.
func (a *addFriendRqService) GetAddFrRqsToUser(id string, ctx context.Context) (*[]business_object.AddFrRequest, error) {
	return a.addFrRepo.GetAddFrRqsToUser(id, ctx)
}

// GetAllAddFrRqs implements service.IAddFriendRqService.
func (a *addFriendRqService) GetAllAddFrRqs(ctx context.Context) (*[]business_object.AddFrRequest, error) {
	return a.addFrRepo.GetAllAddFrRqs(ctx)
}

// GetUserAddFrRqs implements service.IAddFriendRqService.
func (a *addFriendRqService) GetUserAddFrRqs(id string, ctx context.Context) (*[]business_object.AddFrRequest, error) {
	return a.addFrRepo.GetUserAddFrRqs(id, ctx)
}

// ProcessAddFrRq implements service.IAddFriendRqService.
func (a *addFriendRqService) ProcessAddFrRq(req dto.ActionRequest, ctx context.Context) error {
	if req.AuthorId == "" || req.AccountId == "" {
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

		if err := verifyAccount(req.AccountId, id_validate, account, a.userRepo, ctx); err != nil {
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

		if err := verifyUser(req.AuthorId, a.userRepo, ctx); err != nil {
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

	return a.addFrRepo.CreateAddFrRq(business_object.AddFrRequest{
		RequestId: util.GenerateId(),
		AuthorId:  req.AuthorId,
		AccountId: req.AccountId,
		CreatedAt: time.Now().UTC(),
	}, ctx)
}

func verifyAddFrRq(id string, repo repo.IAddFriendRqRepo, req *business_object.AddFrRequest, ctx context.Context) error {
	var res error

	req, res = repo.GetAddFrRq(id, ctx)
	if res != nil {
		return res
	}

	if req == nil {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	return nil
}
