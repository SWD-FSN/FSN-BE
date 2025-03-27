package service

import (
	"context"
	"database/sql"
	"log"
	business_object "social_network/business_object"
	"social_network/dto"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
	"strings"
	"sync"
)

type searchObjectService struct {
	postRepo repo.IPostRepo
	likeRepo repo.ILikeRepo
	userRepo repo.IUserRepo
	logger   *log.Logger
}

func InitializeSearchObjectsService(db *sql.DB, logger *log.Logger) service.ISearchObjectService {
	return &searchObjectService{
		postRepo: repository.InitializePostRepo(db, logger),
		likeRepo: repository.InitializeLikeRepo(db, logger),
		userRepo: repository.InitializeUserRepo(db, logger),
		logger:   logger,
	}
}

func GenerateSearchObjectsService() (service.ISearchObjectService, error) {
	var logger = util.GetLogConfig()

	db, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeSearchObjectsService(db, logger), nil
}

// GetObjectsByKeyword implements service.ISearchObjectService.
func (s *searchObjectService) GetObjectsByKeyword(id string, keyword string, ctx context.Context) *dto.GetObjectsFromEnterSearchBarResponse {
	var usersRes []dto.UserSearchDoneResponse
	var postsRes []dto.PostResponse

	var wg sync.WaitGroup
	wg.Add(2)

	// Fetch users
	go func() {
		defer wg.Done()

		tmpStorage, _ := s.userRepo.GetInvolvedAccountsFromTag(id, ctx)
		var idMap map[string]string = make(map[string]string)

		for _, id := range tmpStorage {
			// Lọc trùng user
			// Vd list follow và list friend có cùng 1 user
			if _, isExist := idMap[id]; !isExist {
				account, _ := s.userRepo.GetUser(id, ctx)

				if account != nil && !account.IsPrivate && strings.Contains(strings.ToLower(account.Username), strings.ToLower(keyword)) {
					usersRes = append(usersRes, dto.UserSearchDoneResponse{
						UserId:            id,
						Username:          account.Username,
						ProfileAvatar:     account.ProfileAvatar,
						IsFriendWithActor: strings.Contains(account.Friends, id),
						FollowerAmount:    len(util.ToSliceString(account.Followers, sepChar)),
					})
				}

				idMap[id] = id
			}
		}

		users, _ := s.userRepo.GetUsersByKeyword(keyword, ctx)
		for _, user := range *users {
			if _, isExist := idMap[user.UserId]; !isExist {
				if user.UserId != "" && !user.IsPrivate {
					usersRes = append(usersRes, dto.UserSearchDoneResponse{
						UserId:            user.UserId,
						Username:          user.Username,
						ProfileAvatar:     user.ProfileAvatar,
						IsFriendWithActor: strings.Contains(user.Friends, id),
						FollowerAmount:    len(util.ToSliceString(user.Followers, sepChar)),
					})

					idMap[user.UserId] = user.UserId
				}
			}
		}

	}()

	// Fetch posts
	go func() {
		defer wg.Done()

		posts, _ := s.postRepo.GetPostsByKeyword(keyword, ctx)

		for _, post := range *posts {
			likes, _ := s.likeRepo.GetLikesFromObject(post.PostId, business_object.GetPostTable(), ctx)
			author, _ := s.userRepo.GetUser(post.AuthorId, ctx)

			if author != nil {
				var likesAmount int = 0
				if likes != nil {
					likesAmount = len(*likes)
				}
				postsRes = append(postsRes, dto.PostResponse{
					PostId:        post.PostId,
					Content:       post.Content,
					IsPrivate:     post.IsPrivate,
					IsHidden:      post.IsHidden,
					LikeAmount:    likesAmount,
					CreatedAt:     post.CreatedAt,
					AuthorId:      author.UserId,
					Username:      author.Username,
					ProfileAvatar: author.ProfileAvatar,
				})
			}
		}
	}()

	wg.Wait()

	return &dto.GetObjectsFromEnterSearchBarResponse{
		Users: &usersRes,
		Posts: &postsRes,
	}
}
