package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gitopia/gitopia/x/gitopia/keeper"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func GitopiaKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)

	amino := codec.NewLegacyAmino()
	ss := typesparams.NewSubspace(appCodec,
		amino,
		storeKey,
		memStoreKey,
		"GitopiaSubSpace",
	)

	ak := authkeeper.NewAccountKeeper(
		appCodec,
		storeKey,
		ss,
		nil,
		nil,
		"gitopia",
	)

	authzKeeper := authzkeeper.NewKeeper(
		storeKey,
		appCodec,
		nil,
		ak,
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		storeKey,
		ak,
		ss,
		nil,
	)

	mintKeeper := mintkeeper.NewKeeper(
		appCodec, storeKey, ss, nil, nil,
		bankKeeper, authtypes.FeeCollectorName)

	k := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		types.MinterAccountName,
		authtypes.FeeCollectorName,
		ak,
		&authzKeeper,
		bankKeeper,
		mintKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)
	return k, ctx
}
