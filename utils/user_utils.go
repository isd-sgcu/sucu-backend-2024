package utils

import "github.com/isd-sgcu/sucu-backend-2024/utils/constant"

func GetRole(role string) (string, error) {
	switch role {
	case constant.SGCU_SUPERADMIN:
		role = constant.SGCU_ADMIN
	case constant.SCCU_SUPERADMIN:
		role = constant.SCCU_ADMIN
	default:
		return "", constant.ErrInvalidRole
	}

	return role, nil
}
