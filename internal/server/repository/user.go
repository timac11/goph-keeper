package repository

import (
	"context"
	internalErrors "errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/timac11/goph-keeper/internal/common/logger"
	"github.com/timac11/goph-keeper/internal/common/model"
	"github.com/timac11/goph-keeper/internal/server/errors"
)

func (client *PgClient) CreateUser(ctx context.Context, user *model.UserLoginDto) (*model.User, error) {
	var userModel model.User
	log := logger.LoggerFromContext(ctx)

	log.Info("Create user with params",
		zap.String("login", user.Login),
	)

	query, args, err := sq.Insert(`"user"`).
		Columns("login", "password").
		Values(user.Login, user.Password).
		Suffix("RETURNING id, login, password").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	err = client.pool.
		QueryRow(ctx, query, args...).
		Scan(&userModel.ID, &userModel.Login, &userModel.Password)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return nil, errors.NewEntityError(err, errors.EntityAlreadyExists, user)
		}
		return nil, err
	}

	return &userModel, nil
}

func (client *PgClient) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	var user model.User

	query, args, err := sq.Select("id", "login", "password").
		From(`"user"`).
		Where(sq.Eq{"login": login}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	err = client.pool.
		QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Login, &user.Password)

	if err != nil {
		if internalErrors.Is(err, pgx.ErrNoRows) {
			return nil, errors.NewEntityError(err, errors.EntityNotFound, login)
		}
		return nil, err
	}

	return &user, nil
}

func (client *PgClient) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	query, args, err := sq.Select("id", "login", "password").
		From(`"user"`).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	err = client.pool.
		QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Login, &user.Password)

	if err != nil {
		if internalErrors.Is(err, pgx.ErrNoRows) {
			return nil, errors.NewEntityError(err, errors.EntityNotFound, id)
		}
		return nil, err
	}

	return &user, nil
}
