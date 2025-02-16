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
	"sync"
	"time"
)

type userService struct {
	logger           *log.Logger
	roleRepo         repo.IRoleRepo
	userSecurityRepo repo.IUserSecurityRepo
	userRepo         repo.IUserRepo
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
	verifyFailLimit int = 5
)

const (
	admin_role string = "Admin"
	user_role  string = "User"
	staff_role string = "Staff"
)

func GenerateUserService() (service.IUserService, error) {
	db, err := db.ConnectDB(business_object.GetUserTable())

	if err != nil {
		return nil, err
	}

	var logger *log.Logger = &log.Logger{}

	return &userService{
		logger:           logger,
		roleRepo:         repository.InitializeRoleRepo(db, logger),
		userSecurityRepo: repository.InitializeUserSecurityRepo(db, logger),
		userRepo:         repository.InitializeUserRepo(db, logger),
	}, nil
}

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

// ChangeUserStatus implements service.IUserService.
func (u *userService) ChangeUserStatus(rawStatus string, userId string, actorId string, c context.Context) (error, string) {
	panic("unimplemented")
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

	// Save new account to database
	if err := u.userRepo.CreateUser(dto.UserDBResModel{
		UserId:          id,
		RoleId:          req.RoleId,
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
