package link

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	cbd "github.com/cybercongress/cyberd/types"
	cbdlink "github.com/cybercongress/cyberd/x/link/internal/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cybercongress/cyberd/types"
)

// NewHandler returns a handler for "link" type messages.
// cis  - cids index storage
// ils  - links storage
// as   - account storage
// imms - in-memory storage
func NewLinksHandler(cis CidNumberKeeper, ls IndexedKeeper, as auth.AccountKeeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {

		linkMsg := msg.(cbdlink.Msg)

		//validations
		//todo: optimize
		for _, link := range linkMsg.Links {
			// if cid not exists it automatically means that this is new link
			fromCidNumber, exists := cis.GetCidNumber(ctx, link.From)
			if !exists {
				continue
			}
			toCidNumber, exists := cis.GetCidNumber(ctx, link.To)
			if !exists {
				continue
			}

			accNumber := cbd.AccNumber(as.GetAccount(ctx, linkMsg.Address).GetAccountNumber())
			compactLink := cbdlink.NewLink(fromCidNumber, toCidNumber, accNumber)

			if ls.IsLinkExist(compactLink) {
				return nil, types.ErrCyberlinkExist
			}
		}

		for _, link := range linkMsg.Links {
			fromCidNumber := cis.GetOrPutCidNumber(ctx, link.From)
			toCidNumber := cis.GetOrPutCidNumber(ctx, link.To)
			accNumber := cbd.AccNumber(as.GetAccount(ctx, linkMsg.Address).GetAccountNumber())

			ls.PutLink(ctx, cbdlink.NewLink(fromCidNumber, toCidNumber, accNumber))
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName), // TODO
			),
		)

		return &sdk.Result{Events: ctx.EventManager().Events()}, nil
	}
}
