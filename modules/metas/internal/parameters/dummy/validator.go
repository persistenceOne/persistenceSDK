/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package dummy

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func validator(i interface{}) error {
	switch value := i.(type) {
	case types.Parameter:
		data, Error := value.GetData().AsDec()
		if Error != nil || value.GetKey() != Key || data.IsZero() {
			return errors.InvalidParameter
		}
		return nil
	default:
		return errors.IncorrectFormat
	}
}