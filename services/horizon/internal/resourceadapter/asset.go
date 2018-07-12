package resourceadapter

import (
	"context"

	"github.com/caoxuwen/go/xdr"
	. "github.com/caoxuwen/go/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
