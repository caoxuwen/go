package resourceadapter

import (
	"context"

	"github.com/caoxuwen/go/amount"
	. "github.com/caoxuwen/go/protocols/horizon"
	"github.com/caoxuwen/go/services/horizon/internal/assets"
	"github.com/caoxuwen/go/services/horizon/internal/db2/core"
	"github.com/caoxuwen/go/xdr"
)

func PopulateBalance(ctx context.Context, dest *Balance, row core.Trustline) (err error) {
	dest.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	dest.Balance = amount.String(row.Balance)
	dest.BuyingLiabilities = amount.String(row.BuyingLiabilities)
	dest.SellingLiabilities = amount.String(row.SellingLiabilities)
	dest.Limit = amount.String(row.Tlimit)
	dest.Liquidation = ((row.Flags & (int32)(xdr.TrustLineFlagsLiquidationFlag)) != 0)
	dest.Issuer = row.Issuer
	dest.Code = row.Assetcode
	dest.Debt = amount.String(row.Debt)
	return
}

func PopulateNativeBalance(dest *Balance, stroops, buyingLiabilities, sellingLiabilities xdr.Int64) (err error) {
	dest.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	dest.Balance = amount.String(stroops)
	dest.BuyingLiabilities = amount.String(buyingLiabilities)
	dest.SellingLiabilities = amount.String(sellingLiabilities)
	dest.Limit = ""
	dest.Issuer = ""
	dest.Code = ""
	return
}
