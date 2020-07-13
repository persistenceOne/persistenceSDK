package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	//message := messageFromInterface(msg)
	//assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(message.AssetID)
	//asset := assets.Get(message.AssetID)
	//if asset == nil {
	//	return constants.EntityNotFound
	//}
	//if !asset.CanBurn(types.NewHeight(context.BlockHeight())) {
	//	return constants.BurnNotAllowed
	//}
	//assets.Remove(asset)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}