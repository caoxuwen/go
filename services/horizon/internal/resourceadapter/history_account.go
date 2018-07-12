package resourceadapter

import (
	"context"

	"github.com/caoxuwen/go/services/horizon/internal/db2/history"
	. "github.com/caoxuwen/go/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
