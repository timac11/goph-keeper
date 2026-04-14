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

func (client *PgClient) CreateSecret(ctx context.Context, user *model.UserLoginDto) (*model.User, error) {
	
}

func (client *PgClient) GetSecret(ctx context.Context, id string) (*model.User, error) {
	
}

func (client *PgClient) UpdateSecret(ctx context.Context, id string) (*model.User, error) {
	
}

func (client *PgClient) DeleteSecret(ctx context.Context, id string) (*model.User, error) {
	
}

func (client *PgClient) ListSecret(ctx context.Context, id string) (*model.User, error) {
	
}
