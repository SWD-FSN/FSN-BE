package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
)

type userSecurityRepo struct {
	db     *sql.DB
	logger *log.Logger
}

// EditUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) EditUserSecurity(usc business_object.UserSecurity, ctx context.Context) error {
	panic("unimplemented")
}

// GetUserSecurity implements repo.IUserSecurityRepo.
func (u *userSecurityRepo) GetUserSecurity(id string, ctx context.Context) (*business_object.UserSecurity, error) {
	var query string = "Select top 1 * from " + business_object.GetUserSecurityTable() + "Where id = ?"
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetUserSecurityTable()) + "GetUserSecurity - "
	defer u.db.Close()

}

func InitializeUserSecurityRepo(db *sql.DB, logger *log.Logger) repo.IUserSecurityRepo {
	return &userSecurityRepo{
		db:     db,
		logger: logger,
	}
}
