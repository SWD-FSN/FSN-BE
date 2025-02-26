package service

import (
	"context"
	"errors"
	"log"
	"os"
	business_object "social_network/business_object"
	action_type "social_network/constant/action_type"
	mail_const "social_network/constant/mail_const"
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

type userService struct {
	logger           *log.Logger
	roleRepo         repo.IRoleRepo
	userSecurityRepo repo.IUserSecurityRepo
	userRepo         repo.IUserRepo
}

func GenerateUserService() (service.IUserService, error) {
	db, err := db.ConnectDB(business_object.GetUserTable())

	if err != nil {
		return nil, err
	}

	var logger = util.GetLogConfig()

	return &userService{
		logger:           logger,
		roleRepo:         repository.InitializeRoleRepo(db, logger),
		userSecurityRepo: repository.InitializeUserSecurityRepo(db, logger),
		userRepo:         repository.InitializeUserRepo(db, logger),
	}, nil
}

const (
	sepChar        string = "|"
	mailSepChar    string = ":"
	id_validate    string = "ID_VALIDATE"
	email_validate string = "EMAIL_VALIDATE"
)

const (
	activateType      string = "1"
	resetPassType     string = "2"
	updateProfileType string = "3"
	verifyType        string = "4"
)

const (
	verifyFailLimit        int = 5
	minLengthVerifyCombine int = 3
)

const (
	admin_role string = "Admin"
	user_role  string = "User"
	staff_role string = "Staff"
)

const (
	friends_involed    string = "FRIENDS_INVOLED"
	blocks_involed     string = "BLOCKEDS_INVOLED"
	followers_involed  string = "FOLLOWERS_INVOLED"
	followings_involed string = "FOLLOWINGS_INVOLED"
)

func getProcessUrl() string {
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//------------------------------------
	return "http://localhost:" + port + "/VerifyAction?rawToken="
}

func getResetPassUrl() string {
	return "Your reset-pass URL page?token="
}

func getLoginUrl() string {
	return ""
}

// ChangeUserStatus implements service.IUserService.
// func (u *userService) ChangeUserStatus(rawStatus string, userId string, actorId string, c context.Context) (string, error) {
// 	panic("unimplemented")
// }

// GetUsersFromSearchBar implements service.IUserService.
func (u *userService) GetUsersFromSearchBar(id string, keyword string, ctx context.Context) *[]dto.GetInvolvedAccountsSearchResponse {
	var res *[]dto.GetInvolvedAccountsSearchResponse
	var maxResLength int = 8

	tmpStorage1, _ := u.userRepo.GetInvolvedAccountsFromTag(id, ctx)
	var idMap map[string]string = make(map[string]string)

	for _, id := range tmpStorage1 {
		// Lọc trùng user
		// Vd list follow và list friend có cùng 1 user
		if _, isExist := idMap[id]; !isExist {
			account, _ := u.userRepo.GetUser(id, ctx)

			if account != nil && strings.Contains(strings.ToLower(account.Username), strings.ToLower(keyword)) {
				*res = append(*res, dto.GetInvolvedAccountsSearchResponse{
					UserId:        id,
					Username:      account.Username,
					ProfileAvatar: account.ProfileAvatar,
				})
			}

			if len(*res) == maxResLength {
				return res
			}

			idMap[id] = id
		}
	}

	users, _ := u.userRepo.GetUsersByKeyword(keyword, ctx)
	for _, user := range *users {
		if _, isExist := idMap[user.UserId]; !isExist {
			*res = append(*res, dto.GetInvolvedAccountsSearchResponse{
				UserId:        user.UserId,
				Username:      user.Username,
				ProfileAvatar: user.ProfileAvatar,
			})

			if len(*res) == maxResLength {
				return res
			}

			idMap[user.UserId] = user.UserId
		}
	}

	return res
}

// GetInvolvedAccountsFromTag implements service.IUserService.
func (u *userService) GetInvolvedAccountsFromTag(id string, keyword string, ctx context.Context) *[]dto.GetInvolvedAccountsSearchResponse {
	tmpStorage, _ := u.userRepo.GetInvolvedAccountsFromTag(id, ctx)

	if len(tmpStorage) == 0 {
		return nil
	}

	var res *[]dto.GetInvolvedAccountsSearchResponse
	var idMap map[string]string = make(map[string]string)

	for _, id := range tmpStorage {
		// Lọc trùng user
		// Vd list follow và list friend có cùng 1 user
		if _, isExist := idMap[id]; !isExist {
			account, _ := u.userRepo.GetUser(id, ctx)

			if account != nil && strings.Contains(strings.ToLower(account.Username), strings.ToLower(keyword)) {
				*res = append(*res, dto.GetInvolvedAccountsSearchResponse{
					UserId:        id,
					Username:      account.Username,
					ProfileAvatar: account.ProfileAvatar,
				})
			}

			idMap[id] = id
		}
	}

	return res
}

// GetInvoledAccountsFromUser implements service.IUserService.
func (u *userService) GetInvoledAccountsFromUser(req dto.GetInvoledAccouuntsRequest, ctx context.Context) (*[]business_object.User, error) {
	var user *dto.UserDBResModel
	if err := verifyAccount(req.UserId, id_validate, user, u.userRepo, ctx); err != nil {
		return nil, err
	}

	var ids []string
	switch req.InvolvedType {
	case friends_involed:
		ids = util.ToSliceString(user.Friends, sepChar)
	case followers_involed:
		ids = util.ToSliceString(user.Followers, sepChar)
	case followings_involed:
		ids = util.ToSliceString(user.Followings, sepChar)
	case blocks_involed:
		ids = util.ToSliceString(user.BlockUsers, sepChar)
	default:
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}

	if len(ids) == 0 {
		return nil, nil
	}

	var res *[]business_object.User

	for _, id := range ids {
		account, _ := u.userRepo.GetUser(id, ctx)
		*res = append(*res, toUserModel(*account))
	}

	return res, nil
}

// GetAllUsers implements service.IUserService.
func (u *userService) GetAllUsers(ctx context.Context) (*[]business_object.User, error) {
	tmpStorage, err := u.userRepo.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// GetUsersByRole implements service.IUserService.
func (u *userService) GetUsersByRole(role string, ctx context.Context) (*[]business_object.User, error) {
	role = util.ToNormalizedString(role)

	var tmpStorage *[]dto.UserDBResModel
	var err error

	if role == "" {
		tmpStorage, err = u.userRepo.GetAllUsers(ctx)
	} else {
		tmpStorage, err = u.userRepo.GetUsersByRole(role, ctx)
	}

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// GetUsersByStatus implements service.IUserService.
func (u *userService) GetUsersByStatus(rawStatus string, ctx context.Context) (*[]business_object.User, error) {
	rawStatus = util.ToNormalizedString(rawStatus)

	var tmpStorage *[]dto.UserDBResModel
	var err error

	if rawStatus == "" {
		tmpStorage, err = u.userRepo.GetAllUsers(ctx)
	} else {
		status, errRes := util.ToBoolean(rawStatus)

		if errRes != nil {
			return nil, errRes
		}

		tmpStorage, err = u.userRepo.GetUsersByStatus(status, ctx)
	}

	if err != nil {
		return nil, err
	}

	return toSliceUserModel(tmpStorage), nil
}

// LogOut implements service.IUserService.
func (u *userService) LogOut(id string, ctx context.Context) error {
	return u.userSecurityRepo.LogOut(id, ctx)
}

// Login implements service.IUserService.
func (u *userService) Login(req dto.LoginRequest, ctx context.Context) (string, string, error) {
	var user *dto.UserDBResModel
	if err := verifyAccount(req.Email, email_validate, user, u.userRepo, ctx); err != nil {
		return "", "", err
	}

	if !util.IsHashStringMatched(req.Password, user.Password) {
		return processFailLogin(user.UserId, u.userSecurityRepo, ctx)
	}

	return processSuccessLogin(user, u.userSecurityRepo, ctx)
}

// CreateUser implements service.IUserService.
func (u *userService) CreateUser(req dto.CreateUserReq, actorId string, ctx context.Context) (string, error) {
	var actor *dto.UserDBResModel

	// If this request executed by an account
	if actorId != "" {
		if err := verifyAccount(actorId, id_validate, actor, u.userRepo, ctx); err != nil {
			return "", err
		}
	}

	// New email exists?
	if verifyAccount(req.Email, email_validate, nil, u.userRepo, ctx) == nil {
		return "", errors.New(noti.EmailRegisteredWarnMsg)
	}

	// Check password secure
	if !util.IsPasswordSecure(req.Password) {
		return "", errors.New(noti.PasswordNotSecureWarnMsg)
	}

	// Hash password
	hashPw, err := util.ToHashString(req.Password)
	if err != nil {
		return "", err
	}

	// Define role for new account
	roles, err := getRoles(u.roleRepo, ctx)
	if err != nil {
		return "", err
	}

	if actorId == "" || req.RoleId == "" {
		req.RoleId = roles[user_role]
	} else {
		if err := validateCreateAction(req.RoleId, actor.RoleId, roles); err != nil {
			return "", err
		}
	}

	var id string = util.GenerateId()

	// Generate token
	token, err := util.GenerateActionToken(req.Email, "", req.RoleId, u.logger)
	if err != nil {
		return "", err
	}

	// Belongs to last fail access of a new account
	var tmpTime = util.GetPrimitiveTime()

	// Flag if account need to reset password (in case staff role creates)
	var isHaveToResetPw *bool = nil
	if actorId != "" {
		var flag bool = true
		isHaveToResetPw = &flag
	}

	// Set status
	var isPrivate bool = true
	if req.IsPrivate == nil {
		isPrivate = false
	} else {
		isPrivate = *req.IsPrivate
	}

	var isActive bool = false
	if req.IsActive == nil {
		isActive = true
	} else {
		isPrivate = *req.IsActive
	}

	var fullName string = req.FullName
	if fullName == "" {
		fullName = req.Username
	}

	// Save new account to database
	if err := u.userRepo.CreateUser(dto.UserDBResModel{
		UserId:          id,
		RoleId:          req.RoleId,
		FullName:        fullName,
		Username:        req.Username,
		Email:           req.Email,
		Password:        hashPw,
		DateOfBirth:     req.DateOfBirth,
		ProfileAvatar:   req.ProfileAvatar,
		Bio:             req.Bio,
		Friends:         util.ToCombinedString(*req.Friends, sepChar),
		Followers:       util.ToCombinedString(*req.Followers, sepChar),
		Followings:      util.ToCombinedString(*req.Followings, sepChar),
		BlockUsers:      util.ToCombinedString(*req.BlockUsers, sepChar),
		IsPrivate:       isPrivate,
		IsActive:        isActive,
		IsActivated:     false,
		IsHaveToResetPw: isHaveToResetPw,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}, ctx); err != nil {
		return "", err
	}

	// Save new account security to database
	if err := u.userSecurityRepo.CreateUserSecurity(business_object.UserSecurity{
		UserId:      id,
		ActionToken: &token,
		FailAccess:  0,
		LastFail:    &tmpTime,
	}, ctx); err != nil {
		return "", err
	}

	// Send confirmation mail
	if err := util.SendMail(dto.SendMailRequest{
		Body: dto.MailBody{ // Mail body
			Email:    req.Email,
			Password: req.Password,
			Url: util.ToCombinedString([]string{ // Call back url when guest clicks to the confirmation, it will call back to the api endpoint which generate here to verify and finish the registration process
				getProcessUrl(),
				token,
				id,
				activateType,
			},
				mailSepChar),
		},

		TemplatePath: mail_const.AccountRegistrationTemplate, // Template path

		Subject: noti.RegistrationAccountSubject, // Mail subject

		Logger: u.logger, // Logger
	}); err != nil {
		return "", err
	}

	var msg string = "Success"
	if actorId == "" {
		msg = noti.RegistrationAccountMsg
	}

	return msg, nil
}

// GetUser implements service.IUserService.
func (u *userService) GetUser(id string, ctx context.Context) (*business_object.User, error) {
	var user *dto.UserDBResModel
	if err := verifyAccount(id, id_validate, user, u.userRepo, ctx); err != nil {
		return nil, err
	}

	var res = toUserModel(*user)
	return &res, nil
}

// ChangeUserStatus implements service.IUserService.
func (u *userService) ChangeUserStatus(rawStatus string, userId string, actorId string, ctx context.Context) (string, error) {
	panic("unimplemented")
}

// ResetPassword implements service.IUserService.
func (u *userService) ResetPassword(newPass string, re_newPass string, token string, ctx context.Context) (string, error) {
	id, _, exp, err := util.ExtractDataFromToken(token, u.logger)
	if err != nil {
		return getLoginUrl(), err
	}

	var user *dto.UserDBResModel
	if err := verifyAccount(id, id_validate, user, u.userRepo, ctx); err != nil {
		return getLoginUrl(), err
	}

	// User state doesn't have to reset password
	if user.IsHaveToResetPw == nil {
		return getLoginUrl(), errors.New(noti.GenericsErrorWarnMsg)
	}

	// Expired
	if util.IsActionExpired(exp) {
		return getLoginUrl(), errors.New("")
	}

	usc, err := u.userSecurityRepo.GetUserSecurity(id, ctx)
	if err != nil {
		return getLoginUrl(), err
	}

	if *usc.ActionToken != token {
		return getLoginUrl(), errors.New(noti.GenericsErrorWarnMsg)
	}

	// Passwords not matched
	if newPass != re_newPass {
		return util.ToCombinedString([]string{
			getResetPassUrl(),
			token,
		}, sepChar), errors.New(noti.GenericsErrorWarnMsg)
	}

	hashPw, err := util.ToHashString(newPass)
	if err != nil {
		return getLoginUrl(), err
	}

	// Setting new data
	user.Password = hashPw
	user.IsHaveToResetPw = nil

	usc.ActionToken = nil

	// Process update
	var capturedErr error

	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	// Update user data
	go func() {
		defer wg.Done()

		if err := u.userRepo.UpdateUser(*user, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	// Update user security
	go func() {
		defer wg.Done()

		if err := u.userSecurityRepo.EditUserSecurity(*usc, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	wg.Wait()

	return getLoginUrl(), capturedErr
}

// UpdateUser implements service.IUserService.
func (u *userService) UpdateUser(req dto.UpdateUserReq, actorId string, ctx context.Context) (string, error) {
	var res string
	var capturedErr error

	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	var account, actor *dto.UserDBResModel

	// Verify account
	go func() {
		defer wg.Done()

		if err := verifyAccount(req.UserId, id_validate, account, u.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	// Verify actor
	go func() {
		defer wg.Done()

		if err := verifyAccount(actorId, id_validate, actor, u.userRepo, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err
				cancel()
			}

			mu.Unlock()
		}
	}()

	// Wait for 2 goroutines to be done
	wg.Wait()

	if capturedErr != nil {
		return res, capturedErr
	}

	// Get roles
	roles, err := getRoles(u.roleRepo, ctx)
	if err != nil {
		return res, err
	}

	if err := verifyEditUserAuthorization(req, account, actor, roles); err != nil {
		return res, err
	}

	// Edit themselves
	if req.UserId == actorId {
		if !util.IsPasswordSecure(req.Password) {
			return res, errors.New(noti.PasswordNotSecureWarnMsg)
		}
	}

	hashPw, err := util.ToHashString(req.Password)
	if err != nil {
		return res, err
	}

	account.Password = hashPw

	var email string = util.ToNormalizedString(req.Email)
	var isHaveToVerify bool = false

	// Check if need to verify new email
	if email != account.Email {
		if err := verifyAccount(email, email_validate, nil, u.userRepo, ctx); err != errors.New("Lỗi email ko tồn tại") { // ~~ Tức mail tồn tại -> invalid
			return res, errors.New(noti.EmailRegisteredWarnMsg)
		}

		isHaveToVerify = true
	}

	if req.Username != "" {
		account.Username = req.Username
	}

	if req.DateOfBirth != nil {
		account.DateOfBirth = *req.DateOfBirth
	}

	if req.ProfileAvatar != "" {
		account.ProfileAvatar = req.ProfileAvatar
	}

	if req.Bio != "" {
		account.Bio = req.Bio
	}

	if req.IsPrivate != nil {
		account.IsPrivate = *req.IsPrivate
	}

	if err := u.userRepo.UpdateUser(*account, ctx); err != nil {
		return res, err
	}

	// Have to confirm new email before updating
	if isHaveToVerify {
		// Generate token
		token, err := util.GenerateActionToken(email, req.UserId, req.RoleId, u.logger)
		if err != nil {
			return res, err
		}

		// Get user security data to update token for new action - update new email
		usc, err := u.userSecurityRepo.GetUserSecurity(req.UserId, ctx)
		if err != nil {
			return res, err
		}

		// Update user security to db
		usc.ActionToken = &token
		if err := u.userSecurityRepo.EditUserSecurity(*usc, ctx); err != nil {
			return res, err
		}

		// Send verification mail
		if util.SendMail(dto.SendMailRequest{
			Body: dto.MailBody{
				Email: email,
				Url: util.ToCombinedString([]string{
					getProcessUrl(),
					token,
					req.UserId,
					updateProfileType,
					email,
				}, mailSepChar),
			},

			TemplatePath: mail_const.UpdateMailTemplate,

			Subject: noti.UpdateMailSubject,

			Logger: u.logger,
		}) != nil {
			return res, errors.New("We have updated your other information. " + noti.GenerateMailWarnMsg)
		}

		res = noti.UpdateMailMsg
	} else {
		res = "Success"
	}

	return res, nil
}

// VerifyAction implements service.IUserService.
func (u *userService) VerifyAction(rawToken string, ctx context.Context) (string, error) {
	var errRes error = errors.New(noti.GenericsErrorWarnMsg)

	var cmps []string = util.ToSliceString(rawToken, mailSepChar)
	if len(cmps) < minLengthVerifyCombine { // Min length of combination of information in a call back url
		return "", errRes
	}

	var token string = cmps[0]
	var id string = cmps[1]
	var actionType string = cmps[2]

	usc, err := u.userSecurityRepo.GetUserSecurity(id, ctx)
	if err != nil {
		return "", errRes
	}

	if *usc.ActionToken != token {
		return "", errRes
	}

	// Extract data from token
	extractId, _, exp, err := util.ExtractDataFromToken(rawToken, u.logger)
	if err != nil {
		return "", err
	}

	if extractId != id {
		return "", errRes
	}

	// Expired
	if util.IsActionExpired(exp) {
		return "", errors.New("")
	}

	// Invalid action type
	if actionType != activateType && actionType != resetPassType && actionType != updateProfileType && actionType != verifyType {
		return "", errors.New(noti.GenericsErrorWarnMsg)
	}

	var res string

	if actionType == activateType || actionType == updateProfileType {
		user, err := u.userRepo.GetUser(id, ctx)
		if err != nil {
			return res, err
		}

		if actionType == activateType {
			user.IsActivated = true
		} else {
			user.Email = cmps[len(cmps)-1]
		}

		if err := u.userRepo.UpdateUser(*user, ctx); err != nil {
			return res, err
		}

		res = getLoginUrl()
	} else {
		redirectUrl, err := setUpBeforeResetPw(usc, u.logger)
		if err != nil {
			return res, err
		}

		res = redirectUrl
	}

	return res, nil
}

func toSliceUserModel(src *[]dto.UserDBResModel) *[]business_object.User {
	var res *[]business_object.User

	for _, user := range *src {
		*res = append(*res, toUserModel(user))
	}

	return res
}

func toUserModel(src dto.UserDBResModel) business_object.User {
	var friends = util.ToSliceString(src.Friends, sepChar)
	var followers = util.ToSliceString(src.Followers, sepChar)
	var followings = util.ToSliceString(src.Followings, sepChar)
	var blockUsers = util.ToSliceString(src.BlockUsers, sepChar)

	return business_object.User{
		UserId:        src.UserId,
		RoleId:        src.RoleId,
		FullName:      src.FullName,
		Username:      src.Username,
		Email:         src.Email,
		Password:      src.Password,
		DateOfBirth:   src.DateOfBirth,
		ProfileAvatar: src.ProfileAvatar,
		Bio:           src.Bio,
		IsPrivate:     &src.IsPrivate,
		IsActive:      &src.IsActive,
		Friends:       &friends,
		Followers:     &followers,
		Followings:    &followings,
		BlockUsers:    &blockUsers,
		CreatedAt:     src.CreatedAt,
		UpdatedAt:     src.UpdatedAt,
	}
}

func verifyAccount(field, validateField string, user *dto.UserDBResModel, repo repo.IUserRepo, ctx context.Context) error {
	if field == "" {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	var res error

	switch validateField {
	case id_validate:
		user, res = repo.GetUser(field, ctx)
	case email_validate:
		user, res = repo.GetUserByEmail(field, ctx)
	}

	// User ko tồn tại
	if user == nil && res == nil {
		// Phân chia lỗi trả về
		switch validateField {
		case id_validate:
			res = errors.New("")
		case email_validate:
			res = errors.New("")
		}
	}

	return res
}

func prepareActivateAccount(security *business_object.UserSecurity, email, actionType, templatePath, subject string) error {
	var logger *log.Logger

	token, err := util.GenerateActionToken(email, security.UserId, "", logger)
	if err != nil {
		return err
	}

	*security.ActionToken = token

	return util.SendMail(dto.SendMailRequest{
		Body: dto.MailBody{ // Mail body
			Email: email,
			Url: util.ToCombinedString([]string{
				getProcessUrl(),
				token,
				security.UserId,
				actionType,
			}, mailSepChar),
		},

		TemplatePath: templatePath, // Đường dẫn đến template html để tạo mail theo format gửi đến user

		Subject: subject, // Tiêu đề mail

		Logger: logger,
	})
}

func setUpVerifyAccount(security *business_object.UserSecurity, email string, secureRepo repo.IUserSecurityRepo, ctx context.Context) (string, string, error) {
	var capturedErr error
	_, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mu sync.Mutex

	var res1, res2 string

	wg.Add(2)

	go func() {
		defer wg.Done()

		var actionType, templatePath, subject string

		if security.FailAccess > verifyFailLimit {
			actionType = verifyType
			templatePath = mail_const.AccountRecoveryTemplate
			subject = noti.VerifyAccountSubject

			res1 = action_type.VerifyCase
			res2 = noti.VerifyAccountMsg

		} else {
			actionType = activateType
			templatePath = mail_const.AccountRegistrationTemplate
			subject = noti.RegistrationAccountSubject

			res1 = action_type.ActivateCase
			res2 = noti.ActivateAccountMsg
		}

		if err := prepareActivateAccount(security, email, actionType, templatePath, subject); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel()          // Cancel the other goroutine
			}

			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		security.FailAccess = 0
		*security.LastFail = util.GetPrimitiveTime()

		if err := secureRepo.EditUserSecurity(*security, ctx); err != nil {
			mu.Lock()

			if capturedErr == nil {
				capturedErr = err // Capture the first error
				cancel()          // Cancel the other goroutine
			}

			mu.Unlock()
		}
	}()

	wg.Wait()

	if capturedErr != nil {
		return "", "", capturedErr
	}

	return res1, res2, nil
}

func processSuccessLogin(user *dto.UserDBResModel, securityRepo repo.IUserSecurityRepo, ctx context.Context) (string, string, error) {
	security, err := securityRepo.GetUserSecurity(user.UserId, ctx)

	if err != nil {
		return "", "", err
	}

	if !user.IsActivated || security.FailAccess > verifyFailLimit {
		return setUpVerifyAccount(security, user.Email, securityRepo, ctx)
	}

	var res1, res2 string

	if user.IsHaveToResetPw != nil && *user.IsHaveToResetPw {
		token, err := util.GenerateActionToken(user.Email, user.UserId, user.RoleId, log.Default())

		if err != nil {
			return "", "", err
		}

		res2 = util.ToCombinedString([]string{
			getResetPassUrl(),
			token,
		}, mailSepChar)

		*security.ActionToken = token
	} else {
		accessToken, refreshToken, err := util.GenerateTokens(user.Email, user.UserId, user.RoleId, log.Default())

		if err != nil {
			return "", "", err
		}

		res1 = accessToken
		res2 = refreshToken

		*security.AccessToken = accessToken
		*security.RefreshToken = refreshToken
	}

	return res1, res2, securityRepo.EditUserSecurity(*security, ctx)
}

func processFailLogin(id string, securityRepo repo.IUserSecurityRepo, ctx context.Context) (string, string, error) {
	security, _ := securityRepo.GetUserSecurity(id, ctx)

	if security != nil {
		security.FailAccess += 1
		*security.LastFail = time.Now().UTC()
		securityRepo.EditUserSecurity(*security, ctx)
	}

	return "", "", errors.New(noti.WrongCredentialsWarnMsg)
}

func getRoles(repo repo.IRoleRepo, ctx context.Context) (map[string]string, error) {
	roles, err := repo.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}

	var res = make(map[string]string)

	for _, role := range *roles {
		res[role.RoleName] = role.RoleId
	}

	return res, nil
}

func validateCreateAction(accountRole, actorRole string, roles map[string]string) error {
	var isAccountRoleExist bool = false
	for _, v := range roles {
		if v == accountRole {
			isAccountRoleExist = true
			break
		}
	}

	if !isAccountRoleExist {
		return errors.New("")
	}

	if actorRole == roles[staff_role] {
		if accountRole == roles[admin_role] {
			return errors.New("")
		}
	}

	return nil
}

func verifyEditUserAuthorization(req dto.UpdateUserReq, account, actor *dto.UserDBResModel, roles map[string]string) error {
	// Edited
	if req.UserId != actor.UserId {
		return verifyEditedAuth(req.RoleId, account.RoleId, actor.RoleId, roles)
	}

	// Self edit
	if req.RoleId != account.RoleId {
		return errors.New("")
	}

	return nil
}

func verifyEditedAuth(inputedRole, orgRole, actorRole string, roles map[string]string) error {
	var res error
	var authErr error = errors.New(noti.AuthorizationWarnMsg)

	switch actorRole {
	case roles[admin_role]:
		if orgRole == roles[admin_role] { // Admin edited admin
			res = authErr
		}
	case roles[staff_role]:
		if inputedRole != "" || inputedRole != orgRole { // Staff edits other
			res = authErr
		}
	}

	return res
}

func setUpBeforeResetPw(usc *business_object.UserSecurity, logger *log.Logger) (string, error) {
	token, err := util.GenerateActionToken("", usc.UserId, "", logger)
	if err != nil {
		return "", err
	}

	*usc.ActionToken = token

	return util.ToCombinedString([]string{
		getProcessUrl(),
		token,
	}, sepChar), nil
}
