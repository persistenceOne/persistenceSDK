/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Parameters interface {
	String() string

	GetList() []types.Parameter

	Validate() error
	Equal(Parameters) bool
	KeyTable() subspace.KeyTable
	subspace.ParamSet
}