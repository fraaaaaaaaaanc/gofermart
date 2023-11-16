package utils

import (
	"bytes"
	"io"
)

func IsRequestBodyEmpty(body io.Reader) (bool, error) {
	newBody, err := io.ReadAll(body)
	if err != nil {
		return false, err
	}
	body = io.NopCloser(bytes.NewReader(newBody))
	if len(newBody) == 0 {
		return true, nil
	}
	return false, nil
}
