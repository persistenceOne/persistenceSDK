/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type fact struct {
	Hash       string           `json:"hash"`
	Signatures types.Signatures `json:"signatures"`
	Meta       bool             `json:"meta"`
}

var _ types.Fact = (*fact)(nil)

func (fact fact) Get() string                     { return "" }
func (fact fact) GetHash() string                 { return fact.Hash }
func (fact fact) GetSignatures() types.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return fact.Meta
}
func (fact fact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return fact
}

func NewFact(Fact string) types.Fact {
	return fact{
		Hash:       metaUtilities.Hash(Fact),
		Signatures: signatures{},
		Meta:       false,
	}
}

func MetaFactToFact(MetaFact types.Fact) types.Fact {
	switch value := MetaFact.(type) {
	case fact:
		return value
	case metaFact:
		return fact{
			Hash:       MetaFact.GetHash(),
			Signatures: MetaFact.GetSignatures(),
			Meta:       MetaFact.IsMeta(),
		}
	default:
		return fact{}
	}
}
