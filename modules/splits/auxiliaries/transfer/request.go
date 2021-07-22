/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package transfer

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type auxiliaryRequest struct {
	FromID    test_types.ID     `json:"fromID" valid:"required~required field fromID missing"`
	ToID      test_types.ID     `json:"toID" valid:"required~required field toID missing"`
	OwnableID test_types.ID     `json:"ownableID" valid:"required~required field ownableID missing"`
	Value     sdkTypes.Dec `json:"value" valid:"required~required field value missing"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(auxiliaryRequest)
	return Error
}

func auxiliaryRequestFromInterface(request helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := request.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(fromID test_types.ID, toID test_types.ID, ownableID test_types.ID, value sdkTypes.Dec) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		FromID:    fromID,
		ToID:      toID,
		OwnableID: ownableID,
		Value:     value,
	}
}
