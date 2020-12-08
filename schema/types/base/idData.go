/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type idData struct {
	Value types.ID `json:"value"`
}

var _ types.Data = (*idData)(nil)

func (IDData idData) String() string {
	return IDData.Value.String()
}
func (IDData idData) ZeroValue() types.Data {
	return NewIDData(NewID(""))
}
func (IDData idData) Type() string {
	return "I"
}
func (IDData idData) GenerateHash() string {
	if IDData.Value.String() == "" {
		return IDData.Type() + constants.DataTypeAndValueSeparator
	}
	return IDData.Type() + constants.DataTypeAndValueSeparator + meta.Hash(IDData.Value.String())
}
func (IDData idData) AsString() (string, error) {
	return "", errors.EntityNotFound
}
func (IDData idData) AsDec() (sdkTypes.Dec, error) {
	return sdkTypes.Dec{}, errors.EntityNotFound
}
func (IDData idData) AsHeight() (types.Height, error) {
	return height{}, errors.EntityNotFound
}
func (IDData idData) AsID() (types.ID, error) {
	return IDData.Value, nil
}
func (IDData idData) Get() interface{} {
	return IDData.Value
}
func (IDData idData) Equal(data types.Data) bool {
	switch value := data.(type) {
	case idData:
		return value.Value.Equals(IDData.Value)
	default:
		return false
	}
}

func NewIDData(value types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(NewID(idData)), nil
}
