package mutate

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper mapper.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	assetID := message.AssetID
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	asset := assets.Get(assetID)
	if asset == nil {
		return constants.EntityNotFound
	}
	mutableProperties := asset.GetMutables().Get()
	for _, property := range message.Properties.GetList() {
		if mutableProperties.Get(property.GetID()) == nil {
			return constants.MutableNotFound
		}
		mutableProperties = mutableProperties.Mutate(property)
	}
	asset = mapper.NewAsset(asset.GetID(), types.NewMutables(mutableProperties, asset.GetMaintainersID()), asset.GetImmutables(), asset.GetLock(), asset.GetBurn())
	assets = assets.Mutate(asset)
	return nil
}

func NewTransactionKeeper(mapper mapper.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}