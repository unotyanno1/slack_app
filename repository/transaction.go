package repository

import (
	"context"
	"database/sql"
)

var txKey = struct{}{}

func GetTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sql.Tx)
	return tx, ok
}
