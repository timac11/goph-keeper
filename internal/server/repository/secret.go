package repository

import (
	"context"
	internalErrors "errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/timac11/goph-keeper/internal/common/model"
	"github.com/timac11/goph-keeper/internal/server/errors"
)

func (client *PgClient) CreateSecret(ctx context.Context, secret *model.SecretCreateDto, userId string) (*model.SecretInfoDto, error) {
	var secretModel model.SecretInfoDto

	query, _, err := sq.Insert("user").
		Columns("id", "name", "type", "data", "metadata", "public_key").
		Values(userId, secret.Name, secret.Type, secret.Data, secret.Metadata, secret.PublicKey).
		Suffix("RETURNING id, name").
		ToSql()

	err = client.pool.
		QueryRow(ctx, query).
		Scan(&secretModel.ID, &secretModel.Name)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return nil, errors.NewEntityError(err, errors.EntityAlreadyExists, secret)
		}
		return nil, err
	}

	return &secretModel, nil
}

func (client *PgClient) GetSecret(ctx context.Context, id string) (*model.Secret, error) {
	var secret model.Secret

	query, _, err := sq.Select("id", "name", "type", "data", "metadata", "public_key", "created_at", "updated_at").
		From("secret").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = client.pool.
		QueryRow(ctx, query).
		Scan(&secret.ID, &secret.Name, &secret.Type, &secret.Data, &secret.Metadata, &secret.PublicKey, &secret.CreatedAt, &secret.UpdatedAt)

	if err != nil {
		if internalErrors.Is(err, pgx.ErrNoRows) {
			return nil, errors.NewEntityError(err, errors.EntityNotFound, id)
		}
		return nil, err
	}

	return &secret, nil
}

func (client *PgClient) UpdateSecret(ctx context.Context, secret *model.Secret) (*model.Secret, error) {
	return nil, internalErrors.New("Not implemented")
}

func (client *PgClient) DeleteSecret(ctx context.Context, id string) error {
	query, _, err := sq.Delete("secret").Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	_, err = client.pool.Exec(ctx, query)

	return err

}

func (client *PgClient) GetSecretList(ctx context.Context, id string) (*[]model.SecretInfoDto, error) {
	query, _, err := sq.
		Select("id", "name", "type", "updated_at", "created_at").
		From("secret").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := client.pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []model.SecretInfoDto

	for rows.Next() {
		var secret model.SecretInfoDto
		if err := rows.Scan(&secret.ID, &secret.Name, &secret.Type, &secret.UpdatedAt, &secret.CreatedAt); err != nil {
			return nil, err
		}

		result = append(result, secret)
	}

	return &result, nil
}
