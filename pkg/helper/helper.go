package helper

import (
	"errors"
	"strconv"
)

func ConvertToUint(value string) (uint, error) {
	idUint, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return uint(idUint), nil
}
