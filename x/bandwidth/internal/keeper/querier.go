package keeper

import (
	//"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cybercongress/cyberd/x/bandwidth/exported"
	"github.com/cybercongress/cyberd/x/bandwidth/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier returns a minting Querier handler. k exported.StateKeeper
func NewQuerier(k exported.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, _ abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParameters:
			return queryParams(ctx, k)

		case types.QueryDesirableBandwidth:
			return queryDesirableBandwidth(ctx, k)

		case types.QueryMaxBlockBandwidth:
			return queryMaxBlockBandwidth(ctx, k)

		case types.QueryRecoveryPeriod:
				return queryRecoveryPeriod(ctx, k)

		case types.QueryAdjustPricePeriod:
			return queryAdjustPricePeriod(ctx, k)

		case types.QueryBaseCreditPrice:
			return queryBaseCreditPrice(ctx, k)

		case types.QueryTxCost:
			return queryTxCost(ctx, k)

		case types.QueryLinkMsgCost:
			return queryLinkMsgCost(ctx, k)

		case types.QueryNonLinkMsgCost:
			return queryNonLinkMsgCost(ctx, k)


		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryParams(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryDesirableBandwidth(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.DesirableBandwidth)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryMaxBlockBandwidth(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.MaxBlockBandwidth)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryRecoveryPeriod(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.RecoveryPeriod)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAdjustPricePeriod(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.AdjustPricePeriod)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryBaseCreditPrice(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.BaseCreditPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTxCost(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.TxCost)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryLinkMsgCost(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.LinkMsgCost)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNonLinkMsgCost(ctx sdk.Context, k exported.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params.NonLinkMsgCost)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}


