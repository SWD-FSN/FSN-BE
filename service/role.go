package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
	"social_network/interfaces/service"
	"social_network/repository"
	"social_network/repository/db"
	"social_network/util"
	"strings"
	"time"

	"github.com/google/uuid"
)

type roleService struct {
	roleRepo repo.IRoleRepo
	logger   *log.Logger
}

const (
	id_type   string = "ID"
	name_type string = "NAME"
)

func InitializeRoleService(db *sql.DB, logger *log.Logger) service.IRoleService {
	return &roleService{
		roleRepo: repository.InitializeRoleRepo(db, logger),
		logger:   logger,
	}
}

func GenerateRoleService() (service.IRoleService, error) {
	var logger = util.GetLogConfig()

	cnn, err := db.ConnectDB(logger)

	if err != nil {
		return nil, err
	}

	return InitializeRoleService(cnn, logger), nil
}

// ActivateRole implements service.IRoleService.
func (r *roleService) ActivateRole(id string, ctx context.Context) error {
	return r.roleRepo.ActivateRole(id, ctx)
}

// CreateRole implements service.IRoleService.
func (r *roleService) CreateRole(name string, ctx context.Context) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New(noti.FieldEmptyWarnMsg)
	}
	//---------------------------------------
	if isRoleExist(r.roleRepo, name, name_type, ctx) {
		return errors.New(noti.ItemExistedWarnMsg)
	}
	//---------------------------------------
	return r.roleRepo.CreateRole(business_object.Role{
		RoleId:       fmt.Sprint(uuid.New()),
		RoleName:     name,
		ActiveStatus: true,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}, ctx)
}

// GetAllRoles implements service.IRoleService.
func (r *roleService) GetAllRoles(ctx context.Context) (*[]business_object.Role, error) {
	return r.roleRepo.GetAllRoles(ctx)
}

// GetRoleById implements service.IRoleService.
func (r *roleService) GetRoleById(id string, ctx context.Context) (*business_object.Role, error) {
	if id == "" {
		return nil, errors.New(noti.GenericsErrorWarnMsg)
	}
	//---------------------------------------
	res, err := r.roleRepo.GetRoleById(id, ctx)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New(business_object.GetRoleTable() + noti.UndefinedObjectWarnMsg)
	}
	//---------------------------------------
	return res, nil
}

// GetRolesByName implements service.IRoleService.
func (r *roleService) GetRolesByName(name string, ctx context.Context) (*[]business_object.Role, error) {
	if name == "" {
		return r.roleRepo.GetAllRoles(ctx)
	}

	return r.roleRepo.GetRolesByName(util.ToNormalizedString(name), ctx)
}

// GetRolesByStatus implements service.IRoleService.
func (r *roleService) GetRolesByStatus(rawStatus string, ctx context.Context) (*[]business_object.Role, error) {
	if rawStatus == "" {
		return r.roleRepo.GetAllRoles(ctx)
	}

	status, err := util.ToBoolean(rawStatus)
	if err != nil {
		return nil, err
	}

	return r.roleRepo.GetRolesByStatus(status, ctx)
}

// RemoveRole implements service.IRoleService.
func (r *roleService) RemoveRole(id string, ctx context.Context) error {
	if id == "" {
		return errors.New(noti.GenericsErrorWarnMsg)
	}

	return r.roleRepo.RemoveRole(id, ctx)
}

// UpdateRole implements service.IRoleService.
func (r *roleService) UpdateRole(role business_object.Role, ctx context.Context) error {
	if !isRoleExist(r.roleRepo, role.RoleId, id_type, ctx) {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetRoleTable()))
	}

	if role.RoleName == "" {
		return nil
	}

	role.RoleName = strings.TrimSpace(role.RoleName)
	role.UpdatedAt = time.Now().UTC()
	return r.roleRepo.UpdateRole(role, ctx)
}

func isRoleExist(repo repo.IRoleRepo, prob, validateProb string, ctx context.Context) bool {
	var res interface{}

	if validateProb == id_type {
		res, _ = repo.GetRoleById(prob, ctx)
	} else {
		res, _ = repo.GetRolesByName(prob, ctx)
	}

	return res != nil
}
