package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	business_object "social_network/business_object"
	"social_network/constant/noti"
	"social_network/interfaces/repo"
	"time"
)

type roleRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeRoleRepo(db *sql.DB, logger *log.Logger) repo.IRoleRepo {
	return &roleRepo{
		db:     db,
		logger: logger,
	}
}

// ActivateRole implements repo.IRoleRepo.
func (r *roleRepo) ActivateRole(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "ActivateRole - "
	var query string = "Update " + business_object.GetRoleTable() + " set active_status = true, updated_at = ? where id = ?"
	defer r.db.Close()

	if _, err := r.db.Exec(query, fmt.Sprint(time.Now().UTC()), id); err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// CreateRole implements repo.IRoleRepo.
func (r *roleRepo) CreateRole(role business_object.Role, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "CreateRole - "
	var query string = "Insert into " + business_object.GetRoleTable() + "(role_id, role_name, active_status, created_at, updated_at) values (?, ?, ?, ?, ?)"
	defer r.db.Close()

	if _, err := r.db.Exec(query, role.RoleId, role.RoleName, role.ActiveStatus, role.CreatedAt, role.UpdatedAt); err != nil {
		r.logger.Println(errLogMsg, err)
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAllRoles implements repo.IRoleRepo.
func (r *roleRepo) GetAllRoles(ctx context.Context) (*[]business_object.Role, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "GetAllRoles - "
	var query string = "Select * from " + business_object.GetRoleTable()
	var internalErr error = errors.New(noti.InternalErr)
	defer r.db.Close()

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Role
	for rows.Next() {
		var x business_object.Role
		if err := rows.Scan(&x.RoleId, &x.RoleName, &x.ActiveStatus, &x.CreatedAt, &x.UpdatedAt); err != nil {
			r.logger.Println(errLogMsg, err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetRoleById implements repo.IRoleRepo.
func (r *roleRepo) GetRoleById(id string, ctx context.Context) (*business_object.Role, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "GetRoleById - "
	var query string = "Select * from " + business_object.GetRoleTable() + " where id = ?"
	defer r.db.Close()

	var res *business_object.Role
	if err := r.db.QueryRow(query, id).Scan(&res.RoleId, &res.RoleName, &res.ActiveStatus, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Println(errLogMsg, err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return res, nil
}

// GetRolesByName implements repo.IRoleRepo.
func (r *roleRepo) GetRolesByName(name string, ctx context.Context) (*[]business_object.Role, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "GetRolesByName - "
	var query string = "Select * from " + business_object.GetRoleTable() + " where lower(name) = lower($1)"
	var internalErr error = errors.New(noti.InternalErr)
	defer r.db.Close()

	rows, err := r.db.Query(query, name)
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Role
	for rows.Next() {
		var x business_object.Role
		if err := rows.Scan(&x.RoleId, &x.RoleName, &x.ActiveStatus, &x.CreatedAt, &x.UpdatedAt); err != nil {
			r.logger.Println(errLogMsg, err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// GetRolesByStatus implements repo.IRoleRepo.
func (r *roleRepo) GetRolesByStatus(status bool, ctx context.Context) (*[]business_object.Role, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "GetRolesByStatus - "
	var query string = "Select * from " + business_object.GetRoleTable() + " where active_status = ?"
	var internalErr error = errors.New(noti.InternalErr)
	defer r.db.Close()

	rows, err := r.db.Query(query, fmt.Sprint(status))
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return nil, internalErr
	}

	var res *[]business_object.Role
	for rows.Next() {
		var x business_object.Role
		if err := rows.Scan(&x.RoleId, &x.RoleName, &x.ActiveStatus, &x.CreatedAt, &x.UpdatedAt); err != nil {
			r.logger.Println(errLogMsg, err.Error())
			return nil, internalErr
		}

		*res = append(*res, x)
	}

	return res, nil
}

// RemoveRole implements repo.IRoleRepo.
func (r *roleRepo) RemoveRole(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "RemoveRole - "
	var query string = "Update " + business_object.GetRoleTable() + " set active_status = false, updated_at = ? where id = ?"
	var internalErr error = errors.New(noti.InternalErr)
	defer r.db.Close()

	res, err := r.db.Exec(query, fmt.Sprint(time.Now().UTC()), id)
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetRoleTable()))
	}

	return nil
}

// UpdateRole implements repo.IRoleRepo.
func (r *roleRepo) UpdateRole(role business_object.Role, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetRoleTable()) + "UpdateRole - "
	var query string = "Update " + business_object.GetRoleTable() + " set role_name = ?, active_status = ?, updated_at = ? where id = ?"
	var internalErr error = errors.New(noti.InternalErr)
	defer r.db.Close()

	res, err := r.db.Exec(query, role.RoleName, role.ActiveStatus, role.UpdatedAt, role.RoleId)
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetRoleTable()))
	}

	return nil
}
