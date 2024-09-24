package utils

import (
	"errors"
	"strings"

	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
)

func GetDocType(docType string) (string, error) {
	switch strings.ToUpper(docType) {
	case constant.ANNOUNCEMENT:
		docType = constant.ANNOUNCEMENT
	case constant.BUDGET:
		docType = constant.BUDGET
	case constant.STATISTIC:
		docType = constant.STATISTIC
	default:
		return "", errors.New(constant.ErrInvalidDocType)
	}

	return docType, nil
}
