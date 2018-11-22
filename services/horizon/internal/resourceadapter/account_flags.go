package resourceadapter

import (
	"github.com/caoxuwen/go/services/horizon/internal/db2/core"
	. "github.com/caoxuwen/go/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
