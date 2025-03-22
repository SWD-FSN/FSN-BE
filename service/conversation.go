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
	"strings"
	"sync"
	"time"
)

type conersationService struct {
	userRepo         repo.IUserRepo
	msgRepo          repo.IMessageRepo
	conversationRepo repo.IConversationRepo
	logger           *log.Logger
}

func InitializeConversationService(db *sql.DB, logger *log.Logger) service.IConversationService {
	return &conersationService{
		userRepo:         repository.InitializeUserRepo(db, logger),
		msgRepo:          repository.InitializeMessageRepo(db, logger),
		conversationRepo: repository.InitializeConversationRepo(db, logger),
		logger:           logger,
	}
}

func GenerateMessageService() (service.IConversationService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeConversationService(cnn, logger), nil
}

const (
	group_name_property   string = "group_name_property"
	group_avatar_property string = "group_avatar_property"
)

// CreateMessage implements service.IConversationService.
func (c *conersationService) CreateMessage(req dto.CreateMessageRequest, ctx context.Context) error {
	var conversation *business_object.Conversation
	if err := verifyActorToConversationAction(req.AuthorId, req.ConversationId, false, conversation, c.conversationRepo, ctx); err != nil {
		return err
	}

	// Create message
	var curTime time.Time = time.Now()
	if err := c.msgRepo.CreateMessage(business_object.Message{
		MessageId:     util.GenerateId(),
		AuthorId:      req.AuthorId,
		CoversationId: req.ConversationId,
		Content:       req.Content,
		CreatedAt:     curTime,
	}, ctx); err != nil {
		return err
	}

	actor, _ := c.userRepo.GetUser(req.AuthorId, ctx)

	sendMsgSocket(req.ConversationId, conversation_object,
		actor.Username, actor.ProfileAvatar, "message",
		"", curTime, &req.ConversationId, nil,
		nil, c.conversationRepo, ctx)

	return nil
}

// GetConversationFromUser implements service.IConversationService.
func (c *conersationService) GetConversationFromUser(actorId string, conversationId string, ctx context.Context) (*dto.InternalConversationUIResponse, error) {
	var conversation *business_object.Conversation
	if err := verifyActorToConversationAction(actorId, conversationId, false, conversation, c.conversationRepo, ctx); err != nil {
		return nil, err
	}

	msgs, _ := c.msgRepo.GetMessagesFromConversation(conversationId, ctx)

	// Sort messages
	util.SortByTime(*msgs, func(entity business_object.Message) time.Time {
		return entity.CreatedAt
	}, false)

	var actorMsgs, memberMsgs *[]dto.MessageUIResponse
	for _, msg := range *msgs {
		msgAuthor, _ := c.userRepo.GetUser(msg.AuthorId, ctx)

		var msgRes = dto.MessageUIResponse{
			ConvesationId: conversationId,
			MessageId:     msg.MessageId,
			AuthorId:      msg.AuthorId,
			AuthorAvatar:  msgAuthor.ProfileAvatar,
			Content:       msg.Content,
			CreatedAt:     msg.CreatedAt,
		}

		if msgAuthor.UserId == actorId {
			*actorMsgs = append(*actorMsgs, msgRes)
		} else {
			*memberMsgs = append(*memberMsgs, msgRes)
		}
	}

	avatar, name, err := processConversationAvatarAndName(conversation, actorId, c.userRepo, ctx)
	if err != nil {
		return nil, err
	}

	return &dto.InternalConversationUIResponse{
		ConversationId:     conversationId,
		ConversationAvatar: avatar,
		ConversationName:   name,
		ActorMessages:      actorMsgs,
		MemberMessages:     memberMsgs,
	}, nil
}

// EditGroupChatProperty implements service.IConversationService.
func (c *conersationService) EditGroupChatProperty(req dto.EditGroupChatPropRequest, ctx context.Context) error {
	var conversation *business_object.Conversation
	if err := verifyActorToConversationAction(req.ActorId, req.ConversationId, false, conversation, c.conversationRepo, ctx); err != nil {
		return err
	}

	// Not group chat or empty data
	if conversation.HostId == nil || req.Value == "" {
		return nil
	}

	switch req.Property {
	case group_avatar_property:
		conversation.ConversationAvatar = req.Property
	case group_name_property:
		conversation.ConversationName = req.Property
	default:
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	conversation.UpdatedAt = time.Now()

	return c.conversationRepo.UpdateConversation(*conversation, ctx)
}

// GetMessagesInChatByKeyword implements service.IConversationService.
func (c *conersationService) GetMessagesInChatByKeyword(req dto.SearchMessagesInChatRequest, ctx context.Context) (*dto.SearchMessagesInChatResponse, error) {
	var conversation *business_object.Conversation
	if err := verifyActorToConversationAction(req.ActorId, req.ConversationId, false, conversation, c.conversationRepo, ctx); err != nil {
		return nil, err
	}

	msgs, _ := c.msgRepo.GetMessagesFromConversationByKeyword(req.ConversationId, req.Keyword, ctx)

	// Sort messages
	util.SortByTime(*msgs, func(entity business_object.Message) time.Time {
		return entity.CreatedAt
	}, false)

	var msgsRes *[]dto.MessageUIResponse
	for _, msg := range *msgs {
		msgAuthor, _ := c.userRepo.GetUser(msg.AuthorId, ctx)

		*msgsRes = append(*msgsRes, dto.MessageUIResponse{
			ConvesationId: req.ConversationId,
			MessageId:     msg.MessageId,
			AuthorId:      msg.AuthorId,
			AuthorAvatar:  msgAuthor.ProfileAvatar,
			Content:       msg.Content,
			CreatedAt:     msg.CreatedAt,
		})
	}

	avatar, name, err := processConversationAvatarAndName(conversation, req.ActorId, c.userRepo, ctx)
	if err != nil {
		return nil, err
	}

	return &dto.SearchMessagesInChatResponse{
		ConversationId:     req.ConversationId,
		ConversationName:   name,
		ConversationAvatar: avatar,
		Messages:           msgsRes,
	}, nil
}

// CreateConversation implements service.IConversationService.
func (c *conersationService) CreateConversation(req dto.CreateConversationRequest, ctx context.Context) (*dto.ConversationUIResponse, error) {
	var capturedErr error
	var actor *dto.UserDBResModel
	var members *[]dto.UserDBResModel
	var memberUsernames []string

	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	// Verify actor
	go func() {
		defer wg.Done()

		if err := verifyAccount(req.ActorId, id_validate, actor, c.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	// Verify member(s)
	go func() {
		defer wg.Done()

		for _, id := range req.Members {
			var member *dto.UserDBResModel

			if err := verifyAccount(id, id_validate, member, c.userRepo, ctx); err != nil {
				mu.Lock()

				if capturedErr == nil {
					capturedErr = err
					cancel()
				}

				mu.Unlock()
				break
			}

			*members = append(*members, *member)
			memberUsernames = append(memberUsernames, *&member.Username)
		}
	}()

	wg.Wait()

	if capturedErr != nil {
		return nil, capturedErr
	}

	// Prepare creating conversation
	var isGroup bool = (len(req.Members) + 1) > 2

	var hostId *string
	if isGroup {
		*hostId = req.ActorId
	}

	var conversationName string
	if !isGroup {
		conversationName = actor.Username + sepChar + (*members)[0].Username
	} else {
		conversationName = util.ToCombinedString(append(memberUsernames, req.ActorId), sepChar)
	}

	var avatar string
	if !isGroup {
		avatar = actor.ProfileAvatar + sepChar + (*members)[0].ProfileAvatar
	}

	var curTime time.Time = time.Now()

	var id string = util.GenerateId()

	if err := c.conversationRepo.CreateConversation(business_object.Conversation{
		ConversationId:     id,
		ConversationAvatar: avatar,
		ConversationName:   conversationName,
		HostId:             hostId,
		Members:            util.ToCombinedString(append(req.Members, req.ActorId), sepChar),
		IsGroup:            isGroup,
		CreatedAt:          curTime,
		UpdatedAt:          curTime,
	}, ctx); err != nil {
		return nil, err
	}

	// Re-set conversation UI to actor if this is not a group chat
	if !isGroup {
		conversationName = (*members)[0].Username
		avatar = (*members)[0].ProfileAvatar
	}

	// Send back conversation UI
	return &dto.ConversationUIResponse{
		ConversationId:     id,
		ConversationAvatar: avatar,
		ConversationName:   conversationName,
	}, nil
}

// DissovelGroupConversation implements service.IConversationService.
func (c *conersationService) DissovelGroupConversation(actorId string, conversationId string, ctx context.Context) error {
	if err := verifyActorToConversationAction(actorId, conversationId, true, nil, c.conversationRepo, ctx); err != nil {
		return err
	}

	return c.conversationRepo.DissovelGroupConversation(conversationId, ctx)
}

// GetAllConversations implements service.IConversationService.
func (c *conersationService) GetAllConversations(ctx context.Context) (*[]dto.ConversationResponse, error) {
	panic("unimplemented")
}

// GetConversationsByKeywordFromUser implements service.IConversationService.
func (c *conersationService) GetConversationsByKeywordFromUser(id string, keyword string, ctx context.Context) *[]dto.ConversationSearchBarResponse {
	tmpStorage, _ := c.conversationRepo.GetConversationsByKeyword(id, keyword, ctx)

	// Extract conversation(s)
	var res *[]dto.ConversationSearchBarResponse
	for _, chat := range *tmpStorage {
		avatar, name, err := processConversationAvatarAndName(&chat, id, c.userRepo, ctx)

		if err != nil {
			*res = append(*res, dto.ConversationSearchBarResponse{
				ConversationId:     chat.ConversationId,
				ConversationAvatar: avatar,
				ConversationName:   name,
			})
		}
	}

	return res
}

// GetConversationsFromUser implements service.IConversationService.
func (c *conersationService) GetConversationsFromUser(id string, ctx context.Context) *[]dto.ConversationUIResponse {
	panic("unimplemented")
}

// LeaveGroupConversation implements service.IConversationService.
func (c *conersationService) LeaveGroupConversation(memberId string, conversationId string, ctx context.Context) error {
	var conversation *business_object.Conversation
	if err := verifyActorToConversationAction(memberId, conversationId, false, conversation, c.conversationRepo, ctx); err != nil {
		return err
	}

	var memberSlice []string = util.ToSliceString(conversation.Members, sepChar)
	var members string

	for index, id := range memberSlice {
		if id != memberId {
			members += id

			if index < len(memberSlice)-2 && index > 0 {
				members += sepChar
			}
		}
	}

	// Host
	if memberId == *conversation.HostId {
		*conversation.HostId = strings.Split(members, sepChar)[0]
	}

	conversation.UpdatedAt = time.Now()

	return c.conversationRepo.UpdateConversation(*conversation, ctx)
}

func verifyActorToConversationAction(actorId, conversationId string, isCheckHost bool, chat *business_object.Conversation, repo repo.IConversationRepo, ctx context.Context) error {
	isMember, err := isActorBelongToChat(actorId, conversationId, isCheckHost, chat, repo, ctx)

	if err != nil {
		return err
	}

	// Not belong to this conversation
	if !isMember {
		return errors.New(noti.GenericsRightAccessWarnMsg)
	}

	return nil
}

func isActorBelongToChat(actorId, conversationId string, isCheckHost bool, chat *business_object.Conversation, repo repo.IConversationRepo, ctx context.Context) (bool, error) {
	var errRes error

	chat, errRes = repo.GetConversation(conversationId, ctx)

	if errRes != nil {
		return false, errRes
	}

	if isCheckHost {
		return *chat.HostId == actorId, nil
	}

	return strings.Contains(chat.Members, actorId), nil
}

func processConversationAvatarAndName(conversation *business_object.Conversation, actorId string, userRepo repo.IUserRepo, ctx context.Context) (string, string, error) {
	if conversation.HostId == nil {
		var memberId string

		for _, id := range util.ToSliceString(conversation.Members, sepChar) {
			if id != actorId {
				memberId = id
				break
			}
		}

		member, err := userRepo.GetUser(memberId, ctx)

		if err != nil {
			return "", "", err
		}

		return member.ProfileAvatar, member.Username, nil
	}

	return conversation.ConversationAvatar, conversation.ConversationName, nil
}
