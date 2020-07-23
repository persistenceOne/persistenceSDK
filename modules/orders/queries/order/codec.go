package order

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, QueryRoute+"/"+"request", nil)
	codec.RegisterConcrete(queryResponse{}, QueryRoute+"/"+"response", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	types.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}