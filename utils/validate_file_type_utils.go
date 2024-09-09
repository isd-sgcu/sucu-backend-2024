package utils

import (
	"errors"
	"strings"

	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
)

func ValidateFileType(fileName string) (*string, error) {
	// substring to access only the file extension
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex == -1 {
		return nil, errors.New("not found file extension")
	}
	fileExt := strings.ToLower(fileName[lastDotIndex:])

	// iterate ext over 2 arrays
	for _, ext := range constant.AllowedImageFileTypes {
		if ext == fileExt {
			result := constant.IMAGE
			return &result, nil
		}
	}

	for _, ext := range constant.AllowedDocsFileTypes {
		if ext == fileExt {
			result := constant.DOCS
			return &result, nil
		}
	}

	return nil, errors.New("invalid file type")
}
