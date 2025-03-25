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

type postRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializePostRepo(db *sql.DB, logger *log.Logger) repo.IPostRepo {
	return &postRepo{
		db:     db,
		logger: logger,
	}
}

// CreatePost implements repo.IPostRepo.
func (p *postRepo) CreatePost(post business_object.Post, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "CreatePost - "
	var query string = "INSERT INTO " + business_object.GetPostTable() + "(id, author_id, content, is_private, is_hidden, created_at, updated_at, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	//defer p.db.Close()

	if _, err := p.db.Exec(query, post.PostId, post.AuthorId, post.Content, post.IsPrivate, post.IsHidden, post.CreatedAt, post.UpdatedAt, post.Status); err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	return nil
}

// GetAllPosts implements repo.IPostRepo.
func (p *postRepo) GetAllPosts(ctx context.Context) (*[]business_object.Post, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "GetAllPosts - "
	var query string = "SELECT * FROM " + business_object.GetPostTable()
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Post
	for rows.Next() {
		var x business_object.Post

		if err := rows.Scan(&x.PostId, &x.AuthorId, &x.Content, &x.IsPrivate, &x.IsHidden, &x.CreatedAt, &x.UpdatedAt, &x.Status); err != nil {
			p.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetPostsByKeyword implements repo.IPostRepo.
func (p *postRepo) GetPostsByKeyword(keyword string, ctx context.Context) (*[]business_object.Post, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "GetPostsByKeyword - "
	var query string = "SELECT * FROM " + business_object.GetPostTable() + " WHERE is_private = false, is_hidden = false, status = true AND LOWER(content) LIKE LOWER('%?%')"
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	rows, err := p.db.Query(query, keyword)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Post
	for rows.Next() {
		var x business_object.Post

		if err := rows.Scan(&x.PostId, &x.AuthorId, &x.Content, &x.IsPrivate, &x.IsHidden, &x.CreatedAt, &x.UpdatedAt, &x.Status); err != nil {
			p.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetPosts implements repo.IPostRepo.
func (p *postRepo) GetPosts(ctx context.Context) (*[]business_object.Post, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "GetPosts - "
	var query string = "SELECT * FROM " + business_object.GetPostTable() + " WHERE is_private = false, is_hidden = false AND status = true"
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Post
	for rows.Next() {
		var x business_object.Post

		if err := rows.Scan(&x.PostId, &x.AuthorId, &x.Content, &x.IsPrivate, &x.IsHidden, &x.CreatedAt, &x.UpdatedAt, &x.Status); err != nil {
			p.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetPost implements repo.IPostRepo.
func (p *postRepo) GetPost(id string, ctx context.Context) (*business_object.Post, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "GetPost - "
	var query string = "SELECT * FROM " + business_object.GetPostTable() + " WHERE id = $1"

	//defer p.db.Close()

	var res business_object.Post
	if err := p.db.QueryRow(query, id).Scan(&res.PostId, &res.AuthorId, &res.Content, &res.IsPrivate, &res.IsHidden, &res.CreatedAt, &res, res.UpdatedAt, &res.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		p.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.InternalErr)
	}

	return &res, nil
}

// GetPostsByUser implements repo.IPostRepo.
func (p *postRepo) GetPostsByUser(id string, ctx context.Context) (*[]business_object.Post, error) {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetLikeTable()) + "GetPostsByUser - "
	var query string = "SELECT * FROM " + business_object.GetPostTable() + " WHERE author_id = $1"
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	rows, err := p.db.Query(query, id)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []business_object.Post
	for rows.Next() {
		var x business_object.Post

		if err := rows.Scan(&x.PostId, &x.AuthorId, &x.Content, &x.IsPrivate, &x.IsHidden, &x.CreatedAt, &x.UpdatedAt, &x.Status); err != nil {
			p.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// RemovePost implements repo.IPostRepo.
func (p *postRepo) RemovePost(id string, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "RemovePost - "
	var query string = "UPDATE " + business_object.GetPostTable() + " SET status = false, updated_at = $1 WHERE id = $2"
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	res, err := p.db.Exec(query, time.Now().String(), id)
	if err != nil {
		p.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		p.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetRoleTable()))
	}

	return nil
}

// UpdatePost implements repo.IPostRepo.
func (p *postRepo) UpdatePost(post business_object.Post, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.RepoErrMsg, business_object.GetPostTable()) + "UpdatePost - "
	var query string = "UPDATE " + business_object.GetPostTable() + " SET content = $1, is_private = $2, is_hidden = $3, updated_at = $4 AND status = $5 WHERE id = $6"
	var internalErr error = errors.New(noti.InternalErr)

	//defer p.db.Close()

	res, err := p.db.Exec(query, post.Content, post.IsPrivate, post.IsHidden, post.UpdatedAt, post.Status, post.PostId)
	if err != nil {
		p.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		p.logger.Println(errLogMsg, err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UndefinedObjectWarnMsg, business_object.GetRoleTable()))
	}

	return nil
}
