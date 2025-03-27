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
	var user dto.UserDBResModel
	if verifyAccount(userId, id_validate, &user, p.userRepo, ctx) != nil {
		return nil
	}

	// Block
	if strings.Contains(user.BlockUsers, actorId) {
		return nil
	}

	tmpStorage, _ := p.postRepo.GetPostsByUser(userId, ctx)

	// View themselves
	if actorId == userId {
		var posts []dto.PostResponse

		for _, post := range *tmpStorage {
			likes, _ := p.likeRepo.GetLikesFromObject(post.PostId, post_obj, ctx)

			var likeAmount int = 0
			if likes != nil {
				likeAmount = len(*likes)
			}

			posts = append(posts, dto.PostResponse{
				PostId:        post.PostId,
				Content:       post.Content,
				Attachment:    post.Attachment,
				IsPrivate:     post.IsPrivate,
				IsHidden:      post.IsHidden,
				LikeAmount:    likeAmount,
				AuthorId:      userId,
				CreatedAt:     post.CreatedAt,
				Username:      user.Username,
				ProfileAvatar: user.ProfileAvatar,
			})
		}

		return &dto.PersonalProfileUIResponse{
			UserId:        userId,
			Username:      user.Username,
			ProfileAvatar: user.ProfileAvatar,
			IsSelf:        true,
			Posts:         &posts,
		}
	}

	var conversationId string
	var posts []dto.PostResponse
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

		for _, post := range *tmpStorage {
			// Ko ẩn bài viết
			if !post.IsHidden {
				// Lấy lượt like của bài
				likes, _ := p.likeRepo.GetLikesFromObject(post.PostId, post_obj, ctx)
				var likeAmount int = 0
				if likes != nil {
					likeAmount = len(*likes)
				}

				// Tạo post theo model
				var postResponse = dto.PostResponse{
					PostId:        post.PostId,
					Content:       post.Content,
					Attachment:    post.Attachment,
					IsPrivate:     post.IsPrivate,
					IsHidden:      post.IsHidden,
					LikeAmount:    likeAmount,
					AuthorId:      userId,
					CreatedAt:     post.CreatedAt,
					Username:      user.Username,
					ProfileAvatar: user.ProfileAvatar,
				}

				// Giới hạn bài viết cho bạn bè
				if post.IsPrivate {
					// Actor và account đc view là bạn bè
					if isFriend {
						posts = append(posts, postResponse)
					}
				} else { // Ko giới hạn bài viết
					posts = append(posts, postResponse)
				}
			}
		}
	}()

	wg.Wait()

	return &dto.PersonalProfileUIResponse{
		UserId:         userId,
		Username:       user.Username,
		ProfileAvatar:  user.ProfileAvatar,
		IsFriend:       isFriend,
		ConversationId: conversationId,
		Posts:          &posts,
	}
}
