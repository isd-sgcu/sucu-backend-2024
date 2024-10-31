package utils

import (
	"strings"

	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
)

func validate(word string, words []string) bool {
	if word == "" {
		return true
	}

	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func ValidateDocType(docType string) bool {
	docs := []string{
		constant.ANNOUNCEMENT,
		constant.STATISTIC,
		constant.BUDGET,
	}

	return validate(strings.ToUpper(docType), docs)
}

func ValidateAttachmentType(attachmentType string) bool {
	attachment := []string{
		constant.IMAGE,
		constant.DOCS,
	}

	return validate(strings.ToUpper(attachmentType), attachment)
}

func ValidateOrg(org string) bool {
	orgs := []string{
		constant.SCCU,
		constant.SGCU,
	}

	return validate(strings.ToUpper(org), orgs)
}

func ValidateRole(role string) bool {
	roles := []string{
		constant.SCCU_ADMIN,
		constant.SCCU_SUPERADMIN,
		constant.SGCU_ADMIN,
		constant.SGCU_SUPERADMIN,
	}

	return validate(strings.ToUpper(role), roles)
}
