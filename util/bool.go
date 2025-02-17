package util

import (
	"errors"
	"social_network/constant/noti"
	"strconv"
)

func ToBoolean(rawStatus string) (bool, error) {
	res, err := strconv.ParseBool(ToNormalizedString(rawStatus))

	if err != nil {
		return false, errors.New(noti.InvalidStatusWarnMsg)
	}

	return res, nil
}

func IsBooleanRemain(input *bool, org bool) bool {
	// Remain
	if input == nil {
		return true
	}

	return *input == org
}
