package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	mutables := types.NewMutables(message.Properties, types.NewID("maintainerID"))
	immutables := types.NewImmutables(message.Properties)
	assetID := mapper.NewOrderID(types.NewID("chainID"), types.NewID("maintainerID"), types.NewID("classificationID"), immutables.GetHashID())
	asset := mapper.NewOrder(assetID, message.BuyCoins, message.SellCoins, mutables, immutables, types.NewHeight(-1), types.NewHeight(-1))
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExists
	}
	assets.Add(asset)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}