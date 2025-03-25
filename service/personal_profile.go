package service

import (
	"context"
	"database/sql"
	"log"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
	"strings"
	"sync"
	"time"
)

type personalProfileService struct {
	conversationRepo repo.IConversationRepo
	postRepo         repo.IPostRepo
	likeRepo         repo.ILikeRepo
	userRepo         repo.IUserRepo
	logger           *log.Logger
}

func InitializePersonalProfileService(db *sql.DB, logger *log.Logger) service.IPersonalProfileService {
	return &personalProfileService{
		conversationRepo: repository.InitializeConversationRepo(db, logger),
		postRepo:         repository.InitializePostRepo(db, logger),
		likeRepo:         repository.InitializeLikeRepo(db, logger),
		userRepo:         repository.InitializeUserRepo(db, logger),
		logger:           logger,
	}
}

func GeneratePersonalProfileService() (service.IPersonalProfileService, error) {
	var logger = util.GetLogConfig()

	db, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializePersonalProfileService(db, logger), nil
}

// GetPersonalProfile implements service.IPersonalProfileService.
func (p *personalProfileService) GetPersonalProfile(actorId string, userId string, ctx context.Context) *dto.PersonalProfileUIResponse {
	var user *dto.UserDBResModel
	if verifyAccount(userId, id_validate, &user, p.userRepo, ctx) != nil {
		return nil
	}

	// Block
	if strings.Contains(user.BlockUsers, actorId) {
		return nil
	}

	// View themselves
	if actorId == userId {
		var posts *[]dto.PostResponse

		tmpStorage, _ := p.postRepo.GetPostsByUser(userId, ctx)
		for _, post := range *tmpStorage {
			likes, _ := p.likeRepo.GetLikesFromObject(post.PostId, post_obj, ctx)

			*posts = append(*posts, dto.PostResponse{
				PostId:        post.PostId,
				Content:       post.Content,
				IsPrivate:     post.IsPrivate,
				IsHidden:      post.IsHidden,
				LikeAmount:    len(*likes),
				AuthorId:      userId,
				CreatedAt:     post.CreatedAt,
				Username:      user.Username,
				ProfileAvatar: user.ProfileAvatar,
			})
		}

		// Sort post(s)
		util.SortByTime(*posts, func(entity dto.PostResponse) time.Time {
			return entity.CreatedAt
		}, false)

		return &dto.PersonalProfileUIResponse{
			UserId:        userId,
			Username:      user.Username,
			ProfileAvatar: user.ProfileAvatar,
			IsSelf:        true,
			Posts:         posts,
		}
	}

	var conversationId string
	var posts *[]dto.PostResponse
	var isFriend bool = strings.Contains(user.Friends, actorId)

	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch conversation
	go func() {
		defer wg.Done()
		conversation, _ := p.conversationRepo.GetConversationOfTwoUsers(actorId, userId, ctx)
		conversationId = conversation.ConversationId
	}()

	// Fetch post(s)
	go func() {
		defer wg.Done()

		tmpStorage, _ := p.postRepo.GetPostsByUser(userId, ctx)
		for _, post := range *tmpStorage {
			// Ko ẩn bài viết
			if !post.IsHidden {
				// Lấy lượt like của bài
				likes, _ := p.likeRepo.GetLikesFromObject(post.PostId, post_obj, ctx)

				// Tạo post theo model
				var postResponse = dto.PostResponse{
					PostId:        post.PostId,
					Content:       post.Content,
					IsPrivate:     post.IsPrivate,
					IsHidden:      post.IsHidden,
					LikeAmount:    len(*likes),
					AuthorId:      userId,
					CreatedAt:     post.CreatedAt,
					Username:      user.Username,
					ProfileAvatar: user.ProfileAvatar,
				}

				// Giới hạn bài viết cho bạn bè
				if post.IsPrivate {
					// Actor và account đc view là bạn bè
					if isFriend {
						*posts = append(*posts, postResponse)
					}
				} else { // Ko giới hạn bài viết
					*posts = append(*posts, postResponse)
				}
			}
		}

		// Sort post(s)
		util.SortByTime(*posts, func(entity dto.PostResponse) time.Time {
			return entity.CreatedAt
		}, false)
	}()

	wg.Wait()

	return &dto.PersonalProfileUIResponse{
		UserId:         userId,
		Username:       user.Username,
		ProfileAvatar:  user.ProfileAvatar,
		IsFriend:       isFriend,
		ConversationId: conversationId,
		Posts:          posts,
	}
}
