package storage

import (
	"context"
	"database/sql"
	"errors"

	"magnet-test-task/internal/generated/models"

	_ "github.com/mattn/go-sqlite3"
)

var UserNotFound = errors.New("user not found")

type Storage struct {
	db *sql.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetInstance() *sql.DB {
	return s.db
}

func (s *Storage) Store(ctx context.Context, user *models.User) error {
	if user.UID == 0 {
		return s.new(ctx, user)
	}

	return s.update(ctx, user)
}

func (s *Storage) Delete(ctx context.Context, uid int64) error {
	query := `DELETE FROM user WHERE id = ?`
	if _, err := s.db.ExecContext(ctx, query, uid); err != nil {
		return err
	}

	return nil
}

func (s *Storage) OneByUID(ctx context.Context, uid int64) (*models.User, error) {
	query := `SELECT name, birth_date FROM user WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, uid)
	user := &models.User{UID: uid}

	if err := row.Scan(&user.Name, &user.BirthDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, UserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *Storage) new(ctx context.Context, user *models.User) error {
	query := `INSERT INTO user (name, birth_date) VALUES (?, ?)`
	result, err := s.db.ExecContext(ctx, query, user.Name, user.BirthDate)
	if err != nil {
		return err
	}

	uid, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.UID = uid

	return nil
}

func (s *Storage) update(ctx context.Context, user *models.User) error {
	query := `UPDATE user SET name = ?, birth_date = ? WHERE id = ?`
	res, err := s.db.ExecContext(ctx, query, user.Name, user.BirthDate, user.UID)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected == 0 {
		return UserNotFound
	}

	return nil
}
